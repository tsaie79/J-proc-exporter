package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"github.com/containers/psgo"
	"io/ioutil"

    "gopkg.in/yaml.v2"
)


func writeYaml() {
    data := map[string]string{
        "process_names": []map[string]string{
				"names": "{{.ExeBase}}:{{.Username}}", 
				"comm": "sshd",
				"pid": "1",
		}
	}
    

    // Marshal data to YAML
    yamlData, err := yaml.Marshal(data)
    if err != nil {
        panic(err)
    }

    // Write YAML data to file
    err = ioutil.WriteFile("data.yaml", yamlData, 0644)
    if err != nil {
        panic(err)
    }

    fmt.Println("YAML data written to file.")
}






func getComd() {
	data, err := psgo.ProcessInfoByPids([]string{"1"}, []string{"user", "pid", "ppid", "pgid", "state", "pcpu", "comm"})
	// data, err := psgo.ProcessInfoByPids([]string{"1"}, []string{"pid", "state"})

	if err != nil {
		log.Fatal(err)
	}

	cmds := []string{}
	pid := []string{}
	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	for _, d := range data {
		if d[4] == "S" {
			fmt.Fprintln(tw, strings.Join(d, "\t"))
			cmds = append(cmds, d[6])
			pid = append(pid, d[1])
		}
	}	
	tw.Flush()
	fmt.Println(cmds)
	fmt.Println(pid)
}

func main() {
	writeYaml()
}
