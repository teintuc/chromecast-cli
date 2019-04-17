package discover

import (
  "log"
  "io/ioutil"

  "github.com/hashicorp/mdns"
)

const chromecastMdns = "_googlecast._tcp"

func Lookup() []*mdns.ServiceEntry {
  // Disabling logs to prevent logs from the third party library
  log.SetOutput(ioutil.Discard)

  chromecasts := []*mdns.ServiceEntry{}

  // Make a channel for results and start listening
  entriesCh := make(chan *mdns.ServiceEntry, 4)
  go func() {
      for entry := range entriesCh {
          chromecasts = append(chromecasts, entry)
      }
  }()

  // Start the lookup
  mdns.Lookup(chromecastMdns, entriesCh)
  close(entriesCh)

  return chromecasts
}