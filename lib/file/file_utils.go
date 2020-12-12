package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(b), "\n", " ", -1)
}
