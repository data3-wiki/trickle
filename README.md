# Trickle

<img width="128px" src="https://raw.githubusercontent.com/data3-wiki/static-assets/main/Trickle-Logo.png">

Trickle is an open source no-code framework for spinning up web3 data services painlessly. Focus on dApp logic rather than reinventing the blockchain data plumbling wheel.

## Goals

It is designed with the following goals in mind:

- Batteries Included: Easy to use for simple use cases with minimal set up. No need to set up a complicated deployment if your data fits on a single node.
- Extensible: Modular design to accomodate different L1 chains, databases, etc.
- Community First: Encourage collaboration on data tooling in the web3 community through open source with a permissive license (Apache).

## Features

- Schema-Driven: Read APIs for indexed blockchain data automatically generated based on schema definition (e.g. Solana Anchor Framework's `idl.json`).
- Web UI: Currently supports Swagger UI. Developer-friendly interface for API use.
- L1 Integrations: Currently supports Solana. Integration with more L1 chains on the roadmap. No need to bother learning to use the RPC Node APIs directly.
- Database Integrations: Currently supports SQLite. Integrations with other databases on the roadmap.

## Usage

```
go run main.go -config ./test/config.yaml
```
