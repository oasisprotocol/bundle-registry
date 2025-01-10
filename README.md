# Oasis Bundle Registry

This repository contains metadata that is used by the nodes for automated
bundle discovery and distribution.

## List of Bundles

### Mainnet

| Bundle                                                                                                              | Manifest Hash                                                    |
| ------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Sapphire 0.8.2](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.8.2/sapphire-paratime.orc) | e523903e480a8bef7caf18b846aefaa17913878b67eee13ac618849dd0bb8741 |

### Testnet

| Bundle                                                                                                                              | Manifest Hash                                                    |
| ----------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Sapphire 0.9.0-testnet](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.9.0-testnet/sapphire-paratime.orc) | 7af652553a1fe1e2b2ec4535458eed4dd760cce1dd4df00c30b164523ecd674e |

## Verify Metadata

To ensure that the name of the metadata file matches the hash of the bundle
manifest linked within the metadata file, run the following commands:

```bash
cd src
go test ./... -count=1 -v
```
