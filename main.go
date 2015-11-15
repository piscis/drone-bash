package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	"os/user"
	"log"
	"io/ioutil"
	"path/filepath"

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

	if err := writeKey(&workspace); err != nil {
		log.Println("Unable to write private key")
		log.Println(err)
		os.Exit(1)
	}

	if err := os.Chdir(workspace.Path); err != nil {
		log.Println("Unable to dc into workspace.Path")
		os.Exit(1)
	}

	for _, c := range vargs.Commands {
		fabArgs := strings.Split(c, " ")
		c := exec.Command("fab", fabArgs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			log.Println(err)
			os.Exit(101)
		}
	}
}

func writeKey(in *plugin.Workspace) error {
	if in.Keys == nil || len(in.Keys.Private) == 0 {
		return nil
	}
	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	sshpath := filepath.Join(home, ".ssh")
	if err := os.MkdirAll(sshpath, 0700); err != nil {
		return err
	}
	confpath := filepath.Join(sshpath, "config")
	privpath := filepath.Join(sshpath, "id_rsa")
	ioutil.WriteFile(confpath, []byte("StrictHostKeyChecking no\n"), 0700)
	return ioutil.WriteFile(privpath, []byte(in.Keys.Private), 0600)
}
