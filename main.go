package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	// "io/ioutil"

	"github.com/drone/drone-plugin-go/plugin"
)

type Fabric struct {
	Commands  []string `json:"commands"`
}

func main() {
	workspace := plugin.Workspace{}
	vargs := Fabric{}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &vargs)

	// parse the parameters
	if err := plugin.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// err := ioutil.WriteFile("/tmp/ssh_key", []byte(workspace.Keys.Public), 0644)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	fabfile := fmt.Sprintf("--fabfile=%s/fabfile.py", workspace.Path)

	for _, c := range vargs.Commands {
		command := fmt.Sprintf("%s %s", fabfile, c)
		fabArgs := strings.Split(command, " ")
		// fabArgs = append(fabArgs, "-i", "/tmp/ssh_key")

		c := exec.Command("fab", fabArgs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(101)
		}
	}
}
