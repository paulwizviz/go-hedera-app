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

package wallet

import (
	"errors"
	"fmt"

	"github.com/hiero-ledger/hiero-sdk-go/v2"
)

var (
	ErrInvalidPrivKeyStringFormat = errors.New("invalid private key string format")
	ErrInvalidAcctIDStringFormat  = errors.New("invalid account ID string format")
)

func CreateClient(privKey string, acctID string) (*hiero.Client, error) {

	var key hiero.PrivateKey
	var err error
	if privKey[0:2] == "0x" {
		key, err = hiero.PrivateKeyFromString(privKey[2:])
		if err != nil {
			return nil, fmt.Errorf("%w-%v", ErrInvalidPrivKeyStringFormat, err)
		}
	} else {
		key, err = hiero.PrivateKeyFromString(privKey)
		if err != nil {
			return nil, fmt.Errorf("%w-%v", ErrInvalidPrivKeyStringFormat, err)
		}
	}

	id, err := hiero.AccountIDFromString(acctID)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrInvalidAcctIDStringFormat, err)
	}

	return createClient(key, id, false)
}

func createClient(privKey hiero.PrivateKey, acctID hiero.AccountID, forMainnet bool) (*hiero.Client, error) {

	if forMainnet {
		client := hiero.ClientForMainnet()
		client.SetOperator(acctID, privKey)
		return client, nil
	}

	client := hiero.ClientForTestnet()
	client.SetOperator(acctID, privKey)
	return client, nil
}
