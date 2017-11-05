package main

import (
	"fmt"
	"syscall"

	"github.com/docopt/docopt-go"
	"github.com/prologic/ircd/irc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	version := irc.FullVersion()
	usage := `ircd.
Usage:
	ircd genpasswd [--conf <filename>]
	ircd run [--conf <filename>] [ -d | --debug ]
	ircd -h | --help
	ircd -v | --version
Options:
	-c --conf <filename>  Configuration file to use [default: ircd.yml].
	-h --help          Show this screen.
	-v --version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, version, false)

	if arguments["-d"].(bool) || arguments["--debug"].(bool) {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

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
		server := irc.NewServer(config)
		log.Println(irc.FullVersion(), "running")
		defer log.Println(irc.FullVersion(), "exiting")
		server.Run()
	}
}
