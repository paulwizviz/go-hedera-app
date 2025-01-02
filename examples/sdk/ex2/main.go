// Copyright 2025 The Contributors to go-hedera-app
// This file is part of the go-hedera-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

// This example demonstrate opeartion to create HTS Token

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hiero-ledger/hiero-sdk-go/v2"
)

func main() {
	privKey := os.Getenv("HEDERA_DEV_PRIV_KEY")
	acctID := os.Getenv("HEDERA_ACCT_ID")

	key, err := hiero.PrivateKeyFromString(privKey[2:])
	if err != nil {
		log.Fatal(err)
	}

	acct, err := hiero.AccountIDFromString(acctID)
	if err != nil {
		log.Fatal(err)
	}

	client := hiero.ClientForTestnet()
	client.SetOperator(acct, key)
	defer client.Close()

	// Generate a supply key
	supplyKey, err := hiero.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate supply key: %v", err)
	}
	fmt.Printf("Generated supply key: %s\n", supplyKey.String())

	// Create the token
	tokenCreateTx, err := hiero.NewTokenCreateTransaction().
		SetTokenName("MyToken").
		SetTokenSymbol("MTK").
		SetDecimals(2).
		SetInitialSupply(100000). // Total supply: 1,000.00
		SetTreasuryAccountID(acct).
		SetSupplyKey(supplyKey). // Set the supply key here
		FreezeWith(client)
	if err != nil {
		log.Fatal(err)
	}

	// Sign and execute the transaction
	signedTx := tokenCreateTx.Sign(key)
	if err != nil {
		log.Fatalf("Failed to sign token create transaction: %v", err)
	}
	txResponse, err := signedTx.Execute(client)
	if err != nil {
		log.Fatalf("Failed to execute token create transaction: %v", err)
	}

	// Get the receipt to confirm the transaction
	receipt, err := txResponse.GetReceipt(client)
	if err != nil {
		log.Fatalf("Failed to get receipt: %v", err)
	}

	// Extract the token ID from the receipt
	tokenID := receipt.TokenID
	fmt.Printf("Created token with ID: %s\n", tokenID)
}
