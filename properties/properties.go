package properties

import (
	"encoding/json"
	"os"
)

var (
	Props *properties = &properties{}
)

type properties struct {
	BucketName     string `json:"bucketName"`
	AccessKey      string `json:"accessKey"`
	SecretKey      string `json:"secretKey"`
	LoggerMode     string `json:"loggerMode"`
	LogApplication string `json:"logApplication"`
	LogDirectory   string `json:"logDirectory"`
}

func Load(path string) error {
	return loadConfig(path)
}
func loadConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &Props)

}
