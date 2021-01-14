package decode

import "testing"

type TestData struct {
	Name  string `yaml:"name"`
	Value int    `yaml:"value"`
}

func TestYaml(t *testing.T) {
	var data TestData

	if err := Yaml("test.yaml", &data); err != nil {
		t.Error(err)
	}
	if data.Name != "test" || data.Value != 1 {
		t.Error("값이 맞지 않습니다.")
	}
}
