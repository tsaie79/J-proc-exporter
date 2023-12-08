package main

import (
	"fmt"
	"os"
	"jprocexporter/util" // replace "yourproject" with the name of your project
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: <program> <input_yaml_file> <output_yaml_file> <new_name> <commands...>")
		os.Exit(1)
	}

	inputYamlFile := os.Args[1]
	outputYamlFile := os.Args[2]
	newName := os.Args[3]
	commands := os.Args[4:]

	// Read config from file
	config, err := util.ReadConfigFromFile(inputYamlFile)
	if err != nil {
		panic(err)
	}

	// Modify process names and add commands
	config.ProcessNames = util.ModifyProcessNames(config.ProcessNames, newName, commands...)

	// Print the config to verify it was read correctly
	fmt.Printf("%+v\n", config)

	// Write the updated config to a file
	err = util.WriteConfigToFile(config, outputYamlFile)
	if err != nil {
		panic(err)
	}
}