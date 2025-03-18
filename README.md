# Oasis Bundle Registry

This repository contains metadata that is used by the nodes for automated
bundle discovery and distribution.

## List of Bundles

### Mainnet

| Bundle                                                                                                              | Manifest Hash                                                    |
| ------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Sapphire 0.9.2](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.9.2/sapphire-paratime.orc) | 91b1ff8b5cc9f6a99e97be9c69f22d73d75e037b8bbca7fda64a7a41bc997841 |
| [Sapphire 0.8.2](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.8.2/sapphire-paratime.orc) | e523903e480a8bef7caf18b846aefaa17913878b67eee13ac618849dd0bb8741 |
| [Cipher 3.2.2](https://github.com/oasisprotocol/cipher-paratime/releases/download/v3.2.2/cipher-paratime.orc) | e653d9fc0cc950ce018dc11b36ade72b5fc322d70e78501549a89e60ed22ded0 |

### Testnet

| Bundle                                                                                                                              | Manifest Hash                                                    |
| ----------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| [Sapphire 0.9.2-testnet](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.9.2-testnet/sapphire-paratime.orc) | ff9ab01f1a5a043c3b105a67691fbd2223b76136be734c6a322493ce9f3a7280 |
| [Sapphire 0.9.1-testnet](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.9.1-testnet/sapphire-paratime.orc) | 30e8e286c45ef0b5381fbcf5ff3e3532ed763ef18e6de736fe15013d482e030d |
| [Sapphire 0.9.0-testnet](https://github.com/oasisprotocol/sapphire-paratime/releases/download/v0.9.0-testnet/sapphire-paratime.orc) | 7af652553a1fe1e2b2ec4535458eed4dd760cce1dd4df00c30b164523ecd674e |
| [Cipher 3.2.2-testnet](https://github.com/oasisprotocol/cipher-paratime/releases/download/v3.2.2-testnet/cipher-paratime.orc) | aa424b14c14a58274372ff710702d08f3a2261808d00d3a6e132534949309320 |
| [Cipher 3.2.1-testnet](https://github.com/oasisprotocol/cipher-paratime/releases/download/v3.2.1-testnet/cipher-paratime.orc) | c3580cc119821926478135974cfe705d291042e4010a0f3b1173c98e6223e155 |
| [Cipher 3.2.0-testnet](https://github.com/oasisprotocol/cipher-paratime/releases/download/v3.2.0-testnet/cipher-paratime.orc) | 3e5b5c266b9db4f67ca7b3f8e235eab37ce569cffa93bc3e01e15c3a23b9ccc9 |

## Verify Metadata

To ensure that the name of the metadata file matches the hash of the bundle
manifest linked within the metadata file, run the following commands:

```bash
cd src
go test ./... -count=1 -v
```
