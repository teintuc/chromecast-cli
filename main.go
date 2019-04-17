package main


import (
	"os"

	"github.com/teintuc/chromecast-cli/commands"
)

func main() {
	commands.Run(os.Args[1:])
}
