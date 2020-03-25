package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pushthat/dbless/cmd/dbless"
)

func main() {
	serializerType := "json"
	s3Bucket := "dbless-test"
	index := "id"
	prefix := "db-less/"

	s3Session, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	var DbLessSessions []*dbless.DbLessSession

	dblessSession := dbless.DbLessSession{
		S3Bucket:  &s3Bucket,
		Index:     &index,
		Prefix:    &prefix,
		S3Session: s3Session,
	}

	DbLessSessions = append(DbLessSessions, &dblessSession)

	dbless := dbless.DbLess{
		DbLessSessions: DbLessSessions,
		SerializerType: &serializerType,
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
