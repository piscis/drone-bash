package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"

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

	fabfile := fmt.Sprintf("--fabfile=%s/fabfile.py", workspace.Path)

	for _, c := range vargs.Commands {
		command := fmt.Sprintf("%s %s", fabfile, c)
		fabArgs := strings.Split(command, " ")
		_, err := exec.Command("fab", fabArgs...).Output()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
	}
}
