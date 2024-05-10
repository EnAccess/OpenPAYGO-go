package shared

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/dchest/siphash"
)

const (
	MaxBase            = 999
	MaxActivationValue = 995
	PaygDisableValue   = 998
	CounterSyncValue   = 999
	TokenValueOffset   = 1000
)

type tokenType uint8

const (
	tokenTypeAddTime     tokenType = 1
	tokenTypeSetTime               = 2
	tokenTypeDisablePayg           = 3
	tokenTypeCounterSync           = 4
	tokenTypeInvalid               = 10
	tokenTypeAlreadyUsed           = 11
)

func getTokenBase(code uint64) uint64 {
	return code % TokenValueOffset
}

func putBaseInToken(token, tokenbase uint64) (uint64, error) {
	if tokenbase > MaxBase {
		return 0, fmt.Errorf("invalid value")
	}

	return token - getTokenBase(token) + tokenbase, nil
}

func generateNextToken(lastCode uint32, key []byte) uint32 {
	conformedToken := make([]byte, 4)

	binary.LittleEndian.PutUint32(conformedToken, lastCode)

	extendedToken := append(conformedToken, conformedToken...)

	return convertHashToToken(generateHash(key, extendedToken))
}

func generateHash(key []byte, token []byte) uint64 {
	h := siphash.New(key)
	h.Write(token)

	return h.Sum64()
}

func convertHashToToken(hash uint64) uint32 {
	binHash := make([]byte, 8)
	binary.LittleEndian.PutUint64(binHash, hash)

	hiHash := binary.LittleEndian.Uint32(binHash[0:4])
	loHash := binary.LittleEndian.Uint32(binHash[4:8])

	return convertTo29BitsAndHalf(uint64((hiHash ^ loHash)))
}

func loadSecretKeyFromHex(hexKey string) ([]byte, error) {
	secretKey, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf(
			"the secret key provided is not correctly formatted, it should be 32 hexadecimal characters: %w", err)
	}

	return secretKey, nil

}

func GenerateStartingCode(key []byte) uint32 {
	startingHash := generateHash(key, key)

	return convertHashToToken(startingHash)
}

func convertTo29BitsAndHalf(source uint64) uint32 {
	//TODO: check this mask value
	mask := ((uint64(1) << (32 - 2 + 1)) - 1) << 2

	temp := (source & mask) >> 2
	if temp > 999999999 {
		temp = temp - 73741825
	}

	return uint32(temp)
}

func convertTo4DigitsToken(source uint64) string {
	var restrictedDigitToken strings.Builder

	bitArray := getBitArrayFromInt(source, 30)

	for i := range 15 {
		thisArray := bitArray[i*2 : (i*2)+2]
		restrictedDigitToken.WriteString(
			fmt.Sprint(bitArrayToInt(thisArray) + 1))
	}

	return restrictedDigitToken.String()
}

func convertFrom4DigitsToken(digits string) uint64 {
	bits := make([]byte, 4)

	for _, digit := range digits {
		digit = digit - 1
		tmp := getBitArrayFromInt(uint64(digit), 2)
		bits = append(bits, tmp...)
	}

	return bitArrayToInt(bits)
}

func getBitArrayFromInt(source uint64, nbOfBits int) []byte {
	bitsArray := make([]byte, (nbOfBits/8)+1)
	binary.LittleEndian.PutUint64(bitsArray, source)

	return bitsArray
}

func bitArrayToInt(bits []byte) uint64 {
	return binary.LittleEndian.Uint64(bits)
}
