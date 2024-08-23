package token

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/EnAccess/OpenPAYGO-go/token/shared"
	"io"
	"os"
	"testing"
)

type TTokenContext struct {
	Context TokenContext
	Token   string
}

type TokenJSON struct {
	SerialNumber       string `json:"serial_number"`
	StartingCode       uint32 `json:"starting_code"`
	Key                string `json:"key"`
	TokenCount         uint8  `json:"token_count"`
	RestrictedDigitSet bool   `json:"restricted_digit_set"`
	TokenType          int    `json:"token_type"`
	ValueRaw           int    `json:"value_raw"`
	Token              string `json:"token"`
}

func getTestTokens() ([]TTokenContext, error) {
	file, err := os.Open("sample_tokens.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []TTokenContext{}, nil
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []TTokenContext{}, nil
	}

	var tokens []TokenJSON
	if err := json.Unmarshal(data, &tokens); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return []TTokenContext{}, nil
	}

	var tokenContexts []TTokenContext
	for _, t := range tokens {
		keyBytes, err := hex.DecodeString(t.Key)
		if err != nil {
			fmt.Println("Error decoding key:", err)
			continue
		}
		tc := TokenContext{
			Key:              keyBytes,
			Count:            t.TokenCount,
			Value:            t.ValueRaw,
			TokenType:        shared.TokenType(t.TokenType),
			StartCode:        t.StartingCode,
			ValueDivider:     1, // You can adjust this as needed
			RestrictDigitSet: t.RestrictedDigitSet,
			ExtendToken:      false, // You can adjust this as needed
		}
		tokenContexts = append(tokenContexts, TTokenContext{Context: tc, Token: t.Token})
	}
	return tokenContexts, nil
}
func TestGenerateToken(t *testing.T) {
	tokenCxts, _ := getTestTokens()

	for _, tokenCxt := range tokenCxts {

		token, err := generateToken(tokenCxt.Context)
		if err != nil {
			t.Errorf("Error generating token: %s", err.Error())
			continue
		}

		if token.Token != tokenCxt.Token {
			t.Errorf("Error generating token: expected %s, got %s", tokenCxt.Token, token.Token)
		}

	}
}
