package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/docopt/docopt-go"
	"github.com/prologic/ircd/irc"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	version := irc.FullVersion()
	usage := `ircd.
Usage:
	ircd genpasswd [--conf <filename>]
	ircd run [--conf <filename>]
	ircd -h | --help
	ircd -v | --version
Options:
	-c --conf <filename>  Configuration file to use [default: ircd.yml].
	-h --help          Show this screen.
	-v --version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, version, false)

	// Special case -- We do not need to load the config file here
	if arguments["genpasswd"].(bool) {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal("Error reading password:", err.Error())
		}
		password := string(bytePassword)
		encoded, err := irc.GenerateEncodedPassword(password)
		if err != nil {
			log.Fatalln("encoding error:", err)
		}
		fmt.Print("\n")
		fmt.Println(encoded)
		return
	}

	configfile := arguments["--conf"].(string)
	config, err := irc.LoadConfig(configfile)
	if err != nil {
		log.Fatal("Config file did not load successfully:", err.Error())
	}

	if arguments["run"].(bool) {
		irc.Log.SetLevel(config.Server.Log)
		server := irc.NewServer(config)
		log.Println(irc.FullVersion(), "running")
		defer log.Println(irc.FullVersion(), "exiting")
		server.Run()
	}
}