package main

import (
	"fmt"
	"log"

	"github.com/hiero-ledger/hiero-sdk-go/v2"
)

func main() {
	// Step 1: Generate a Wallet (Private/Public Key Pair)
	privateKey, err := hiero.PrivateKeyGenerateEd25519()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}
	publicKey := privateKey.PublicKey()

	fmt.Printf("Generated Wallet:\n")
	fmt.Printf("Private Key: %s\n", privateKey.String())
	fmt.Printf("Public Key: %s\n", publicKey.String())

	// Step 2: Connect to the Hedera Testnet
	operatorID := "0.0.12345"                            // Replace with your operator account ID
	operatorKey := "302e020100300506032b657004220420..." // Replace with your operator private key
	client := connectToTestnet(operatorID, operatorKey)
	defer client.Close()

	// Step 3: Create a Hedera Account with the Public Key
	accountID := createAccount(client, publicKey)
	fmt.Printf("Created Account ID: %s\n", accountID.String())
}

// Function to Connect to the Hedera Testnet
func connectToTestnet(operatorID string, operatorKey string) *hiero.Client {
	client := hiero.ClientForTestnet()

	// Parse operator credentials
	accountID, err := hiero.AccountIDFromString(operatorID)
	if err != nil {
		log.Fatalf("Error parsing account ID: %v", err)
	}
	privateKey, err := hiero.PrivateKeyFromString(operatorKey)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	// Set operator credentials
	client.SetOperator(accountID, privateKey)
	fmt.Println("Connected to Hedera Testnet!")
	return client
}

// Function to Create a Hedera Account
func createAccount(client *hiero.Client, publicKey hiero.PublicKey) *hiero.AccountID {
	transactionResponse, err := hiero.NewAccountCreateTransaction().
		SetKey(publicKey).
		SetInitialBalance(hiero.HbarFrom(10, "hbar")). // Provide initial balance
		Execute(client)
	if err != nil {
		log.Fatalf("Error creating account: %v", err)
	}

	// Retrieve the account ID from the receipt
	receipt, err := transactionResponse.GetReceipt(client)
	if err != nil {
		log.Fatalf("Error getting receipt: %v", err)
	}

	accountID := receipt.AccountID
	fmt.Printf("New Account ID: %v\n", accountID)
	return accountID
}
