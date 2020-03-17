package dbless

import (
	"bytes"

	"github.com/pushthat/dbless/internal/dblessutils"
	"github.com/pushthat/dbless/pkg/jsonSerializer"
	"github.com/pushthat/dbless/pkg/serializer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type DbLess struct {
	S3Bucket       *string
	SerializerType *string
	Index          *string
	Prefix         *string
	S3Session      *session.Session
	serializer     serializer.Serializer
	s3Uploader     *s3manager.Uploader
	s3Downloader   *s3manager.Downloader
}

func (dbless *DbLess) Init() error {
	switch *dbless.SerializerType {
	case "json":
		serializer := jsonSerializer.JsonSerializer{}
		dbless.serializer = &serializer
	}

	uploader := s3manager.NewUploader(dbless.S3Session)
	dbless.s3Uploader = uploader
	downloader := s3manager.NewDownloader(dbless.S3Session)
	dbless.s3Downloader = downloader
	return nil
}

// TODO error management
func (dbless *DbLess) Save(obj map[string]interface{}) error {
	keyName, err := dblessutils.GetObjectIndex(obj, dbless.Index)
	if err != nil {
		return err
	}

	keyName = *dbless.Prefix + keyName
	data, err := dbless.serializer.Serialize(obj)
	if err != nil {
		return err
	}
	upParams := &s3manager.UploadInput{
		Bucket: dbless.S3Bucket,
		Key:    &keyName,
		Body:   bytes.NewReader(data),
	}
	_, err = dbless.s3Uploader.Upload(upParams)
	if err != nil {
		return err
	}
	return nil
}

func (dbless *DbLess) Load(obj map[string]interface{}) (map[string]interface{}, error) {
	keyName, _ := dblessutils.GetObjectIndex(obj, dbless.Index)
	buf := aws.NewWriteAtBuffer([]byte{})
	keyName = *dbless.Prefix + keyName

	downloadParam := s3.GetObjectInput{
		Bucket: dbless.S3Bucket,
		Key:    &keyName,
	}

	dbless.s3Downloader.Download(buf, &downloadParam)
	data, err := dbless.serializer.Deserialize(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dbless *DbLess) Delete(obj interface{}) {

}
