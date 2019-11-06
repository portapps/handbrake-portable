//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"path"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("handbrake-portable", "HandBrake"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "HandBrake.exe")

	defer func() {
		handbrakeTeamPath := path.Join(utl.RoamingPath(), "HandBrake Team")
		if _, err := os.Stat(handbrakeTeamPath); err == nil {
			if err := os.RemoveAll(handbrakeTeamPath); err != nil {
				Log.Error().Err(err).Msg("Cannot remove old appdata folder")
			}
		}
	}()

	app.Launch(os.Args[1:])
}
