package config

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

var (
	Cfg Config
)

type Config struct {
	Port       string `yaml:"port"`
	DBHost     string `yaml:"dbHost"`
	DBUser     string `yaml:"dbUser"`
	DBPassword string `yaml:"dbPassword"`
	DBSchema   string `yaml:"dbName"`
	DBPort     string `yaml:"dbPort"`
	Secret     string `yaml:"secret"`
}

func Connect() (*gorm.DB, error) {
	var cfg Config

	yamlFile, fileErr := ioutil.ReadFile("config.yaml")
	if fileErr != nil {
		return nil, fmt.Errorf("Error while loading config file: %v\n", fileErr)
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, fmt.Errorf("Problem while unmarshaling config file: %v\n", err)
	}
	Cfg = cfg

	param := "charset=utf8&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBSchema, param)

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("Error while connecting to database: %v", err)
	}
	return db, nil
}
