package maps

import (
	"io/ioutil"

	"gitlab.com/detl/detl-common"
	"gopkg.in/yaml.v2"
)

type MappingEl struct {
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
	AppendObj  map[string]string
	PrependObj map[string]interface{}
}

type Mapping map[string]MappingEl

// GetMapping returns YamlMapping object from yaml file
func GetMapping(settingsName string) Mapping {
	var doy Mapping

	//var mapping interface{}
	var mapping Mapping

	fileData, err := ioutil.ReadFile(settingsName + ".yaml")
	err2 := yaml.Unmarshal(fileData, &mapping)

	detl.Check(err)
	detl.Check(err2)

	return doy
	//return mapping
}
