package jsonSerializer

import (
	"encoding/json"
	"errors"
)

type JsonSerializer struct {
}

func tryCastToMap(obj interface{}) (map[string]interface{}, error) {
	mapValue, isMap := obj.(map[string]interface{})
	if isMap {
		return mapValue, nil
	} else {
		return nil, errors.New("Invalid object content")
	}
}

func (serializer *JsonSerializer) Serialize(data map[string]interface{}) ([]byte, error) {
	jsonContent, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}
	return jsonContent, nil
}

func (serializer *JsonSerializer) Deserialize(data []byte) (map[string]interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	objMap, err := tryCastToMap(obj)
	if err != nil {
		return nil, err
	}
	return objMap, nil
}
