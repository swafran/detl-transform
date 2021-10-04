package maps

import (
	"io/ioutil"

	"gitlab.com/detl/detl-common"
	"gopkg.in/yaml.v2"
)

type YamlMappingEl struct {
	Eltype     string
	Omit       bool
	Filters    []string
	Targets    []string
	Value      string
	Num        float64
	Boolean    bool
	Recurse    int
	Append     []string
	Prepend    []string
	AppendObj  map[string]interface{}
	PrependObj map[string]interface{}
}

type YamlMapping map[string]YamlMappingEl

// GetYamlMapping returns YamlMapping object from yaml file
func GetYamlMapping(settingsName string) YamlMapping {
	var mapping YamlMapping

	fileData, err := ioutil.ReadFile(settingsName + ".yaml")
	err2 := yaml.Unmarshal(fileData, &mapping)

	detl.Check(err)
	detl.Check(err2)

	return mapping
}
