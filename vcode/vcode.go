package vcode

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"
)

// GenRandomBiliVCode generates a random vcode for bilidanmuauth
// params: 1. client_uuid 2.uid 3. prefix 4. length
// return: 1. vcode
func GenRandomBiliVCode(clientUUID string, uid int, prefix string, length uint) string {
	data := "bili" + clientUUID + string(rune(uid)) + strconv.FormatInt((time.Now().UnixNano()), 10)
	hashString := hashSHA1(data)
	if prefix == "" {
		return "vc-" + hashString[:length]
	}
	return prefix + hashString[:length]
}

// GenBiliVCode generates a fixed vcode via extra infomation for bilidanmuauth
// Could be used to generate a vcode for a specific user or check the vcode
// params: 1. client_uuid 2.uid 3. extraInfo 4. prefix 5. length
// return: 1. vcode
func GenBiliVCodeWithExtraInfo(clientUUID string, uid string, extraInfo string, prefix string, length uint) string {
	data := "bili" + clientUUID + uid + extraInfo
	hashString := hashSHA1(data)
	if prefix == "" {
		return "vc-" + hashString[:length]
	}
	return prefix + hashString[:length]
}

func hashSHA1(s string) string {
	res := sha1.Sum([]byte(s))
	result := hex.EncodeToString(res[:])
	return result
}
