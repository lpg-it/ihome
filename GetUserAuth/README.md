# GetUserAuth Service

This is the GetUserAuth service

Generated with

```
micro new ihome/GetUserAuth --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetUserAuth
- Type: srv
- Alias: GetUserAuth

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
./GetUserAuth-srv
```

Build a docker image
```
make docker
```