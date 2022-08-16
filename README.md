# Trickle

<div align="center">
    <img width="128px" src="https://raw.githubusercontent.com/data3-wiki/static-assets/main/Trickle-Logo.png">
</div>

[![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/data3_wiki.svg?style=social&label=Follow%20%40Data3_Wiki)](https://twitter.com/data3_wiki)

Trickle is an open source no-code framework for spinning up web3 data services painlessly. Focus on dApp logic rather than reinventing the blockchain data plumbing wheel. This project is supported by the [Data3 Wiki](https://www.data3.wiki/).

## Features

- **Schema-Driven Indexing**: APIs for indexed blockchain data are automatically generated based on schema definition (e.g. Solana Anchor Framework's `idl.json`).
- **Web UI**: Bundled with Swagger UI. Developer-friendly interface for API use.
- **Blockchain Integrations**: Currently supports pulling data from Solana. Integration with more blockchains on the roadmap. No need to bother learning to use the different blockchain RPC node APIs directly.
- **Database Integrations**: Currently supports storing data in SQLite. Integrations with other databases on the roadmap.

## Getting Started

0. Write a configuration file to specify the RPC node url, program id and `idl.json` of the Solana program you want to index. An example config has been provided for you in `test/config.yaml` which is displayed below:

```yaml
version: 1
database:
  sqlite:
    file: ./test/test.db
chains:
  - solana:
      node: https://api.mainnet-beta.solana.com
      programs:
        - program_id: SMPLecH534NA9acpos4G6x7uf3LWbCAwZQE9e8ZekMu
          idl: ./test/squads_mpl.json
```

1. Build and run the service with the test config file.

```
make
./bin/trickle -config ./test/config.yaml
```

2. The server will load the account data at start up and start serving it afterwards. To see what endpoints are available, you can navigate to the Swagger UI url http://localhost:8080/swagger.

## Goals

It is designed with the following goals in mind:

- **Batteries Included**: Easy to use for simple use cases with minimal set up. No need to set up a complicated deployment if your data fits on a single node.
- **Extensible**: Modular design to accomodate different blockchains, databases, etc.
- **Community First**: Encourage collaboration on data tooling in the web3 community through open source with a permissive license (Apache).

### Roadmap

- [ ] Data Loading
    - [ ] Solana
        - [x] Account data loading using getProgramAccounts (at start up)
        - [ ] Account data loading using getProgramAccounts (ongoing)
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