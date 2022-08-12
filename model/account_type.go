package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AccountType struct {
	Name          string          `yaml:"name"`
	PropertyTypes []*PropertyType `yaml:"properties"`
}

type PropertyType struct {
	Name     string   `yaml:"name"`
	DataType DataType `yaml:"data_type"`
}

func FromYamlFile(filepath string) (*AccountType, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	accountType := &AccountType{}
	err = yaml.Unmarshal(buf, &accountType)
	if err != nil {
		return nil, err
	}

	// TODO: Validate data types.

	return accountType, nil
}