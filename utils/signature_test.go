package utils_test

import (
	"fmt"
	"testing"
	"toy-blockchain/transaction"
	"toy-blockchain/utils"
)

func TestDigitalSignature(t *testing.T) {
	// Test 1: Create valid signed transaction, verify signature succeeds.
	priv, pub, err := utils.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	tx := transaction.Transaction{
		Sender:   "alice",
		Receiver: "bob",
		Amount:   100.0,
	}

	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	sig, err := utils.SignTransaction(data, priv)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	tx.PublicKey = pub
	tx.Signature = sig

	if !utils.VerifyTransactionSignature(tx) {
		t.Error("VerifyTransactionSignature failed for valid signed transaction")
	}

	// Test 2: Modify transaction amount after signing. Verification must fail.
	txModifiedAmount := tx
	txModifiedAmount.Amount = 200.0
	if utils.VerifyTransactionSignature(txModifiedAmount) {
		t.Error("Verification should have failed when amount was modified")
	}

	// Test 3: Modify receiver after signing. Verification must fail.
	txModifiedReceiver := tx
	txModifiedReceiver.Receiver = "charlie"
	if utils.VerifyTransactionSignature(txModifiedReceiver) {
		t.Error("Verification should have failed when receiver was modified")
	}

	// Test 4: Modify sender after signing. Verification must fail.
	txModifiedSender := tx
	txModifiedSender.Sender = "eve"
	if utils.VerifyTransactionSignature(txModifiedSender) {
		t.Error("Verification should have failed when sender was modified")
	}

	// Test 5: Invalid signature should reject transaction.
	txInvalidSig := tx
	txInvalidSig.Signature = "invalid_signature_data_here"
	if utils.VerifyTransactionSignature(txInvalidSig) {
		t.Error("Verification should have failed with invalid signature")
	}
}
