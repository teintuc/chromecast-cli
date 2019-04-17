package commands

import (
  "fmt"

  "github.com/teintuc/chromecast-cli/pkg/discover"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type DiscoverCommand struct {}

func (dc *DiscoverCommand) run(c *kingpin.ParseContext) error {
  chromecasts := discover.Lookup()

  for key, chromecast := range chromecasts {
    fmt.Printf("%d:\t%s\n", key, chromecast.Name)
    fmt.Printf("\t%s\n", chromecast.Host)
    fmt.Printf("\t%s\n", chromecast.AddrV4)
    fmt.Printf("\t%s\n", chromecast.AddrV6)
    fmt.Printf("\t%d\n", chromecast.Port)
    fmt.Printf("\t%s\n\n", chromecast.Info)
  }
  return nil
}

func getDiscoverCommand(app *kingpin.Application) {
	dc := &DiscoverCommand{}
	app.Command("discover", "Discover Chromecasts on the network.").Action(dc.run)
}