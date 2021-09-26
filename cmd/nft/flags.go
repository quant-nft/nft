package main

import "github.com/urfave/cli/v2"

var (
	EndpointFlag = &cli.StringFlag{
		Name:    "endpoint",
		Aliases: []string{"e"},
		Value:   "https://mainnet.infura.io/v3/",
		Usage:   "Ethereum RPC endpoint",
	}
	AddressFlag = &cli.StringFlag{
		Name:    "address",
		Aliases: []string{"a"},
		Usage:   "contract `address`",
	}
	tokenIdFlag = &cli.Int64Flag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "token `Id`",
	}
)
