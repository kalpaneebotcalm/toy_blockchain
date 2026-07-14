package blockchain_test

import (
	"testing"
	"toy-blockchain/block"
	"toy-blockchain/blockchain"
)

func TestDifficultyRemainsSame(t *testing.T) {
	// Start with blocks having normal block times (60 seconds apart)
	// Base difficulty = 4
	blocks := []block.Block{
		{Index: 0, Timestamp: 1000, Difficulty: 4},
		{Index: 1, Timestamp: 1060, Difficulty: 4},
		{Index: 2, Timestamp: 1120, Difficulty: 4},
		{Index: 3, Timestamp: 1180, Difficulty: 4},
		{Index: 4, Timestamp: 1240, Difficulty: 4},
	}

	nextDiff := blockchain.CalculateNextDifficulty(blocks)
	if nextDiff != 4 {
		t.Errorf("Expected difficulty to remain 4, got %d", nextDiff)
	}
}

func TestDifficultyIncrease(t *testing.T) {
	// Very fast block creation (1 second apart)
	// expectedTime = 5 * 60 = 300 seconds
	// actualTime = 1004 - 1000 = 4 seconds
	// 4 < 150 (expectedTime / 2), so difficulty should increase from 4 to 5
	blocks := []block.Block{
		{Index: 0, Timestamp: 1000, Difficulty: 4},
		{Index: 1, Timestamp: 1001, Difficulty: 4},
		{Index: 2, Timestamp: 1002, Difficulty: 4},
		{Index: 3, Timestamp: 1003, Difficulty: 4},
		{Index: 4, Timestamp: 1004, Difficulty: 4},
	}

	nextDiff := blockchain.CalculateNextDifficulty(blocks)
	if nextDiff != 5 {
		t.Errorf("Expected difficulty to increase to 5, got %d", nextDiff)
	}
}

func TestDifficultyDecrease(t *testing.T) {
	// Very slow block creation (300 seconds apart)
	// expectedTime = 300 seconds
	// actualTime = 2200 - 1000 = 1200 seconds
	// 1200 > 600 (expectedTime * 2), so difficulty should decrease from 4 to 3
	blocks := []block.Block{
		{Index: 0, Timestamp: 1000, Difficulty: 4},
		{Index: 1, Timestamp: 1300, Difficulty: 4},
		{Index: 2, Timestamp: 1600, Difficulty: 4},
		{Index: 3, Timestamp: 1900, Difficulty: 4},
		{Index: 4, Timestamp: 2200, Difficulty: 4},
	}

	nextDiff := blockchain.CalculateNextDifficulty(blocks)
	if nextDiff != 3 {
		t.Errorf("Expected difficulty to decrease to 3, got %d", nextDiff)
	}
}

func TestDifficultyLimits(t *testing.T) {
	// Verify difficulty never goes below 1
	// Start with difficulty 1, and make blocks very slow so difficulty tries to decrease.
	blocksLow := []block.Block{
		{Index: 0, Timestamp: 1000, Difficulty: 1},
		{Index: 1, Timestamp: 1300, Difficulty: 1},
		{Index: 2, Timestamp: 1600, Difficulty: 1},
		{Index: 3, Timestamp: 1900, Difficulty: 1},
		{Index: 4, Timestamp: 2200, Difficulty: 1},
	}

	nextDiffLow := blockchain.CalculateNextDifficulty(blocksLow)
	if nextDiffLow != 1 {
		t.Errorf("Expected difficulty to be clamped to 1, got %d", nextDiffLow)
	}

	// Verify difficulty never exceeds 10
	// Start with difficulty 10, and make blocks very fast so difficulty tries to increase.
	blocksHigh := []block.Block{
		{Index: 0, Timestamp: 1000, Difficulty: 10},
		{Index: 1, Timestamp: 1001, Difficulty: 10},
		{Index: 2, Timestamp: 1002, Difficulty: 10},
		{Index: 3, Timestamp: 1003, Difficulty: 10},
		{Index: 4, Timestamp: 1004, Difficulty: 10},
	}

	nextDiffHigh := blockchain.CalculateNextDifficulty(blocksHigh)
	if nextDiffHigh != 10 {
		t.Errorf("Expected difficulty to be clamped to 10, got %d", nextDiffHigh)
	}
}
