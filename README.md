# mediateq

mediateq is a file storage REST API microservice that allows users to upload and download files, as well as resize images on the fly.

![CI](https://github.com/behouba/mediateq/actions/workflows/main.yml/badge.svg)

## Prerequisites

- libvips 8.3+ (8.8+ recommended)
- PostgreSQL 14+

## Installation

To install mediateq, clone the project repository:

```bash
git clone https://github.com/behouba/mediateq.git
```

Mediateq depend on libvips. Run the following script from [bimg](https://github.com/h2non/bimg) as sudo (supports OSX, Debian/Ubuntu, Redhat, Fedora, Amazon Linux):

```bash
curl -s https://raw.githubusercontent.com/h2non/bimg/master/preinstall.sh | sudo bash -
```

The install script requires curl and pkg-config

You can then build mediateq by running the `build.sh` script:

```bash
./build.sh
```

You will also need to setup a database. Currently mediateq only support postgreSQL.

Read the instructions to setup a postgreSQL database [here](database/postgres/README.md)

To run mediateq, use the following command:

```bash
./bin/mediateq -config=mediateq.yaml
```

You can also run tests for mediateq by using the run_tests.sh script:

```bash
./run_tests.sh
```

### API specification

The API specification of mediateq can be found [here](docs/mediateq-0.0.1.yaml)

### Docker container

TODO
