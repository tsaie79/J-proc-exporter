package main

import (
    "fmt"
    "os"
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

func addCommandsToProcessNames(processNames []ProcessName, commands ...string) []ProcessName {
    for i := range processNames {
        processNames[i].Comm = append(processNames[i].Comm, commands...)
    }
    return processNames
}

func readConfigFromFile(filename string) (Config, error) {
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

func writeConfigToFile(config Config, filename string) error {
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

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: <program> <input_yaml_file> <output_yaml_file> <commands...>")
        os.Exit(1)
    }

    inputYamlFile := os.Args[1]
    outputYamlFile := os.Args[2]
    commands := os.Args[3:]

    // Read config from file
    config, err := readConfigFromFile(inputYamlFile)
    if err != nil {
        panic(err)
    }

    // Add commands to comm of each process_name
    config.ProcessNames = addCommandsToProcessNames(config.ProcessNames, commands...)

    // Print the config to verify it was read correctly
    fmt.Printf("%+v\n", config)

    // Write the updated config to a file
    err = writeConfigToFile(config, outputYamlFile)
    if err != nil {
        panic(err)
    }
}