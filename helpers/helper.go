package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func Reader(s string) []byte {
	ans, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Error")
	}
	return (ans)
}

func Unmarshal(r []byte, p *map[string]interface{}) {
	err := json.Unmarshal(r, p)
	if err != nil {
		log.Fatalln(err)
	}
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Encode64(s []byte) string {
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	return s64
}
func Decode64(s string) string {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bs)
}

