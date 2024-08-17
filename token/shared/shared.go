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

type TokenType uint8

const (
	TokenTypeAddTime     TokenType = 1
	TokenTypeSetTime               = 2
	TokenTypeDisablePayg           = 3
	TokenTypeCounterSync           = 4
	TokenTypeInvalid               = 10
	TokenTypeAlreadyUsed           = 11
)

func GetTokenBase(code uint64) uint64 {
	return code % TokenValueOffset
}

func PutBaseInToken(token, tokenbase uint64) (uint64, error) {
	if tokenbase > MaxBase {
		return 0, fmt.Errorf("invalid value")
	}

	return token - GetTokenBase(token) + tokenbase, nil
}

func GenerateNextToken(lastCode uint32, key []byte) uint32 {
	conformedToken := make([]byte, 4)

	binary.BigEndian.PutUint32(conformedToken, lastCode)

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
	binary.BigEndian.PutUint64(binHash, hash)

	hiHash := binary.BigEndian.Uint32(binHash[0:4])
	loHash := binary.BigEndian.Uint32(binHash[4:8])

	return convertTo29BitsAndHalf(uint64(hiHash ^ loHash))
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

func ConvertTo4DigitsToken(source uint64) string {
	var restrictedDigitToken strings.Builder

	bitArray := getBitArrayFromInt(source, 4*2)

	//for i := range 4 {
	//	thisArray := bitArray[i*2 : (i*2)+2]
	//	restrictedDigitToken.WriteString(
	//		fmt.Sprint(bitArrayToInt(thisArray) + 1))
	//}

	ir := bitArrayToInt(bitArray)

	restrictedDigitToken.WriteString(fmt.Sprintf("%d", ir))

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
	bitsArray := make([]byte, nbOfBits)
	binary.BigEndian.PutUint64(bitsArray, source)

	return bitsArray
}

func bitArrayToInt(bits []byte) uint64 {
	return binary.BigEndian.Uint64(bits)
}
