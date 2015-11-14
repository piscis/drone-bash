package main

import (
	"fmt"
	"os"
    "os/exec"

	"github.com/drone/drone-plugin-go/plugin"
)

type Fabric struct {
	Commands  string `json:"commands"`
}

func main() {
	repo := plugin.Repo{}
	build := plugin.Build{}
	system := plugin.System{}
	vargs := Fabric{}

	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("system", &system)
	plugin.Param("vargs", &vargs)

	// parse the parameters
	if err := plugin.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, err := exec.Command("fab", vargs.Commands).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
