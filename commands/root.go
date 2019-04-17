package commands

import (
	"fmt"

	Applog "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("chromecast", "Chromecast cli.")
	debug      = app.Flag("debug", "Enable debug mode").Default("false").Bool()
)

//Buildstamp is used for storing the timestamp of the build
var Buildstamp = "Not set"

//Githash is used for storing the commit hash of the build
var Githash = "Not set"

// Version is used to store the tagged version of the build
var Version = "Not set"

func init() {
	// Configure the app
	app.DefaultEnvars()
	app.Version(getVersion())
	app.PreAction(setUpLog)

	getDiscoverCommand(app)
}

func Run(args []string) {
	kingpin.MustParse(app.Parse(args))
}

func setUpLog(c *kingpin.ParseContext) error {
	if *debug {
		Applog.SetLevel(Applog.DebugLevel)
	}

	return nil
}

func getVersion() string {
	return fmt.Sprintf("Version: %s\nBuildtime: %s\nGitHash: %s", Version, Buildstamp, Githash)
}
