package config

import (
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

func ReadConfig() Config {
    filename, _ := filepath.Abs("config.yml")
    yamlFile, err := os.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    return config
}

type Config struct {
    ApiKey         string `yaml:"apiKey"`
    CompanyID      int64  `yaml:"companyID"`
    Cookie         string `yaml:"cookie"`
    OrganizationID int64  `yaml:"organizationID"`
}
