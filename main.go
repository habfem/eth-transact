package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//"github.com/ethereum/go-ethereum/accounts/keystore"

var (
	url = "https://sepolia.infura.io/v3/c76438a754b54ef2bea366c87b7500fb"
)

func main() {
	/* ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount("xxxxxx")
	if err != nil {
		log.Fatal(err)
	}

	_, err = ks.NewAccount("xxxxxx")
	if err != nil {
		log.Fatal(err)
	}
	"33cf0dc27a3334ac2d9c579b7e2ee56cceaa4b62"
	"6b62f88057a82449d3bb1ab5041460b71be4c77c" */

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	a1 := common.HexToAddress("33cf0dc27a3334ac2d9c579b7e2ee56cceaa4b62")
	a2 := common.HexToAddress("6b62f88057a82449d3bb1ab5041460b71be4c77c")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	//1 eth = 10^18 wei
	amount := big.NewInt(10000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2023-11-07T08-26-52.205876600Z--33cf0dc27a3334ac2d9c579b7e2ee56cceaa4b62")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "habfem")
	if err != nil {
		log.Fatal(err)
	}
	tx1, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	client.SendTransaction(context.Background(), tx1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx1.Hash().Hex())
}
