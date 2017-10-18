package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/docopt/docopt-go"
	"github.com/prologic/ircd/irc"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	version := irc.SEM_VER
	usage := `ircd.
Usage:
	ircd initdb [--conf <filename>]
	ircd upgradedb [--conf <filename>]
	ircd genpasswd [--conf <filename>]
	ircd run [--conf <filename>]
	ircd -h | --help
	ircd --version
Options:
	--conf <filename>  Configuration file to use [default: ircd.yaml].
	-h --help          Show this screen.
	--version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, version, false)

	// Special case -- We do not need to load the config file here
	if arguments["genpasswd"].(bool) {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
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

	if arguments["initdb"].(bool) {
		irc.InitDB(config.Server.Database)
		log.Println("database initialized: ", config.Server.Database)
	} else if arguments["upgradedb"].(bool) {
		irc.UpgradeDB(config.Server.Database)
		log.Println("database upgraded: ", config.Server.Database)
	} else if arguments["run"].(bool) {
		// Create database if it isn't already created
		if _, err := os.Stat(config.Server.Database); err != nil {
			if os.IsNotExist(err) {
				irc.InitDB(config.Server.Database)
				log.Println("database initialized: ", config.Server.Database)
			}
		}

		irc.Log.SetLevel(config.Server.Log)
		server := irc.NewServer(config)
		log.Println(irc.SEM_VER, "running")
		defer log.Println(irc.SEM_VER, "exiting")
		server.Run()
	}
}
