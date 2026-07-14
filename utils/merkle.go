package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"toy-blockchain/transaction"
)

// hashTransaction converts a transaction into deterministic string data
// and generates its SHA-256 hash in hex representation.
func hashTransaction(tx transaction.Transaction) string {
	data := fmt.Sprintf("%s:%s:%f:%s:%s", tx.Sender, tx.Receiver, tx.Amount, tx.PublicKey, tx.Signature)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// CalculateMerkleRoot calculates the Merkle root for a slice of transactions.
// It returns an empty string if there are no transactions.
// If there is only one transaction, it returns its hash.
// For multiple transactions, it builds the Merkle tree recursively or iteratively
// and handles odd-length levels by duplicating the last hash.
func CalculateMerkleRoot(transactions []transaction.Transaction) string {
	if len(transactions) == 0 {
		return ""
	}

	// Initialize the leaf hashes
	var level []string
	for _, tx := range transactions {
		level = append(level, hashTransaction(tx))
	}

	// Iteratively combine hashes until only one remains
	for len(level) > 1 {
		var nextLevel []string
		for i := 0; i < len(level); i += 2 {
			if i+1 < len(level) {
				// Pair exists
				combined := level[i] + level[i+1]
				hash := sha256.Sum256([]byte(combined))
				nextLevel = append(nextLevel, hex.EncodeToString(hash[:]))
			} else {
				// Odd element: duplicate the last hash
				combined := level[i] + level[i]
				hash := sha256.Sum256([]byte(combined))
				nextLevel = append(nextLevel, hex.EncodeToString(hash[:]))
			}
		}
		level = nextLevel
	}

	return level[0]
}
