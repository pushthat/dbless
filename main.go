package main

import (
	"fmt"
	"log"

	"github.com/pushthat/dbless/cmd/dbless"
)

func main() {
	serializerType := "json"
	s3Bucket := "dbless-test"
	index := "id"
	prefix := "db-less/"

	dbless := dbless.DbLess{
		SerializerType: &serializerType,
		S3Bucket:       &s3Bucket,
		Index:          &index,
		Prefix:         &prefix,
	}
	dbless.Init()

	objToSave := map[string]interface{}{
		"id":   "1",
		"name": "matias",
	}
	err := dbless.Save(objToSave)
	if err != nil {
		log.Fatal(err)
	}

	objToLoad := map[string]interface{}{
		"id": "1",
	}
	data, err := dbless.Load(objToLoad)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
