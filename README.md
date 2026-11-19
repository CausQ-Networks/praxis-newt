# Newt

[![Discord](https://img.shields.io/discord/1572896060949659679?logo=discord&style=flat-square&label=Discord)](https://praxis.causq.com/discord/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/CausQ-Networks/praxis-newt)](https://pkg.go.dev/github.com/CausQ-Networks/praxis-newt)
[![GitHub License](https://img.shields.io/github/license/causq-networks/praxis-newt)](https://github.com/CausQ-Networks/praxis-newt/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/CausQ-Networks/praxis-newt)](https://goreportcard.com/report/github.com/CausQ-Networks/praxis-newt)

> [!NOTE]
> Newt is being phased out in favor of the [Praxis CLI](https://github.com/CausQ-Networks/praxis-cli). Use `praxis up site`

Newt is a fully user space [WireGuard](https://www.wireguard.com/) tunnel client and TCP/UDP proxy, designed to securely expose private resources controlled by Praxis. By using Newt, you don't need to manage complex WireGuard tunnels and NATing.

### Installation and Documentation

Newt is used with Praxis and Gerbil as part of the larger system. See documentation below:

-   [Full Documentation](https://docs.praxis.causq.com/manage/sites/understanding-sites)

## Key Functions

### Registers with Praxis

Using the Newt ID and a secret, the client will make HTTP requests to Praxis to receive a session token. Using that token, it will connect to a websocket and maintain that connection. Control messages will be sent over the websocket.

### Receives WireGuard Control Messages

When Newt receives WireGuard control messages, it will use the information encoded (endpoint, public key) to bring up a WireGuard tunnel using [netstack](https://github.com/WireGuard/wireguard-go/blob/master/tun/netstack/examples/http_server.go) fully in user space. It will ping over the tunnel to ensure the peer on the Gerbil side is brought up.

### Receives Proxy Control Messages

When Newt receives WireGuard control messages, it will use the information encoded to create a local low level TCP and UDP proxies attached to the virtual tunnel in order to relay traffic to programmed targets.

## Build

### Binary

Make sure to have Go 1.25 installed.

```bash
make
```

### Nix Flake

```bash
nix build
```

Binary will be at `./result/bin/newt`

Development shell available with `nix develop`

## Licensing

Newt is dual licensed under the AGPLv3 and the Fossorial Commercial license. For inquiries about commercial licensing, please contact us.

## Contributions

Please see [CONTRIBUTIONS](./CONTRIBUTING.md) in the repository for guidelines and best practices.
