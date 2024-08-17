package token

import (
	"errors"
	"fmt"
	"github.com/EnAccess/OpenPAYGO-go/token/extended"
	"github.com/EnAccess/OpenPAYGO-go/token/shared"
	"math"
	"strconv"
	"strings"
)

type FinalToken struct {
	Token string
	Count uint8
}

type TokenContext struct {
	Key              []byte
	Count            uint8
	Value            int
	TokenType        shared.TokenType
	StartCode        uint32
	ValueDivider     uint8
	RestrictDigitSet bool
	ExtendToken      bool
}

func generateToken(tokenContext TokenContext) (FinalToken, error) {
	startCode := tokenContext.StartCode
	value := tokenContext.Value
	valueDivider := tokenContext.ValueDivider
	tokenType := tokenContext.TokenType
	if valueDivider == 0 {
		valueDivider = 1
	}

	if startCode == 0 {
		startCode = shared.GenerateStartingCode(tokenContext.Key)
	}

	if tokenType == shared.TokenTypeAddTime || tokenType == shared.TokenTypeSetTime {
		if value == 0 {
			return FinalToken{}, errors.New("token does not have a value")
		}

		value = int(math.Round(float64(value * int(valueDivider))))
		var maxValue int
		if tokenContext.ExtendToken {
			maxValue = extended.MaxActivationValue
		} else {
			maxValue = shared.MaxActivationValue
		}

		if value > maxValue {
			return FinalToken{}, errors.New("token value provided is too high")
		}
	} else if value != 0 {
		return FinalToken{}, errors.New("a value is not allowed for this token type")
	} else {
		if tokenType == shared.TokenTypeDisablePayg {
			value = shared.PaygDisableValue
		} else if tokenType == shared.TokenTypeCounterSync {
			value = shared.CounterSyncValue
		} else {
			return FinalToken{}, errors.New("the token type provided is not supported")
		}
	}

	// TODO: add support for extended tokens
	return generateStandardToken(TokenContext{
		StartCode:        startCode,
		Key:              tokenContext.Key,
		Count:            tokenContext.Count,
		RestrictDigitSet: tokenContext.RestrictDigitSet,
		Value:            value,
		TokenType:        tokenType,
	})
}

func generateStandardToken(tokenContext TokenContext) (FinalToken, error) {

	startBaseCode := shared.GetTokenBase(uint64(tokenContext.StartCode))
	tokenBase := encodeBase(startBaseCode, tokenContext.Value)
	curToken, err := shared.PutBaseInToken(uint64(tokenContext.StartCode), tokenBase)
	if err != nil {
		return FinalToken{}, fmt.Errorf("generating standard token: %s", err.Error())
	}

	newCount := getNewCount(tokenContext.Count, tokenContext.TokenType)
	for i := 0; i < int(newCount-1); i++ {
		curToken = uint64(shared.GenerateNextToken(uint32(curToken), tokenContext.Key))
	}
	finalToken, err := shared.PutBaseInToken(curToken, tokenBase)
	if err != nil {
		return FinalToken{}, fmt.Errorf("generating standard token: %s", err.Error())
	}

	var token FinalToken
	if tokenContext.RestrictDigitSet {
		token.Token = shared.ConvertTo4DigitsToken(finalToken)
		token.Token = strings.TrimLeft(fmt.Sprintf("%015s", token.Token), " ")
	} else {
		token.Token = strconv.FormatInt(int64(finalToken), 10)
		token.Token = strings.TrimLeft(fmt.Sprintf("%09s", token.Token), " ")
	}
	token.Count = newCount
	return token, nil
}

func encodeBase(baseCode uint64, value int) uint64 {
	if uint64(value)+baseCode > 999 {
		return uint64(value) + baseCode - 1000
	}
	return uint64(value) + baseCode
}

func getNewCount(count uint8, mode shared.TokenType) uint8 {
	var newCnt uint8
	currCountOdd := count % 2

	if mode == shared.TokenTypeSetTime ||
		mode == shared.TokenTypeDisablePayg ||
		mode == shared.TokenTypeCounterSync {
		if currCountOdd != 0 {
			newCnt = count + 2
		} else {
			newCnt = count + 1
		}
	} else {
		if currCountOdd != 0 {
			newCnt = count + 1
		} else {
			newCnt = count + 2
		}
	}
	return newCnt
}
