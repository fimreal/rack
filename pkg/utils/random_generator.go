package utils

import (
	"hash/fnv"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func HashStr2Int(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}

func GenSeed() uint64 {
	return uint64(HashStr2Int(uuid.NewString()))
}

// RandomString generates a random string of specified length from the given composition.
func RandomString(composition string, length int) string {
	r := rand.New(rand.NewSource(GenSeed()))

	b := make([]rune, length)
	compositionRune := []rune(composition)
	for i := 0; i < length; i++ {
		inx := r.Intn(len(compositionRune))
		b[i] = compositionRune[inx]
	}
	return string(b)
}

// SixNumber generates a random six number string.
func SixNumber() string {
	return RandomString("0123456789", 6)
}

func GenPassword(length int) string {
	return RandomString("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM+=-@#.$%^*", length)
}
