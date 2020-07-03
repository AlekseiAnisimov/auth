package auth

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// DbConfig структура хранит конфиги для подключения к БД
type DbConfig struct {
	Development struct {
		Dialect    string
		Datasource string
	}
}

var dbConfigFile = "dbconfig.yml"

// GetDbParamsFromYaml получение из конфига параметров к БД
func (dbconf *DbConfig) GetDbParamsFromYaml() error {
	fopen, err := ioutil.ReadFile(dbConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fopen, &dbconf)
	if err != nil {
		return err
	}

	return nil
}
