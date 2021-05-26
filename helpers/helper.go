package helpers

import (
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

func Reader(fl string) []byte {
	ans, err := ioutil.ReadFile(fl)
	if err != nil {
		log.Fatalln(err)
	}
	return ans
}
