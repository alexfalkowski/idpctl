[![CircleCI](https://circleci.com/gh/alexfalkowski/idpctl.svg?style=svg)](https://circleci.com/gh/alexfalkowski/idpctl)
[![codecov](https://codecov.io/gh/alexfalkowski/idpctl/graph/badge.svg?token=QSRFU8VNST)](https://codecov.io/gh/alexfalkowski/idpctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/idpctl)](https://goreportcard.com/report/github.com/alexfalkowski/idpctl)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/idpctl.svg)](https://pkg.go.dev/github.com/alexfalkowski/idpctl)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)


# Internal Developer Platform

An [Internal Developer Platform](https://internaldeveloperplatform.org/what-is-an-internal-developer-platform/) (IDP) is built by a platform team to build golden paths and enable developer self-service.

### Why a client?

Internal developer portals serve as the interface through which developers can discover and access internal developer platform capabilities.

Separating the client from the service allows us to make sure they move independently.

## Client

The client is broken down to different commands.

### Pipeline

```shell
❯ ./idpctl pipeline --help
Manage pipelines.

Usage:
  idpctl pipeline [flags]

Flags:
  -c, --create string    create a pipeline (path to file)
  -d, --delete string    delete a pipeline (the id of the pipeline)
  -g, --get string       retrieve a pipeline (the id of the pipeline)
  -h, --help             help for pipeline
  -t, --trigger string   trigger a pipeline (the id of the pipeline)
  -u, --update string    update a pipeline (id of the pipeline and the path to file, e.g. id:path)

Global Flags:
  -i, --input string   input config location (format kind:location)
```

Check out the [Makefile](Makefile) to see how you can use it from there.

## Design

The design is heavily influenced by https://github.com/alexfalkowski/servicectl.

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- Ruby
- Golang

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Let's get you setup.

#### Submodules

We need to get the git submodules.

```shell
git submodule sync
git submodule update --init
```

#### Application

To get the application running, do the following:

1. Let's get the dependencies:
```shell
make go-dep
```
2. Let's build the application:
```shell
make build
```
3. Let's create a pipeline:
```shell
make pipeline=pipeline create
```

#### Features

If you want to run the features, do the following:

1. Let's get the dependencies:
```shell
make ruby-dep
```
2. Let's run the features:
```shell
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
