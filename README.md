# Trickle

<div align="center">
    <img width="128px" src="https://raw.githubusercontent.com/data3-wiki/static-assets/main/Trickle-Logo.png">
</div>

Trickle is an open source no-code framework for spinning up web3 data services painlessly. Focus on dApp logic rather than reinventing the blockchain data plumbling wheel. This project is supported by the [Data3 Wiki](https://www.data3.wiki/).

## Features

- **Schema-Driven Indexing**: APIs for indexed blockchain data are automatically generated based on schema definition (e.g. Solana Anchor Framework's `idl.json`).
- **Web UI**: Bundled with Swagger UI. Developer-friendly interface for API use.
- **L1 Integrations**: Currently supports pulling data from Solana. Integration with more L1 chains on the roadmap. No need to bother learning to use the different L1 RPC node APIs directly.
- **Database Integrations**: Currently supports storing data in SQLite. Integrations with other databases on the roadmap.

## Usage (< 5 minutes)

1. Run the service.

```
go run main.go -config ./test/config.yaml
```

## Goals

It is designed with the following goals in mind:

- **Batteries Included**: Easy to use for simple use cases with minimal set up. No need to set up a complicated deployment if your data fits on a single node.
- **Extensible**: Modular design to accomodate different L1 chains, databases, etc.
- **Community First**: Encourage collaboration on data tooling in the web3 community through open source with a permissive license (Apache).

### Roadmap

- [ ] Data Loading
    - [ ] Solana
        - [x] Account data loading using getProgramAccounts
        - [ ] Account data loading using WebSocket
        - [ ] Instruction data loading
- [ ] Data Decoding
    - [ ] Solana
        - [ ] Anchor
            - [x] Accounts
            - [ ] Instructions
        - [ ] Custom Decoders
- [ ] Database Integrations
    - [ ] SQLite
        - [x] Primivite types
        - [ ] Index Management
    - [ ] PostgreSQL