package model

import "github.com/thernande/movie-micro/gen"

// MetadataToProto converts a Metadata struct into a
// generated proto counterpart.

func MetadataToProto(metadata *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id:          metadata.ID,
		Title:       metadata.Title,
		Description: metadata.Description,
		Director:    metadata.Director,
	}
}

// MetadataFromProto converts a generated proto counterpart into a Metadata struct.
func MetadataFromProto(proto *gen.Metadata) *Metadata {
	return &Metadata{
		ID:          proto.Id,
		Title:       proto.Title,
		Description: proto.Description,
		Director:    proto.Director,
	}
}
