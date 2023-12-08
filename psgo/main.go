package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
	"github.com/tsaie79/J-proc-exporter/util" // replace "yourproject" with the name of your project
)

func main() {
	inputYamlFile := flag.String("inyml", "", "Input YAML file")
	outputYamlFile := flag.String("outyml", "", "Output YAML file")
	newName := flag.String("name", "", "New name")
	cmds := flag.String("cmds", "", "Commands")

	flag.Parse()

	if *outputYamlFile == "" || *newName == "" || *cmds == "" {
		fmt.Println("Usage: binary --inyml x.yml --outyml y.yml --name \"xxx\" --cmds \"a b c\"")
		os.Exit(1)
	}

	commands := strings.Split(*cmds, " ")

	// Read config from file
	config, err := util.ReadConfigFromFile(*inputYamlFile)
	if err != nil {
		panic(err)
	}

	// Modify process names and add commands
	config.ProcessNames = util.ModifyProcessNames(config.ProcessNames, *newName, commands...)

	// Print the config to verify it was read correctly
	fmt.Printf("%+v\n", config)

	// Write the updated config to a file
	err = util.WriteConfigToFile(config, *outputYamlFile)
	if err != nil {
		panic(err)
	}
}