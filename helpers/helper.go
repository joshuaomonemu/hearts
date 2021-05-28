package helpers

import (
	"encoding/json"
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

func Unmarshal(r []byte, p *map[string]interface{}) {
	err := json.Unmarshal(r, p)
	if err != nil {
		log.Fatalln(err)
	}
}
