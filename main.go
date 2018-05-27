//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"
	"path"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "handbrake-portable"
	Papp.Name = "HandBrake"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	Papp.Process = PathJoin(Papp.AppPath, "HandBrake.exe")
	Papp.Args = []string{}
	Papp.WorkingDir = Papp.AppPath

	Launch(os.Args[1:])

	// Remove HandBrake Team folder
	handbrakeTeamPath := path.Join(os.Getenv("APPDATA"), "HandBrake Team")
	if _, err := os.Stat(handbrakeTeamPath); err == nil {
		os.RemoveAll(handbrakeTeamPath)
	}
}
