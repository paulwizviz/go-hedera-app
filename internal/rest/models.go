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

package rest

// Token represents a token created for a given transactions
type Token struct {
	TokenID string `json:"token_id"`
	Balance int    `json:"balance"`
}

// Balance represents the balance of a tokens held
type Balance struct {
	Timestamp string  `json:"timestamp"`
	Balance   int     `json:"balance"`
	Tokens    []Token `json:"tokens"`
}

// Key reprsents the cryptographic key
type Key struct {
	Type string `json:"_type"`
	Key  string `json:"key"`
}

// Account represents information related to an account
type Account struct {
	Account           string  `json:"account"`
	Alias             string  `json:"alias"`
	AutoRenewPeriod   *int    `json:"auto_renew_period,omitempty"`
	Balance           Balance `json:"balance"`
	CreatedTimestamp  string  `json:"created_timestamp"`
	DeclineReward     bool    `json:"decline_reward"`
	Deleted           bool    `json:"deleted"`
	EthereumNonce     int     `json:"ethereum_nonce"`
	EVMAddress        string  `json:"evm_address"`
	ExpiryTimestamp   *string `json:"expiry_timestamp,omitempty"`
	Key               *Key    `json:"key,omitempty"`
	MaxAutoTokenAssoc int     `json:"max_automatic_token_associations"`
	Memo              string  `json:"memo"`
	PendingReward     int     `json:"pending_reward"`
	ReceiverSigReq    bool    `json:"receiver_sig_required"`
	StakedAcctID      *string `json:"staked_account_id,omitempty"`
	StakedNodeID      int     `json:"staked_node_id"`
	StakePeriodStart  string  `json:"stake_period_start"`
}
