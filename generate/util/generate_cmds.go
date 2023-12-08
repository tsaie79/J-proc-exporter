package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)


func GetPGIDFromFile(filename string) (int, error) {
	// Read the file
	pgidBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	// Convert the contents to an integer
	pgid, err := strconv.Atoi(strings.TrimSpace(string(pgidBytes)))
	if err != nil {
		return 0, err
	}

	return pgid, nil
}

func GetCommandsFromPGID(procPath string, pgid int) ([]string, error) {
	var commands []string

	// Read all directories in /proc
	files, err := ioutil.ReadDir(procPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// If the file is a directory and its name is a number (i.e., it's a process ID)
		if file.IsDir() && isNumeric(file.Name()) {
			// Read the stat file for the process
			stat, err := ioutil.ReadFile(procPath + "/" + file.Name() + "/stat")
			if err != nil {
				continue
			}

			// Split the stat file into fields
			fields := strings.Fields(string(stat))

			// If the process's PGID matches the given PGID
			if len(fields) > 4 && fields[4] == strconv.Itoa(pgid) {
				// The command is the second field in the stat file, remove parentheses
				command := strings.Trim(fields[1], "()")
				commands = append(commands, command)
			}
		}
	}

	return commands, nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}