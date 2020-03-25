package dbless

import (
	"bytes"
	"errors"
	"log"

	"github.com/pushthat/dbless/internal/dblessutils"
	"github.com/pushthat/dbless/pkg/jsonSerializer"
	"github.com/pushthat/dbless/pkg/serializer"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type DbLessSession struct {
	S3Bucket  *string
	Index     *string
	Prefix    *string
	S3Session *session.Session
}

type DbLess struct {
	DbLessSessions []*DbLessSession
	SerializerType *string
	serializer     serializer.Serializer
}

func (dbless *DbLess) Init() error {
	switch *dbless.SerializerType {
	case "json":
		serializer := jsonSerializer.JsonSerializer{}
		dbless.serializer = &serializer
	}
	return nil
}

func (dbless *DbLess) Save(obj map[string]interface{}) error {
	for _, dblessSession := range dbless.DbLessSessions {
		keyName, err := dblessutils.GetObjectIndex(obj, dblessSession.Index)
		if err != nil {
			return err
		}

		keyName = *dblessSession.Prefix + keyName
		data, err := dbless.serializer.Serialize(obj)
		if err != nil {
			return err
		}
		upParams := &s3manager.UploadInput{
			Bucket: dblessSession.S3Bucket,
			Key:    &keyName,
			Body:   bytes.NewReader(data),
		}
		err = dblessutils.UploadS3File(upParams, dblessSession.S3Session)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dbless *DbLess) Load(obj map[string]interface{}) (map[string]interface{}, error) {
	for _, dblessSession := range dbless.DbLessSessions {
		keyName, _ := dblessutils.GetObjectIndex(obj, dblessSession.Index)
		keyName = *dblessSession.Prefix + keyName

		downloadParam := s3.GetObjectInput{
			Bucket: dblessSession.S3Bucket,
			Key:    &keyName,
		}

		bytes, err := dblessutils.GetS3File(&downloadParam, dblessSession.S3Session)
		if err != nil {
			log.Println("Error while fectching data : " + err.Error())
			continue
		}
		data, err := dbless.serializer.Deserialize(bytes)
		if err != nil {
			log.Println("Error while Deserializing data : " + err.Error())
			continue
		}
		return data, nil
	}
	return nil, errors.New("Can't fetch data from bucket")
}
func (dbless *DbLess) Delete(obj interface{}) {

}
