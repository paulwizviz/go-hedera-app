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

// This example obtain a list of acccounts from mirror nodes via RESTFul API.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/paulwizviz/go-hedera-app/internal/rest"
)

func main() {
	mirrorNodeURL := "https://testnet.mirrornode.hedera.com/api/v1"
	client := rest.NewDefaultClient(mirrorNodeURL)
	accts, err := client.ListAccounts(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for _, acct := range accts.Accounts {
		fmt.Println(acct.Account, *acct.Key)
	}
}
