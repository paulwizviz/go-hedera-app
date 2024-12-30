package main

import (
	"fmt"
	"os"

	"github.com/hiero-ledger/hiero-sdk-go/v2"
)

func createHederaClient() *hiero.Client {

	prvkey, err := hiero.PrivateKeyFromString("302e...") // Replace with your private key
	if err != nil {
		fmt.Println(err)
		return nil
	}

	acctID := hiero.AccountID{Account: 12345} // Replace with your account ID

	client := hiero.ClientForTestnet()
	client.SetOperator(acctID, prvkey)
	return client
}

func deploySmartContract(client *hiero.Client, bytecode []byte) (hiero.ContractID, error) {
	fileTx, err := hiero.NewFileCreateTransaction().
		SetKeys(client.GetOperatorPublicKey()).
		SetContents(bytecode).
		Execute(client)
	if err != nil {
		return hiero.ContractID{}, err
	}

	fileReceipt, err := fileTx.GetReceipt(client)
	if err != nil {
		return hiero.ContractID{}, err
	}

	contractFileID := fileReceipt.FileID

	contractTx, err := hiero.NewContractCreateTransaction().
		SetBytecodeFileID(*contractFileID).
		SetGas(100000).
		Execute(client)
	if err != nil {
		return hiero.ContractID{}, err
	}

	contractReceipt, err := contractTx.GetReceipt(client)
	if err != nil {
		return hiero.ContractID{}, err
	}

	return *contractReceipt.ContractID, nil
}

func callContractFunction(client *hiero.Client, contractID hiero.ContractID) error {
	callTx, err := hiero.NewContractExecuteTransaction().
		SetContractID(contractID).
		SetGas(100000).
		SetFunction("store", hiero.NewContractFunctionParameters().AddUint64(42)).
		Execute(client)
	if err != nil {
		return err
	}

	_, err = callTx.GetReceipt(client)
	return err
}

func main() {
	client := createHederaClient()
	defer client.Close()

	bytecode, err := os.ReadFile("build/SimpleStorage.bin")
	if err != nil {
		fmt.Println("Error reading bytecode:", err)
		return
	}

	contractID, err := deploySmartContract(client, bytecode)
	if err != nil {
		fmt.Println("Error deploying contract:", err)
		return
	}

	fmt.Println("Contract deployed with ID:", contractID)

	err = callContractFunction(client, contractID)
	if err != nil {
		fmt.Println("Error calling contract:", err)
		return
	}

	fmt.Println("Contract function executed successfully!")
}
