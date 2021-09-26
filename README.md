# nft
NFT CLI and Library.

这个代码库只有两个很简单的目的：
1. 提供一个命令行小工具方便快速访问一些NFT的属性，因为他们可能还没来得及开源合约，因此Etherscan上暂时查不到；
2. 提供ERC721的Go语言绑定库，方便其他库直接使用；

## 命令行工具

```shell
NAME:
   nft.0.1.1 - Common NFT CLI

USAGE:
   nft.0.1.1 [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
   name      Get token name
   symbol    Get token symbol
   tokenURI  Get token URI meta data
   ownerOf   Get owner of token
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --endpoint value, -e value  Ethereum RPC endpoint (default: "https://mainnet.infura.io/v3/")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```
