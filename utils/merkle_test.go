package utils_test

import (
	"testing"
	"toy-blockchain/transaction"
	"toy-blockchain/utils"
)

func TestMerkleRootFromMultipleTransactions(t *testing.T) {
	txs := []transaction.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10.0, PublicKey: "pub1", Signature: "sig1"},
		{Sender: "Bob", Receiver: "Charlie", Amount: 5.0, PublicKey: "pub2", Signature: "sig2"},
		{Sender: "Charlie", Receiver: "Alice", Amount: 2.5, PublicKey: "pub3", Signature: "sig3"},
	}

	root := utils.CalculateMerkleRoot(txs)
	if root == "" {
		t.Fatal("Expected Merkle root from multiple transactions to not be empty")
	}
}

func TestMerkleRootChangesOnTransactionTampering(t *testing.T) {
	txs := []transaction.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10.0, PublicKey: "pub1", Signature: "sig1"},
		{Sender: "Bob", Receiver: "Charlie", Amount: 5.0, PublicKey: "pub2", Signature: "sig2"},
	}

	rootOriginal := utils.CalculateMerkleRoot(txs)
	if rootOriginal == "" {
		t.Fatal("Expected original Merkle root to not be empty")
	}

	// Change amount on one transaction
	txs[1].Amount = 50.0

	rootTampered := utils.CalculateMerkleRoot(txs)
	if rootTampered == "" {
		t.Fatal("Expected tampered Merkle root to not be empty")
	}

	if rootOriginal == rootTampered {
		t.Errorf("Expected Merkle root to change when transaction amount is modified, but got identical root: %s", rootOriginal)
	}
}

func TestMerkleRootSingleTransaction(t *testing.T) {
	txs := []transaction.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10.0, PublicKey: "pub1", Signature: "sig1"},
	}

	root := utils.CalculateMerkleRoot(txs)
	if root == "" {
		t.Fatal("Expected Merkle root for a single transaction to not be empty")
	}
}
