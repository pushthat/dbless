package dblessutils

import (
	"encoding/base64"
	"errors"
)

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
