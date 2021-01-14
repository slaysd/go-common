package decode

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Yaml Yaml 파일을 Decode 한다.
func Yaml(filePath string, configuration interface{}) error {
	if filePath == "" {
		return errors.New("invalid filePath")
	}
	if configuration == nil {
		return errors.New("empty configuration")
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, configuration)
	if err != nil {
		return err
	}
	return nil
}
