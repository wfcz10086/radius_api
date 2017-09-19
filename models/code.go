package models
import (
	b64 "encoding/base64"
)


func Base64_de(Base64str string) string {
	sDec, _ := b64.StdEncoding.DecodeString(Base64str)
	return string(sDec)
}
