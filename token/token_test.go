package token

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/EnAccess/OpenPAYGO-go/token/shared"
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
	data := `[
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 3,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 1,
    "token": "380589011"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 5,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 2,
    "token": "283675012"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 7,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 5,
    "token": "034254015"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 9,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 995,
    "token": "409152005"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 11,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 998,
    "token": "071763008"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 13,
    "restricted_digit_set": false,
    "token_type": 1,
    "value_raw": 999,
    "token": "814704009"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 14,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 1,
    "token": "141465011"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 16,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 2,
    "token": "448320012"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 18,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 5,
    "token": "730651015"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 20,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 995,
    "token": "132820005"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 22,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 998,
    "token": "146345008"
  },
  {
    "serial_number": "TEST220000001",
    "starting_code": 516959010,
    "key": "bc41ec9530f6dac86b1a29ab82edc5fb",
    "token_count": 24,
    "restricted_digit_set": false,
    "token_type": 2,
    "value_raw": 999,
    "token": "386863009"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 3,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 1,
    "token": "413441444234331"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 5,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 2,
    "token": "431131331113332"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 7,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 5,
    "token": "423424444232241"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 9,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 995,
    "token": "422313413112333"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 11,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 998,
    "token": "231434142221342"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 13,
    "restricted_digit_set": true,
    "token_type": 1,
    "value_raw": 999,
    "token": "242313431134143"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 14,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 1,
    "token": "113434333414311"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 16,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 2,
    "token": "414212121322332"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 18,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 5,
    "token": "413424224321241"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 20,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 995,
    "token": "342124322343233"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 22,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 998,
    "token": "211422314241142"
  },
  {
    "serial_number": "TEST240000002",
    "starting_code": 432435255,
    "key": "dac86b1a29ab82edc5fbbc41ec9530f6",
    "token_count": 24,
    "restricted_digit_set": true,
    "token_type": 2,
    "value_raw": 999,
    "token": "331233113332423"
  }
]
`
	var tokens []TokenJSON
	if err := json.Unmarshal([]byte(data), &tokens); err != nil {
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
