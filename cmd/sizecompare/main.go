package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/thernande/movie-micro/gen"
	"github.com/thernande/movie-micro/metadata/pkg/model"
	"google.golang.org/protobuf/proto"
)

var metadata = &model.Metadata{
	ID:          "123",
	Title:       "super dragon ball movie",
	Description: "es una pelicula de dragon ball super en la cual el superpoder es super fuerza,es una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerza",
	Director:    "foo bar",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "super dragon ball movie",
	Description: "es una pelicula de dragon ball super en la cual el superpoder es super fuerza,es una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerzaes una pelicula de dragon ball super en la cual el superpoder es super fuerza",
	Director:    "foo bar",
}

func serializeToJSON(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

func serializeToXML(m *model.Metadata) ([]byte, error) {
	return xml.Marshal(m)
}

func serializeToProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}
func main() {
	jsonBytes, err := serializeToJSON(metadata)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serializeToXML(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("json: %d\nxml: %d\nproto: %d\n", len(jsonBytes), len(xmlBytes), len(protoBytes))
}
