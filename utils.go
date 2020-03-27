package main

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"time"
)

func GetCurrentLocalTime() time.Time {
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		log.Println(err)
		return now
	}
	return now.In(loc)
}
func ComputeSHA1(plain string) string {
	plainBytes := []byte(plain)
	hashFunc := sha1.New()
	hashFunc.Write(plainBytes)
	hashStr := base64.URLEncoding.EncodeToString(hashFunc.Sum(nil))
	return hashStr
}
