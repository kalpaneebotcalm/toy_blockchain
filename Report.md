# TOY BLOCKCHAIN AND LEDGER SIMULATOR

**A Simplified Blockchain Implementation in Go**

**Project Report**

Department of Electrical and Information Engineering
Faculty of Engineering
University of Ruhuna

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Objectives](#2-objectives)
3. [System Overview](#3-system-overview)
4. [System Architecture](#4-system-architecture)
5. [Implementation](#5-implementation)
6. [Features](#6-features)
7. [Testing](#7-testing)
8. [Results and Discussion](#8-results-and-discussion)
9. [Future Improvements](#9-future-improvements)
10. [Conclusion](#10-conclusion)

---

## 1. Introduction

### 1.1 Background

Blockchain technology is a distributed ledger technology that enables secure and transparent recording of digital transactions. A blockchain consists of a sequence of blocks, where each block contains transaction information, a timestamp, a reference to the previous block, and a cryptographic hash. Linking blocks through cryptographic hashes ensures data integrity, since modifying information in one block affects the entire chain.

Modern blockchain systems such as Bitcoin and Ethereum rely on cryptographic hashing, consensus mechanisms, digital signatures, and peer-to-peer networking. This project focuses on developing a simplified Toy Blockchain and Ledger Simulator in Go to demonstrate these fundamental concepts by implementing core features from scratch, without using existing blockchain frameworks.

### 1.2 Problem Statement

Traditional database systems allow administrators to modify stored data directly, which can affect data integrity. Blockchain technology solves this by creating an immutable chain of records where each transaction is cryptographically linked. The project addresses the question of how the fundamental mechanisms of a blockchain system — cryptographic hashing, Proof-of-Work mining, transaction validation, blockchain validation, and persistent storage — can be implemented from scratch to provide secure transaction processing and ledger management.

### 1.3 Scope of the Project

The scope covers the following areas:

- **Blockchain Management** — creating the Genesis Block, adding new blocks, linking blocks using previous hashes, and validating blockchain integrity.
- **Transaction Management** — creating transactions, maintaining pending transactions, validating transaction rules, and preventing overspending.
- **Mining** — implementing Proof-of-Work, searching for valid nonce values, and generating blocks that satisfy the mining difficulty.
- **Ledger Management** — calculating account balances, tracking transaction history, and preventing double spending.
- **Data Persistence** — saving blockchain data to JSON and reloading previous state on restart.

---

## 2. Objectives

### 2.1 Main Objective

The main objective of this project is to design and implement a simplified blockchain and ledger simulator using Go to demonstrate the fundamental concepts of blockchain technology, including hashing, mining, transaction validation, and blockchain verification.

### 2.2 Specific Objectives

The project sets out to achieve the following:

1. Implement a custom blockchain structure containing blocks, transactions, hash values, previous block references, nonce values, and timestamps.
2. Implement SHA-256 cryptographic hashing to generate unique block identifiers, maintain data integrity, and detect unauthorized modifications.
3. Develop a Proof-of-Work mining algorithm where the nonce is repeatedly modified until a valid hash is generated according to the configured difficulty.
4. Implement transaction validation that checks amount validity, sender/receiver information, self-transaction prevention, available balance, and pending double-spend prevention.
5. Implement ledger management that calculates account balances by analyzing all confirmed transactions, simulating how cryptocurrency systems track balances without a centralized database.
6. Implement blockchain integrity verification by checking block hashes, previous-hash connections, Proof-of-Work requirements, and timestamp consistency.
7. Implement JSON-based persistent storage to save blockchain information and restore previous states after restart.
8. Develop a command-line interface allowing users to add transactions, mine blocks, view blockchain information, check balances, and validate the blockchain.

---

## 3. System Overview

### 3.1 Overview

The Toy Blockchain and Ledger Simulator is a modular command-line application developed using Go. The system is divided into multiple packages — Block, Transaction, Blockchain, Mining, Ledger, Storage, Utility, and CLI — where each package is responsible for a specific blockchain functionality. This modular design improves maintainability, readability, and separation of responsibilities.

### 3.2 System Workflow

The overall workflow proceeds through six steps:

1. **Blockchain Initialization** — the system checks whether a blockchain data file exists; if not, a new blockchain is created with a Genesis Block.
2. **Transaction Creation** — a user creates a transaction via the CLI (e.g. `addtx faucet alice 100`); the system validates and adds it to the pending pool.
3. **Block Creation** — when mining is requested, pending transactions are collected into a new candidate block referencing the previous block's hash.
4. **Proof-of-Work Mining** — the system generates a hash, checks it against the difficulty target, and increments the nonce until a valid hash is found.
5. **Blockchain Update** — the mined block is appended, its transactions are removed from the pending pool, and the chain is saved.
6. **Blockchain Validation** — the system checks hash correctness, block linkage, Proof-of-Work compliance, and timestamp order to detect tampering.

### 3.3 System Characteristics

The implemented system provides the following characteristics:

- **Security** through SHA-256 hashing, Proof-of-Work, and integrity validation.
- **Data integrity**, since each block references the previous block's hash so unauthorized modifications can be detected.
- **Transaction reliability**, preventing invalid transactions, overspending, and double spending.
- **Persistence** through JSON-based storage across executions.
- **Modularity**, with each component having a clear, testable responsibility.

---

## 4. System Architecture

### 4.1 Overview

The system follows a modular, layered architecture consisting of a User Interface Layer (CLI), a Blockchain Management Layer, Core Blockchain Components (Block, Transaction, Mining, Ledger, Utility), and a Data Persistence Layer. Each package handles one blockchain operation, improving code organization, maintainability, and scalability.

### 4.2 High-Level Architecture

```
                          User
                           |
                           v
              Command Line Interface (main.go)
                           |
              -------------------------------
              |                             |
              v                             v
      Blockchain Module              Storage Module
      (blockchain package)           (storage package)
              |
   -----------------------------------------
   |          |          |          |
   v          v          v          v
 Block    Transaction   Mining    Ledger
 Module     Module      Module    Module
              |
              v
      Hash Utility Module (utils)
```

### 4.3 Key Components

**Block Module**

Defines the block structure: Index, Timestamp, Transactions, PreviousHash, Hash, and Nonce. Each new block stores the hash of the previous block, forming the chain. If any previous block is modified, this linkage becomes invalid and validation fails.

| Field | Description |
|---|---|
| Index | Position of the block in the chain |
| Timestamp | Creation time of the block |
| Transactions | List of transactions stored in the block |
| PreviousHash | Hash of the previous block |
| Hash | Current block hash |
| Nonce | Mining value used for Proof-of-Work |

**Transaction Module**

Represents the transfer of coins between accounts (Sender, Receiver, Amount). Transactions are validated by the Blockchain module before being accepted into the pending pool.

**Mining Module**

Implements Proof-of-Work: it creates a candidate block, sets the nonce to zero, generates a SHA-256 hash, checks it against the difficulty target, and increases the nonce until a valid hash is found.

```
Create Block -> Set Nonce = 0 -> Generate SHA-256 Hash -> Check Difficulty
                    ^                                          |
                    |------------ Invalid: Increase Nonce ------|
                                                                 |
                                                    Valid: Add Block
```

**Hash Utility Module**

Generates SHA-256 hashes from block index, timestamp, transactions, previous hash, and nonce, producing a 64-character hash that uniquely identifies each block.

**Ledger Module**

Instead of storing balances directly, the ledger calculates balances by processing all confirmed transactions. Example: faucet → Alice: 100, Alice → Bob: 40, giving Alice = 60 and Bob = 40.

**Storage Module**

Saves and loads the blockchain state as `blockchain.json`, allowing the application to resume from where it left off after a restart.

### 4.4 Data Flow

```
User Creates Transaction -> CLI validates -> Pending Pool
-> Miner creates candidate block -> Proof-of-Work mining
-> Valid block generated -> Block added to blockchain
-> Blockchain saved to JSON
```

### 4.5 Design Principles

The design emphasizes modularity (independent, testable packages), separation of responsibilities (e.g. mining handles Proof-of-Work, storage handles persistence, ledger handles balances), and data integrity through hash linking, Proof-of-Work, and validation mechanisms.

---

## 5. Implementation

### 5.1 Block and Transaction Structures

```go
type Block struct {
    Index        int
    Timestamp    int64
    Transactions []transaction.Transaction
    PreviousHash string
    Nonce        int
    Hash         string
}

type Transaction struct {
    Sender   string
    Receiver string
    Amount   float64
}
```

The Genesis Block is created during initialization with Index = 0, PreviousHash = "0", no transactions, and a fixed timestamp for deterministic generation.

### 5.2 Transaction Validation

Before a transaction is accepted into the pending pool, the system enforces several rules:

- **Amount Validation** — negative or zero amounts are rejected (e.g. Alice → Bob: -10 or 0 is rejected; Alice → Bob: 50 is accepted).
- **Account Validation** — sender and receiver cannot be empty or identical (Alice → Alice: 100 is rejected).
- **Balance Verification** — the sender must have sufficient balance (e.g. Alice = 40 cannot send 50 to Bob).
- **Pending Transaction Protection** — available balance also accounts for pending transactions, preventing double spending (e.g. if Alice = 100 with 80 pending to Bob, a new 50-coin transaction to Charlie is rejected since only 20 remains available).

### 5.3 Blockchain Structure and Validation

```go
type Blockchain struct {
    Blocks              []block.Block
    PendingTransactions []transaction.Transaction
}
```

Blockchain validation checks block index order, previous-hash linkage, recalculated current hash against the stored value, Proof-of-Work compliance, and chronological timestamp order — any single failure marks the chain as invalid.

### 5.4 Mining Algorithm

```
Nonce = 0
while hash is invalid:
    Nonce++
    Generate Hash
Return valid block
```

After successful mining, the hash is stored in the block, the block is appended to the blockchain, and its transactions are removed from the pending pool.

### 5.5 Command-Line Interface

The CLI, implemented in `main.go`, exposes the following commands:

```
go run . addtx <sender> <receiver> <amount>   # create a transaction
go run . mine                                 # mine a new block
go run . print                                # display the blockchain
go run . balance <username>                   # check account balance
go run . validate                             # verify blockchain integrity
go run . save                                 # manually save blockchain state
```

---

## 6. Features

The following features have been implemented:

- Blockchain creation and management, starting with a Genesis Block and linked block history.
- Transaction creation and validation, with a pending transaction pool for unconfirmed transfers.
- Transaction security — rejection of negative/zero amounts, self-transactions, insufficient balance, and double spending.
- SHA-256 cryptographic hashing for tamper detection and unique block identification.
- Proof-of-Work mining with configurable difficulty and nonce searching.
- Blockchain integrity verification covering hashes, linkage, Proof-of-Work, and timestamps.
- Ledger and balance management, calculated dynamically from transaction history.
- JSON-based persistent storage that survives application restarts.
- A full command-line interface covering all core operations.

---

## 7. Testing

The project was tested using Go's built-in testing package (`go test`). Unit tests were developed to verify the correctness of hashing, mining, transaction validation, and blockchain integrity.

### 7.1 Test Areas

- **Hash Generation** — identical blocks produce identical hashes; modifying any block field produces a different hash.
- **Mining** — mining at a specified difficulty produces a hash with the required number of leading zeros, correctly stored in the block.
- **Blockchain Validation** — an untampered blockchain passes all integrity checks (hash, linkage, Proof-of-Work, timestamp order).
- **Tamper Detection** — modifying a transaction amount, a PreviousHash field, or a block timestamp each causes validation to fail, due respectively to a hash mismatch, broken linkage, or invalid timestamp sequence.
- **Transaction Validation** — negative amounts, zero amounts, insufficient balance, and pending double-spend transactions are correctly rejected, while valid transactions are accepted.

---

## 8. Results and Discussion

### 8.1 Build and Blockchain Creation

The project compiled successfully (`go build ./...`) and correctly created a Genesis Block with Index 0, PreviousHash "0", and no transactions, confirming a deterministic starting point for the chain.

### 8.2 Transaction and Mining Results

Transactions such as `faucet → alice: 100` were validated and correctly moved from the pending pool into a mined block. Mining successfully searched for valid nonces satisfying the configured difficulty — for example, at difficulty 4 the system required a hash beginning with four zeros, found after searching nonce values (nonce = 478 in one test run, producing hash `0008de9a7832f67e...`).

### 8.3 Ledger and Validation Results

The ledger correctly derived balances from transaction history (e.g. Alice = 60, Bob = 40 after faucet → Alice: 100 and Alice → Bob: 40). Blockchain validation correctly reported the chain as VALID when untampered, and correctly detected tampering — changing a transaction amount from 40 to 400 caused validation to fail with a reported block hash mismatch at the affected block index.

### 8.4 Transaction Security Results

All transaction security tests passed:

| Test Case | Expected Result | Status |
|---|---|---|
| Negative amount transaction | Reject | Passed |
| Zero amount transaction | Reject | Passed |
| Sending more coins than balance | Reject | Passed |
| Pending transaction double spending | Reject | Passed |
| Valid transaction | Accept | Passed |

### 8.5 Storage and Overall Discussion

The JSON storage system correctly saved blockchain data to `blockchain.json` and restored it after an application restart, confirming that the persistence mechanism reliably maintains state across executions. All unit tests passed (`go test ./...`) across the blockchain, mining, and utils packages.

Overall, the results confirm that the developed system satisfies its objectives, successfully demonstrating the core mechanisms of real blockchain platforms — secure hashing, Proof-of-Work mining, transaction validation, integrity checking, ledger management, and persistent storage — within a simplified, educational implementation.

---

## 9. Future Improvements

Possible future enhancements include:

- Digital signatures for authenticating transactions.
- Wallet management for key and address handling.
- Peer-to-peer networking between multiple nodes.
- Smart contracts for programmable transaction logic.
- A REST API and a graphical user interface.
- Alternative consensus algorithms beyond Proof-of-Work.
- Blockchain synchronization across multiple nodes.

---

## 10. Conclusion

The Toy Blockchain and Ledger Simulator successfully demonstrates the fundamental concepts of blockchain technology through a modular implementation in Go, covering transaction validation, SHA-256 hashing, Proof-of-Work mining, blockchain verification, ledger management, and persistent storage. Developing this simulator provided practical experience with blockchain architecture, cryptographic hashing, data integrity, and modular software design, and it can serve as a strong educational foundation for more advanced blockchain systems.
