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

// This example demonstrate the use of Go SDK to retrieve account information

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

	acctBal, err := hiero.NewAccountBalanceQuery().SetAccountID(acct).Execute(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(acctBal.Hbars)

}
