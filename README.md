# Next js and Golang example

**trip-planner** is an opinionated *production-ready* RESTful JSON backend template written in [Go](https://golang.org/), highly integrated with [VSCode DevContainers](https://code.visualstudio.com/docs/remote/containers).


## Table of Contents

- [Next js and Golang example](#next-js-and-golang-example)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
    - [Requirements](#requirements)
    - [Quickstart](#quickstart)
    - [Visual Studio Code](#visual-studio-code)
    - [Building and testing](#building-and-testing)
    - [Running](#running)
    - [Uninstall](#uninstall)
  - [License](#license)

## Features

- Full local golang service development environment using [Docker Compose](https://docs.docker.com/compose/install/) and [VSCode devcontainers](https://code.visualstudio.com/docs/remote/containers) that just works with Linux, MacOS and Windows.
- Adheres to the project layout defined in [golang-standard/project-layout](https://github.com/golang-standards/project-layout).
- Provides database migration ([sql-migrate](https://github.com/rubenv/sql-migrate)) and models generation ([SQLBoiler](https://github.com/volatiletech/sqlboiler)) workflows for [PostgreSQL](https://www.postgresql.org/) databases.
- Integrates [IntegreSQL](https://github.com/allaboutapps/integresql) for fast, concurrent and isolated integration testing with real PostgreSQL databases.
- Auto-installs our recommended VSCode extensions for golang development.
- Integrates [go-swagger](https://github.com/go-swagger/go-swagger) for compile-time generation of `swagger.yml`, structs and request/response validation functions.
- Integrates [MailHog](https://github.com/mailhog/MailHog) for easy SMTP-based email testing.
- Integrates [SwaggerUI](https://github.com/swagger-api/swagger-ui) for live-previewing your Swagger v2 schema.
- Integrates [pgFormatter](https://github.com/darold/pgFormatter) and [vscode-pgFormatter](https://marketplace.visualstudio.com/items?itemName=bradymholt.pgformatter) for SQL formatting.
- Comes with fully implemented `auth` package, an OAuth2 RESTful JSON API ready to be extended according to your requirements.
- Implements [OAuth 2.0 Bearer Tokens](https://tools.ietf.org/html/rfc6750) and password authentication using [argon2id](https://godoc.org/github.com/alexedwards/argon2id) hashes.
- Comes with a tested mock and [FCM](https://firebase.google.com/docs/cloud-messaging) provider for sending push notifications and storing push tokens.
- CLI layer provided by [spf13/cobra](https://github.com/spf13/cobra). It's exceptionally easy to [add additional sub-commands via `cobra-cli`](https://github.com/spf13/cobra-cli/blob/main/README.md#add-commands-to-a-project).
- Comes with an initial [PostgreSQL](https://www.postgresql.org/) database structure (see [/migrations](https://github.com/allaboutapps/go-starter/tree/master/migrations)), covering:
  - auth tokens (access-, refresh-, password-reset-tokens),
  - a generic auth-related `user` model
  - an app-specific bare-bones `app_user_profile` model,
  - push notification tokens and
  - a health check sequence (for performing writeable checks).
- API endpoints and CLI for liveness (`/-/healthy`) and readiness (`/-/ready`) probes
- Parallel jobs optimized `Makefile` and various convenience scripts (see all targets and its description via `make help`). A full rebuild only takes seconds.
- Multi-staged `Dockerfile` (`development` -> `builder` -> `app`).

### Requirements

Requires the following local setup for development:

- [Docker CE](https://docs.docker.com/install/) (19.03 or above)
- [Docker Compose](https://docs.docker.com/compose/install/) (1.25 or above)
- [VSCode Extension: Remote - Containers](https://code.visualstudio.com/docs/remote/containers) (`ms-vscode-remote.remote-containers`)

This project makes use of the [Remote - Containers extension](https://code.visualstudio.com/docs/remote/containers) provided by [Visual Studio Code](https://code.visualstudio.com/). A local installation of the Go tool-chain is **no longer required** when using this setup.

Please refer to the [official installation guide](https://code.visualstudio.com/docs/remote/containers) how this works for your host OS and head to our [FAQ: How does our VSCode setup work?](https://github.com/allaboutapps/go-starter/wiki/FAQ#how-does-our-vscode-setup-work) if you encounter issues.

### Quickstart

Create a new git repository through the GitHub template repository feature ([use this template](https://github.com/allaboutapps/go-starter/generate)). You will then start with a **single initial commit** in your own repository.

```bash
# Clone your new repository, cd into it, then easily start the docker-compose dev environment through our helper
./docker-helper.sh --up
```

You should be inside the 'service' docker container with a bash shell.

```bash
development@94242c61cf2b:/app$ # inside your container...

# Shortcut for make init, make build, make info and make test
make all

# Print all available make targets
make help
```

### Visual Studio Code

Run `CMD+SHIFT+P` `Go: Install/Update Tools` **after** attaching to the container with VSCode to auto-install all golang related vscode extensions.


### Building and testing

Other useful commands while developing your service:

```bash
development@94242c61cf2b:/app$ # inside your container...

# Print all available make targets
make help

# Shortcut for make init, make build, make info and make test
make all

# Init install/cache dependencies and install tools to bin
make init

# Rebuild only after changes to files (generate, format, build, lint)
make

# Execute all tests
make test
```

### Running

To run the service locally you may:

```bash
development@94242c61cf2b:/app$ # inside your development container...

# First ensure you have a fresh `app` executable available
make build

# Check if all requirements for becoming are met (db is available, mnt path is writeable)
app probe readiness -v

# Migrate up the database
app db migrate

# Seed the database (if you have any fixtures defined in `/internal/data/fixtures.go`)
app db seed

# Start the locally-built server
app server

# Now available at http://127.0.0.1:8080

# You may also run all the above commands in a single command
app server --probe --migrate --seed # or `app server -pms`
```

### Uninstall

Simply run `./docker-helper --destroy` in your working directory (on your host machine) to wipe all docker related traces of this project (and its volumes!).

## License

[MIT](LICENSE)
