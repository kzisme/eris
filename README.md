# eris - IRC Server / Daemon written in Go

> This project and repository is based off of [ergonomadic](https://github.com/edmund-huber/ergonomadic)
> and much of my original contributions were made in my [fork of ergonomadic](https://github.com/prologic/ergonomadic)
> but the upstream project was ultimately shutdown.
> 
> This repository intends to create a new history and improve upon prior work.

----

> In philosophy and rhetoric, eristic (from Eris, the ancient Greek goddess
> of chaos, strife, and discord) refers to argument that aims to successfully
> dispute another's argument, rather than searching for truth. According to T.H.

From [Eris](https://en.wikipedia.org/wiki/Eris_(mythology))
and [Eristic](https://en.wikipedia.org/wiki/Eristic)

The connotation here is that IRC (*Internet Relay Chat*) is a place of chaos,
strife and discord. IRC is a place where you argue and get into arguments for
the sake of argument.

So `eris` is an IRC daemon written from scratch in Go to factiliate discord
and have arguments for the sake of argument!

Pull requests and issues are welcome.

Discussion at:

* /server irc.mills.io +6697 (*use TLS/SSL*)
* /join #lobby

Or (*not recommended*)P

* /server irc.mills.io (*default port 6667, non-TLS)
* /join #lobby

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
* Secure channels (+Z)

## Installation

```#!bash
$ go install github.com/prologic/eris
$ eris --help
```

## Configuration

See the example [ircd.yml](ircd.yml). Passwords are base64-encoded
bcrypted byte strings. You can generate them with the `mkpasswd` tool
from [prologic/mkpasswd](https://github.com/prologic/mkpasswd):

```#!bash
$ go install github.com/prologic/mkpasswd
$ mkpasswd
```

## Deployment

To run simply run the `eris` binary (*assuming a `ircd.yml` in the current directory*):

```#!bash
$ eris
```

Or you can deploy with [Docker](https://www.docker.com) using the prebuilt [prologic/eris](https://hub.docker.com/r/prologic/eris/):

```#!bash
docker run -d -p 6667:6667 -p 6697:6697 prologic/eris
```

You may want to customize the configuration however and create your own image based off of this; or deploy with `docker stack deploy` on a [Docker Swarm](https://docs.docker.com/engine/swarm/) clsuter like this:

```#!bash
$ docker stack deploy -c docker-compose.yml eris
```

Which assumes a `ircd.yml` coniguration fiel int he current directory which Docker will use to distribute as the configuration. The `docker-compose.yml` (*Docker Stackfile*) is available at the root of this repository.

## Related Proejcts

There are a number of supported accompanying services that are being developed alongside Eris:

* [Soter](https://github.com/prologic/soter) -- An IRC Bot that persists channel modes and topics.
* [Cadmus](https://github.com/prologic/cadmus) -- An IRC Bot that logs channels and provides an interface for viewing and searching logs (*Coming soon...*)

## License

eris is licensed under the MIT License.
