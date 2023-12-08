package util

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type ProcessName struct {
    Name string   `yaml:"name"`
    Comm []string `yaml:"comm"`
}

type Config struct {
    ProcessNames []ProcessName `yaml:"process_names"`
}

func ModifyProcessNames(processNames []ProcessName, newName string, commands ...string) []ProcessName {
    for i := range processNames {
        processNames[i].Name = newName
        processNames[i].Comm = append(processNames[i].Comm, commands...)
    }
    return processNames
}

func ReadConfigFromFile(filename string) (Config, error) {
    if filename == "" {
        // Define a default ProcessName
        defaultProcessName := ProcessName{
            Name: "{{.ExeBase}}:{{.Username}}",
            Comm: []string{},
        }

        // Define a default Config
        defaultConfig := Config{
            ProcessNames: []ProcessName{defaultProcessName},
        }

        return defaultConfig, nil
    }

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return Config{}, err
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return Config{}, err
    }

    return config, nil
}

func WriteConfigToFile(config Config, filename string) error {
    newYaml, err := yaml.Marshal(&config)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filename, newYaml, 0644)
    if err != nil {
        return err
    }

    return nil
}