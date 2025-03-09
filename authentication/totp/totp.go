package totp

import (
	"crypto/hmac"
	"encoding/binary"
	"fmt"
	"hash"
)

var digitsPower = [11]int{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
}

func Generate(
	key []byte,
	timeSteps int64,
	digits int,
	crypto func() hash.Hash,
) string {
	hmacHash := hmac.New(crypto, key)
	binary.Write(hmacHash, binary.BigEndian, timeSteps)
	hash := hmacHash.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	binary := int(binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7FFFFFFF)
	otp := fmt.Sprintf("%0*d", digits, binary%digitsPower[digits])
	return otp
}
