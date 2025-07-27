package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"google.golang.org/protobuf/proto"
	"inventory.com/catalog/pkg/model"
	"inventory.com/gen"
)

var metadata = &model.Category{
	ID:   123,
	Name: "Elecrtonics",
}

var genMetadata = &gen.Category{
	Id:   123,
	Name: "Elecrtonics",
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

	fmt.Printf("JSON size:\t%dB\n", len(jsonBytes))
	fmt.Printf("XML size:\t%dB\n", len(xmlBytes))
	fmt.Printf("Proto size:\t%dB\n", len(protoBytes))
}

func serializeToJSON(m *model.Category) ([]byte, error) {
	return json.Marshal(m)
}

func serializeToXML(m *model.Category) ([]byte, error) {
	return xml.Marshal(m)
}

func serializeToProto(m *gen.Category) ([]byte, error) {
	return proto.Marshal(m)
}
