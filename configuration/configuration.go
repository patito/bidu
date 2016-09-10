package configuration

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration stores general configuration attributes
type Configuration struct {
	Database struct {
		Host     string
		Username string
		Password string
		Port     int
		DbName   string
	}
}

// New creates a new instance of Configuration
func New(filename string) (Configuration, error) {
	var conf Configuration

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return Configuration{}, err
	}

	return conf, nil
}

// StringConnection returns postgres string connection
func (conf Configuration) StringConnection() string {
	db := conf.Database
	return fmt.Sprintf("postgres://%s:%s@%s:%d?dbname=%s&sslmode=disable",
		db.Username, db.Password, db.Host, db.Port, db.DbName)
}
