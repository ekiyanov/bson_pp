package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	var payload = &bytes.Buffer{}
	_, err := io.Copy(payload, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var doc bson.D

	if err = bson.Unmarshal(payload.Bytes(), &doc); err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := bson.MarshalExtJSON(&doc, true, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
