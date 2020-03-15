package serializer

type Serializer interface {
	Serialize(map[string]interface{}) ([]byte, error)
	Deserialize([]byte) (map[string]interface{}, error)
}
