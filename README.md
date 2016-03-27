# hero [![Build Status](https://travis-ci.org/gernest/hero.svg)](https://travis-ci.org/gernest/hero)  [![GoDoc](https://godoc.org/github.com/gernest/hero?status.svg)](https://godoc.org/github.com/gernest/hero)
hero is a feature rich oauth 2 server implementation in Go.


# Features

- [ ]  User account management
- [x]  Encrypted passwords and client secrets.
- [ ]  Client management
- [x]  oauth 2 [rfc 6749](http://tools.ietf.org/html/rfc6749) compliant
- [x]  Configurable.
- [x]  Multiple databases( postgres, mysql and foundation)
- [x]  Hot templates reload for rapid development(no need to recompile when run with `-dev` flag).
- [ ]  csrf protection
- [x]  http2 support
- [x]  https support
- [ ]  Let's encrypt integration

## Table of contents

- [ ] [Getting started]()
  - [ ] [What is hero](docs/introduction.md#what-is-hero-)
  - [x] [Installation](docs/introduction.md#features)
  - [x] [How to use]()
- [ ] [configuration](docs/config.md#configuration)
- [ ] [User management]()
- [ ] [clients management]()
- [ ] [contributing](#contributing)


# Contributing
contributions are welcome. For development please see [DEVELOPER.md](DEVELOPER.md)


# Author
Geofrey Ernest

Twitter  : [@gernesti](https://twitter.com/gernesti)



# Aknowledgement

The project [osin](https://github.com/RangelReale/osin) is the foundation from which I got insight and direction of how oauth 2 works.



# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.
