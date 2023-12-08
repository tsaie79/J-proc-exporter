package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
	"github.com/tsaie79/J-proc-exporter/util"
)



// FILEPATH: /workspaces/J-proc-exporter/psgo/main.go
func Run(inputYamlFile, outputYamlFile, newName, cmds, pgidFile, procPath string) {
    help := flag.Bool("help", false, "Show usage information")
    flag.Parse()

    if *help || outputYamlFile == "" || (cmds == "" && pgidFile == "") {
        fmt.Println("Usage: binary [OPTIONS]")
        fmt.Println("A tool for generating YAML files for the process-exporter of Prometheus.")
        fmt.Println()
        fmt.Println("Options:")
        fmt.Println("  --inyml x.yml    Specifies the input YAML file. Optional.")
        fmt.Println("  --outyml y.yml   Specifies the output YAML file. Required.")
        fmt.Println("  --name \"xxx\"    Specifies the new name. Optional.")
        fmt.Println("  --cmds \"a b c\"  Specifies the commands. Either this or --pgid must be provided.")
        fmt.Println("  --pgid pgid.txt  Specifies the file containing the PGID. Either this or --cmds must be provided.")
        fmt.Println("  --procpath       Specifies the path to the /proc directory. Default is /proc.")
        fmt.Println("  --help           Show this help message.")
        os.Exit(1)
    }

    var commands []string
    if cmds != "" {
        commands = strings.Split(cmds, " ")
    } else {
        // Read PGID from file
        pgid, err := util.GetPGIDFromFile(pgidFile)
        if err != nil {
            panic(err)
        }

        // Get commands from PGID
        commands, err = util.GetCommandsFromPGID(procPath, pgid)
        if err != nil {
            panic(err)
        }
    }

	// Read config from file
	config, err := util.ReadConfigFromFile(inputYamlFile)
	if err != nil {
		panic(err)
	}

	// If newName is an empty string, use the default name from the config
	if newName == "" {
		newName = config.DefaultName
		// println(newName)
		fmt.Printf("%+v\n", newName)
	}

	// Modify process names and add commands
	config.ProcessNames = util.ModifyProcessNames(config.ProcessNames, newName, commands...)

	// Write config to file
	err = util.WriteConfigToFile(config, outputYamlFile)
	if err != nil {
		panic(err)
	}
	

}

func main() {
    inputYamlFile := flag.String("inyml", "", "Input YAML file")
    outputYamlFile := flag.String("outyml", "", "Output YAML file")
    newName := flag.String("name", "", "New name")
    cmds := flag.String("cmds", "", "Commands")
    pgidFile := flag.String("pgid", "", "File containing PGID")
    procPath := flag.String("procpath", "/proc", "Path to the /proc directory")

    flag.Parse()

    Run(*inputYamlFile, *outputYamlFile, *newName, *cmds, *pgidFile, *procPath)
}