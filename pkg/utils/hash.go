package utils

import (
	"hash/fnv"
	"math/rand"

	"github.com/google/uuid"
)

func HashStr2Int(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}

func GenSeed() int64 {
	return int64(HashStr2Int(uuid.NewString()))
}

// 例如生成随机 6 个数字：
// RandomString("1234567890", 6)
func RandomString(composition string, length int) string {
	rand.Seed(GenSeed())

	b := make([]rune, length)
	compositionRune := []rune(composition)
	for i := 0; i < length; i++ {
		inx := rand.Intn(len(compositionRune))
		b[i] = compositionRune[inx]
	}
	return string(b)
}
