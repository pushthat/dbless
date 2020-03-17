package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pushthat/dbless/cmd/dbless"
)

func main() {
	serializerType := "json"
	s3Bucket := "dbless-test"
	index := "id"
	prefix := "db-less/"

	endpoint := "nyc3.digitaloceanspaces.com"
	region := "nyc3"
	s3Session, err := session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	})
	if err != nil {
		log.Fatal(err)
	}

	dbless := dbless.DbLess{
		SerializerType: &serializerType,
		S3Bucket:       &s3Bucket,
		Index:          &index,
		Prefix:         &prefix,
		S3Session:      s3Session,
	}
	dbless.Init()

	objToSave := map[string]interface{}{
		"id":   "1",
		"name": "matias",
	}
	err = dbless.Save(objToSave)
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
