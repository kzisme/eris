# ircd - IRC Daemon

> This project and repository is based off of [ergonomadic](https://github.com/edmund-huber/ergonomadic)
> and much of my original contributions were made in my [fork of ergonomadic](https://github.com/prologic/ergonomadic)
> but the upstream project was ultimately shutdown.
> 
> This repository intends to create a new history and improve upon prior work.

----

ircd is an IRC daemon written from scratch in Go.
Pull requests and issues are welcome.

Discussion at:
* host/port: irc.mills.io:6697 (*use SSL*)
* #lobby

## Features

* follows the RFCs where possible
* UTF-8 nick and channel names
* [yaml](http://yaml.org/) configuration
* server password (PASS command)
* channels with most standard modes
* IRC operators (OPER command)
* passwords stored in [bcrypt][go-crypto] format
* messages are queued in the same order to all connected clients
* SSL/TLS support
* Simple IRC operator privileges (*overrides most things*)
* Secure connection tracking (+z) and SecureOnly user mode (+Z)

## Installation

```#!bash
$ go install github.com/prologic/ircd
$ ircd --help
```

## Configuration

See the example [ircd.yml](ircd.yml). Passwords are base64-encoded
bcrypted byte strings. You can generate them with the `genpasswd` subcommand.

```#!bash
$ ircd genpasswd
```

## Running the server

```#!bash
$ ircd run
```

## Credits

* Jeremy Latt, creator, <https://github.com/jlatt>
* Edmund Huber, maintainer, <https://github.com/edmund-huber>
* Niels Freier, added WebSocket support, <https://github.com/stumpyfr>
* apologies to anyone I forgot.
