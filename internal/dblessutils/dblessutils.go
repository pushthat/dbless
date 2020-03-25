package dblessutils

import (
	"encoding/base64"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetS3File(downloadParam *s3.GetObjectInput, s3Session *session.Session) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})

	downloader := s3manager.NewDownloader(s3Session)
	downloader.Download(buf, downloadParam)
	return buf.Bytes(), nil
}

func UploadS3File(upParams *s3manager.UploadInput, s3Session *session.Session) error {
	uploader := s3manager.NewUploader(s3Session)
	_, err := uploader.Upload(upParams)
	if err != nil {
		return err
	}
	return nil
}

func GetObjectIndex(object map[string]interface{}, index *string) (string, error) {
	value, isKey := object[*index]
	if isKey {
		stringValue, isString := value.(string)
		if isString {
			sEnc := base64.StdEncoding.EncodeToString([]byte(stringValue))
			return sEnc, nil
		} else {
			return "", errors.New("Index does not exist in object")
		}
	} else {
		return "", errors.New("Index does not exist in object")
	}
}
