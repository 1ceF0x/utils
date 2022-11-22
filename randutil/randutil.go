package randutil

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"time"
)

const (
	number            = "0123456789"
	letter            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterLower       = "abcdefghijklmnopqrstuvwxyz"
	letterUpper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberLetter      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberLetterLower = "0123456789abcdefghijklmnopqrstuvwxyz"
	numberLetterUpper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// 从字符串里获取长度随机字符
func RandFromChoices(choices string, length int) string {
	var container string
	b := bytes.NewBufferString(choices)
	bLength := b.Len()
	bigInt := big.NewInt(int64(bLength))
	for i := 0; i < length; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(choices[randomInt.Int64()])
	}
	return container
}

// 随机大小写字母
func RandString(length int) string {
	return RandFromChoices(letter, length)
}

// 随机小写字母
func RandStringLower(length int) string {
	return RandFromChoices(letterLower, length)
}

// 随机大写字母
func RandStringUpper(length int) string {
	return RandFromChoices(letterUpper, length)
}

// 随机数字
func RandNumber(length int) string {
	return RandFromChoices(number, length)
}

// 随机大小写字母和数字
func RandNumberString(length int) string {
	return RandFromChoices(numberLetter, length)
}

// 随机小写字母和数字
func RandNumberStringLower(length int) string {
	return RandFromChoices(numberLetterLower, length)
}

// 随机大写字母和数字
func RandNumberStringUpper(length int) string {
	return RandFromChoices(numberLetterUpper, length)
}

// 随机User-Agent
func RandUserAgent() string {
	userAgent := [...]string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	}
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(userAgent))))
	return userAgent[randomInt.Int64()]
}

// 随机sleep
func RandSleep(millisencond int) {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(millisencond)))
	ms := millisencond + int(randomInt.Int64())
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
