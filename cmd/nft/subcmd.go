package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quant-nft/nft/erc721"
	"github.com/urfave/cli/v2"
	"math/big"
)

var (
	nameCommand = &cli.Command{
		Action: name,
		Name:   "name",
		Usage:  "Get token name",
		Flags: []cli.Flag{
			AddressFlag,
		},
	}
	symbolCommand = &cli.Command{
		Action: symbol,
		Name:   "symbol",
		Usage:  "Get token symbol",
		Flags: []cli.Flag{
			AddressFlag,
		},
	}
	tokenURICommand = &cli.Command{
		Action: tokenURI,
		Name:   "tokenURI",
		Usage:  "Get token URI meta data",
		Flags: []cli.Flag{
			AddressFlag,
			tokenIdFlag,
		},
	}
	ownerOfCommand = &cli.Command{
		Action: ownerOf,
		Name:   "ownerOf",
		Usage:  "Get owner of token ",
		Flags: []cli.Flag{
			AddressFlag,
			tokenIdFlag,
		},
	}
)

func name(c *cli.Context) error {
	endpoint := c.String(EndpointFlag.Name)
	ec, err := ethclient.DialContext(c.Context, endpoint)
	if err != nil {
		return err
	}
	defer ec.Close()
	fmt.Printf("Endpoint is %s\n", endpoint)

	address := common.HexToAddress(c.String(AddressFlag.Name))
	fmt.Printf("Contract is %s\n", address.Hex())

	nft, err := erc721.NewErc721(address, ec)
	if err != nil {
		return err
	}
	n, err := nft.Name(nil)
	if err != nil {
		return err
	}
	fmt.Printf("Symbol is %s\n", n)
	return nil
}

func symbol(c *cli.Context) error {
	endpoint := c.String(EndpointFlag.Name)
	ec, err := ethclient.DialContext(c.Context, endpoint)
	if err != nil {
		return err
	}
	defer ec.Close()
	fmt.Printf("Endpoint is %s\n", endpoint)

	address := common.HexToAddress(c.String(AddressFlag.Name))
	fmt.Printf("Contract is %s\n", address.Hex())

	nft, err := erc721.NewErc721(address, ec)
	if err != nil {
		return err
	}
	s, err := nft.Symbol(nil)
	if err != nil {
		return err
	}
	fmt.Printf("Symbol is %s\n", s)
	return nil
}

func tokenURI(c *cli.Context) error {
	endpoint := c.String(EndpointFlag.Name)
	ec, err := ethclient.DialContext(c.Context, endpoint)
	if err != nil {
		return err
	}
	defer ec.Close()
	fmt.Printf("Endpoint is %s\n", endpoint)

	address := common.HexToAddress(c.String(AddressFlag.Name))
	fmt.Printf("Contract is %s\n", address.Hex())
	tokenId := c.Int64(tokenIdFlag.Name)
	fmt.Printf("Token Id is %d\n", tokenId)

	nft, err := erc721.NewErc721(address, ec)
	if err != nil {
		return err
	}
	uri, err := nft.TokenURI(nil, big.NewInt(tokenId))
	if err != nil {
		return err
	}
	fmt.Printf("Token %d URI is \n%s\n", tokenId, uri)
	return nil
}

func ownerOf(c *cli.Context) error {
	endpoint := c.String(EndpointFlag.Name)
	ec, err := ethclient.DialContext(c.Context, endpoint)
	if err != nil {
		return err
	}
	defer ec.Close()
	fmt.Printf("Endpoint is %s\n", endpoint)

	address := common.HexToAddress(c.String(AddressFlag.Name))
	fmt.Printf("Contract is %s\n", address.Hex())
	tokenId := c.Int64(tokenIdFlag.Name)
	fmt.Printf("Token Id is %d\n", tokenId)

	nft, err := erc721.NewErc721(address, ec)
	if err != nil {
		return err
	}
	owner, err := nft.OwnerOf(nil, big.NewInt(tokenId))
	if err != nil {
		return err
	}
	fmt.Printf("Owner of Token %d is %s\n", tokenId, owner)
	return nil
}
