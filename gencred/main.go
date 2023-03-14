package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hesusruiz/redtmon/redt"
)

func main() {
	var (
		nodeKeyFile = flag.String("nodekey", "", "filename with private key")
		nodeKeyHex  = flag.String("nodekeyhex", "", "private key as hex")
		nodeKey     *ecdsa.PrivateKey
		err         error
	)

	// Parse the command line arguments
	flag.Parse()

	// Get the nodekey either from a file or from the command line
	switch {
	case *nodeKeyFile == "" && *nodeKeyHex == "":
		log.Fatalf("Use -nodekey or -nodekeyhex to specify a private key")
	case *nodeKeyFile != "" && *nodeKeyHex != "":
		log.Fatalf("Options -nodekey and -nodekeyhex are mutually exclusive")
	case *nodeKeyFile != "":
		if nodeKey, err = crypto.LoadECDSA(*nodeKeyFile); err != nil {
			log.Fatalf("-nodekey: %v", err)
		}
	case *nodeKeyHex != "":
		if nodeKey, err = crypto.HexToECDSA(*nodeKeyHex); err != nil {
			log.Fatalf("-nodekeyhex: %v", err)
		}
	}

	fmt.Printf("Public key: %x\n", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])

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

	// We are going to call the Geth API, with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Retrieve the header from the node (may be remote)
	header, err := rt.EthClient().HeaderByNumber(ctx, nil)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Current block", header.Number, header.Hash())

	// Create the message to sign
	dataToSign := fmt.Sprintf("%d,0x%x,%v", header.Number.Int64(), crypto.FromECDSAPub(&nodeKey.PublicKey)[1:], header.Hash().String())
	fmt.Println(dataToSign)

	// Hash the data
	md := sha512.New()
	md.Write([]byte(dataToSign))
	hashToSign := md.Sum(nil)[:32]

	// Sign
	sig, err := crypto.Sign(hashToSign, nodeKey)
	if err != nil {
		log.Panicln(err)
	}
	// sigShort := sig[:20]

	fmt.Printf("%x\n", sig)

	// Now try to verify the signature
	valid := crypto.VerifySignature(crypto.FromECDSAPub(&nodeKey.PublicKey), hashToSign, sig[:64])
	if valid {
		fmt.Println("FANTASTICO!!!")
	} else {
		fmt.Println("Cagada")
	}
}
