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

* /server irc.mills.io:6697 (*use SSL*)
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

## Running the server

```#!bash
$ eris
```

## License

eris is licensed under the MIT License.
