package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Storage Config

func InitConfigFile(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		// todo

		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	info, err := file.Stat()
	if err != nil {
		// todo

		return
	}

	buffer := make([]byte, info.Size())

	_, err = file.Read(buffer)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(buffer, &Storage); err != nil {
		return
	}

	return
}

type Config struct {
	AppName string `json:"app_name" yaml:"app_name"`
	Http    struct {
		Port int64 `json:"port" yaml:"port"`
	} `json:"http" yaml:"http"`
	Grpc struct {
		Port int64 `json:"port" yaml:"port"`
	} `json:"grpc" yaml:"grpc"`
	Db struct {
		DbUrl  string `json:"dburl" yaml:"dburl"`
		Driver string `json:"driver" yaml:"driver"`
	} `json:"db" yaml:"db"`
}

//type Config interface {
//	Get(key string) string
//}
//
//type config struct {
//}
//
//func NewConfig() Config {
//	return &config{}
//}
//
//func (c *config) Get(key string) string {
//
//	var value string
//
//	keys := strings.Split(key, "/")
//
//	if len(keys) == 1 {
//		val := storage[key]
//
//		switch val.(type) {
//		case string:
//			value = val.(string)
//		case int:
//			value = strconv.Itoa(val.(int))
//		case float64:
//			value = strconv.FormatFloat(val.(float64), 'E', 64, 10)
//		case map[string]interface{}:
//			value = ""
//		}
//
//		return value
//	}
//
//	path := ""
//
//	for _, k := range keys {
//		//val := storage[k]
//		//if val == nil {
//		//	break
//		//}
//		path += "/" + k
//		fmt.Println(path)
//		fmt.Println(isDeep(path, storage))
//	}
//
//	return value
//}
//
//func isDeep(key string, sto map[string]interface{}) bool {
//	val, ok := sto[key]
//	if !ok {
//		return false
//	}
//
//	switch val.(type) {
//	case map[string]interface{}:
//		return true
//	}
//
//	return false
//}
