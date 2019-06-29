package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	helper "github.com/shouva/dailyhelper"
)

var setting Setting

func main() {

	currentdir := helper.GetCurrentPath(false)

	err := helper.ReadConfig(currentdir+"/jsbuilder.json", &setting)
	if err != nil {
		panic(err)
	}

	for {
		runUpdate()
		// fmt.Println(time.Now())
		time.Sleep(time.Duration(setting.Delay * 1000000000))
		// fmt.Println("...end")
	}
}

func runUpdate() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover :", r)
		}
	}()
	var cmd *exec.Cmd
	if len(setting.Branch) > 0 {
		cmd = exec.Command("git", "pull", "origin", setting.Branch)
	} else {
		cmd = exec.Command("git", "pull")
	}
	cmd.Dir = setting.Path
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(b))
	if !strings.Contains(string(b), "Already") {
		cmd = exec.Command("npm", "run", "build")
		cmd.Dir = setting.Path
		fmt.Println("Proses rebuild berjalan")
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("prose reupdate selesai")
	}
}

// Setting :
type Setting struct {
	Path   string `json:"path"`
	Delay  int    `json:"delay"`
	Branch string `json:"branch"`
}
