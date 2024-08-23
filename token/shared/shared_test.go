package shared

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	key := []byte("0123456789ABCDEF")
	hash := generateHash(key, []byte(""))

	expectedHash := uint64(3627314469837380007)
	if hash != expectedHash {
		t.Errorf("Expected hash to be %d, got %d", expectedHash, hash)
	}
}

func TestConvertHashToToken(t *testing.T) {
	key, err := hex.DecodeString("bc41ec9530f6dac86b1a29ab82edc5fb")
	if err != nil {
		log.Fatal(err)
	}
	hash := generateHash(key, []byte("hello world"))
	fmt.Printf("hash: %x\n", hash)
	token := convertHashToToken(hash)
	expectedToken := uint32(184900559)
	if token != expectedToken {
		t.Errorf("Expected token to be %d, got %d", expectedToken, token)
	}
}

func TestGenerateNextToken(t *testing.T) {
	startingCode := uint32(516959010)
	key, err := hex.DecodeString("bc41ec9530f6dac86b1a29ab82edc5fb")
	if err != nil {
		log.Fatal(err)
	}
	nextToken := GenerateNextToken(startingCode, key)
	expectedNextToken := uint32(117642353)
	if nextToken != expectedNextToken {
		t.Errorf("Expected next token to be %d, got %d", expectedNextToken, nextToken)
	}
}
