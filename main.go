package main

import (
	"fmt"
	"os"
    // "os/exec"

	"github.com/drone/drone-plugin-go/plugin"
)

type Fabric struct {
	Commands  string `json:"commands"`
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

	fabfile := fmt.Sprintf("%s/fabfile.py", workspace.Path)

	fmt.Println(fabfile)

	// fmt.Println(vargs.Commands)
	// o, err := exec.Command("fab", vargs.Commands).Output()
	// fmt.Println(o)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(2)
	// }
}
