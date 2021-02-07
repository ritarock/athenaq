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

func MakeDir(path string) {
	err := os.Mkdir(path, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func Write(path, data string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(([]byte)(data))
}
