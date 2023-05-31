package helper

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Base64Encode(s string) string {
	en1 := base64.StdEncoding.EncodeToString([]byte(s))
	x := reverse(en1)
	en2 := x + "." + base64.StdEncoding.EncodeToString([]byte("sercet"))
	return base64.StdEncoding.EncodeToString([]byte(en2))
}

func Base64Decode(s string) (string, error) {
	dec1, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	dec2 := strings.Split(string(dec1), ".")
	y := reverse(dec2[0])
	decode, err := base64.StdEncoding.DecodeString(y)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(decode), err
}

func reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
