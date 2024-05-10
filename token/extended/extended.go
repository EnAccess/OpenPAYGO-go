package extended

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/dchest/siphash"
)

const (
	maxBase            = 999999
	maxActivationValue = 999999
	tokenValueOffset   = 1000000
)

type tokenType uint8

func getTokenBase(code uint64) uint64 {
	return code % tokenValueOffset
}

func putBaseInToken(token, tokenbase uint64) (uint64, error) {
	if tokenbase > maxBase {
		return 0, fmt.Errorf("invalid value")
	}

	return token - getTokenBase(token) + tokenbase, nil
}

func generateNextToken(lastCode uint64, key []byte) uint32 {
	conformedToken := make([]byte, 8)

	binary.LittleEndian.PutUint64(conformedToken, lastCode)

	return convertHashToToken(generateHash(key, conformedToken))
}

func convertHashToToken(hash uint64) uint32 {
	return convertTo40Bits(uint64(hash))
}

func convertTo40Bits(source uint64) uint32 {
	// TODO: fix this mask
	// mask := ((uint64(1) << (64 - 24 + 1)) - 1) << 24 // value should be  0x1ffffffffff000000

	// TODO: this mask value is not ther right one
	mask := ^uint64(0)

	temp := (source & mask) >> 24
	if temp > 999999999999 {
		temp = temp - 99511627777
	}

	return uint32(temp)
}

func generateHash(key []byte, token []byte) uint64 {
	h := siphash.New(key)
	h.Write(token)

	return h.Sum64()
}

func loadSecretKeyFromHex(hexKey string) ([]byte, error) {
	secretKey, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf(
			"the secret key provided is not correctly formatted, it should be 32 hexadecimal characters: %w", err)
	}

	return secretKey, nil

}

func generateStartingCde(key []byte) uint32 {
	return convertHashToToken(generateHash(key, key))
}

func convertTo4DigitsToken(source uint64) string {
	var restrictedDigitToken strings.Builder

	bitArray := getBitArrayFromInt(source, 40)

	for i := range 20 {
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
