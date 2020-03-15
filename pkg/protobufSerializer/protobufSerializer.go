package protobufSerializer

// import (
// 	"log"
// 	"reflect"

// 	"github.com/golang/protobuf/proto"
// )

// type ProtobufSerializer struct {
// 	Structure proto.Message
// }

// type AnythingForYou struct {
// 	Anything *proto.Any `protobuf:"bytes,1,opt,name=anything" json:"anything,omitempty"`
// }

// func (serializer ProtobufSerializer) Serialize(data []byte) ([]byte, error) {
// 	protoObject := reflect.ValueOf(obj).Interface().(reflect.ValueOf(serializer.Structure))

// 	data, err := proto.Marshal(protoObject)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// func (serializer ProtobufSerializer) Deserialize(data []byte) (interface{}, error) {
// 	obj := reflect.ValueOf(serializer.Structure).Interface().(proto.Message)

// 	err := proto.Unmarshal(data, obj)
// 	if err != nil {
// 		log.Fatal("unmarshaling error: ", err)
// 	}
// 	return obj, nil
// }
