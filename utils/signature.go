package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"toy-blockchain/transaction"
)

// GenerateKeyPair creates a new ECDSA private key and public key.
//
// Private key:
// - kept secret by the user
// - used to sign transactions
//
// Public key:
// - shared with the blockchain
// - used to verify signatures
func GenerateKeyPair() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)
	if err != nil {
		return nil, "", err
	}

	publicKeyBytes := elliptic.Marshal(
		elliptic.P256(),
		privateKey.PublicKey.X,
		privateKey.PublicKey.Y,
	)

	publicKey := base64.StdEncoding.EncodeToString(publicKeyBytes)

	return privateKey, publicKey, nil
}

// SignTransaction creates a digital signature for transaction data string.
//
// The data string is hashed first using SHA-256.
// The hash is then signed using the private key.
func SignTransaction(data string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	rBytes := r.Bytes()
	sBytes := s.Bytes()

	// Pad both r and s to 32 bytes to ensure safe concatenation and split (64 bytes total)
	rPad := make([]byte, 32)
	sPad := make([]byte, 32)
	copy(rPad[32-len(rBytes):], rBytes)
	copy(sPad[32-len(sBytes):], sBytes)

	signatureBytes := append(rPad, sPad...)
	signature := base64.StdEncoding.EncodeToString(signatureBytes)

	return signature, nil
}

// VerifySignature reconstructs a public key and verifies that the signature of the data is valid.
func VerifySignature(data string, signature string, publicKey string) bool {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false
	}

	x, y := elliptic.Unmarshal(
		elliptic.P256(),
		publicKeyBytes,
	)
	if x == nil || y == nil {
		return false
	}

	pubKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(signatureBytes) != 64 {
		return false
	}

	rBytes := signatureBytes[:32]
	sBytes := signatureBytes[32:]

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	hash := sha256.Sum256([]byte(data))

	return ecdsa.Verify(&pubKey, hash[:], r, s)
}

// VerifyTransactionSignature verifies that the transaction was signed by the owner of the private key.
func VerifyTransactionSignature(tx transaction.Transaction) bool {
	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	return VerifySignature(data, tx.Signature, tx.PublicKey)
}