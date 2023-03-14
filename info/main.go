package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hesusruiz/redtmon/redt"
)

func main() {
	// Connect to the RedT node
	nodeURL := "ws://15.236.0.91:22001"
	rt, err := redt.NewRedTNode(nodeURL)
	if err != nil {
		log.Panicln(err)
	}

	// Check that we can get info about the node
	thisNode, err := rt.NodeInfo()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("name", thisNode.Name)

	// // We are going to call the Geth API, with a timeout of 30 seconds
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	vals := rt.Validators()
	for _, j := range vals {
		info := rt.ValidatorInfo(j)
		fmt.Printf("%v -> %v\n", info.Address, info.Operator)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	txHash := common.HexToHash("0xd5402b7ef13fe9749eead43923eb9a30c6f4694ca3a4e796007586d3957d22f3")
	txReceipt, err := rt.EthClient().TransactionReceipt(ctx, txHash)
	if err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(txReceipt, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}
