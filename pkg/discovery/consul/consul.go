package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"github.com/thernande/movie-micro/pkg/discovery"
)

// Registry defines a service registry.
type Registry struct {
	client *consul.Client
}

// NewRegistry creates a new Consul service registry instance.
func NewRegistry(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client: client}, nil
}

// Register creates a service instance record in the registry.
func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID:      instanceID,
		Name:    serviceName,
		Tags:    []string{hostPort},
		Port:    port,
		Address: parts[0],
		Check: &consul.AgentServiceCheck{
			TTL:     "5s",
			CheckID: instanceID,
		},
	})
}

// Deregister removes a service insttance record from the registry.
func (r *Registry) Deregister(ctx context.Context, instanceID string, _ string) error {
	return r.client.Agent().ServiceDeregister(instanceID)
}

// ServiceAddresses returns the list of addresses of active instances of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	var addresses []string
	services, _, err := r.client.Health().Service(serviceName, "", true, &consul.QueryOptions{})
	if err != nil {
		return nil, err
	} else if len(services) == 0 {
		return nil, discovery.ErrNotFound
	}
	for _, service := range services {
		addresses = append(addresses, fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port))
	}
	return addresses, nil
}

// ReportHealthyState is a push mechanism for reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(instanceID string, serviceName string) error {
	return r.client.Agent().PassTTL(instanceID, "")
}
