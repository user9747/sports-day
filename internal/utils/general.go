package utils

import (
	"context"
	"math/rand"
	"sports-day/internal/entity"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GenerateRandomString generates and returns a random alphanumeric string (with mixed case) of size n
func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

// InArr checks for value present in a slice / array
func InArr[T string | int | int16 | int32 | int64 | bool](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func IsUniqueConstraintViolation(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func GetLoggedInUser(ctx context.Context) *entity.LoggedInUser {
	user, ok := ctx.Value("user").(*entity.LoggedInUser)
	if ok {
		return user
	}
	return nil
}
