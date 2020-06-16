# IhomeWeb Service

This is the IhomeWeb service

Generated with

```
micro new ihome/ihomeWeb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.ihomeWeb
- Type: web
- Alias: ihomeWeb

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ihomeWeb-web
```

Build a docker image
```
make docker
```