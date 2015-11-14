package main

import (
	"fmt"
	"os"
    "os/exec"

	"github.com/drone/drone-plugin-go/plugin"
)

type Bash struct {
	Commands  []string `json:"commands"`
}

func main() {
	repo := plugin.Repo{}
	build := plugin.Build{}
	system := plugin.System{}
	vargs := Bash{}

	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("system", &system)
	plugin.Param("vargs", &vargs)

	// parse the parameters
	if err := plugin.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

    for _, c := range vargs.Commands {
        _, err := exec.Command("sh", c).Output()
        if err != nil {
            fmt.Println(err.Error())
    		os.Exit(1)
        }
    }

}
