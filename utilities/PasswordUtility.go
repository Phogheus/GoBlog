package utilities

import (
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/scrypt"
)

const (
	FAILED_TO_COMPARE_ERROR string = "Failed to compare hash with input."
	N                       int    = 32768
	R                       int    = 8
	P                       int    = 1
	KEY_LEN                 int    = 32
)

var compareFailedError error = errors.New(FAILED_TO_COMPARE_ERROR)

func HashInput(input string) (string, error) {
	salt := make([]byte, 64)
	rand.Read(salt) // Read always returns len(p) and a nil error

	// Create hash using recommended values (we aren't making a real product here, so no need to be extra)
	hash, err := scrypt.Key([]byte(input), salt, N, R, P, KEY_LEN)

	if err != nil {
		return "", err
	}

	inputHash := hex.EncodeToString(hash) + hex.EncodeToString(salt)

	return inputHash, nil
}

func DoesHashedInputCompare(hashedSource string, input string) (bool, error) {
	lengthOfInputHash := len(hashedSource)

	if lengthOfInputHash <= 128 { // 64 byte array encoded to hex is 128 bytes (chars)
		return false, compareFailedError
	}

	encodedHash := hashedSource[:lengthOfInputHash-128]
	encodedSalt := hashedSource[lengthOfInputHash-128:]

	salt, err := hex.DecodeString(encodedSalt)

	if err != nil {
		return false, compareFailedError
	}

	rehash, err := scrypt.Key([]byte(input), salt, N, R, P, KEY_LEN)

	if err != nil {
		return false, compareFailedError
	}

	encodedRehash := hex.EncodeToString(rehash)

	return encodedHash == encodedRehash, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
