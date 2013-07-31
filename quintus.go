package quintus

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type Quintus struct {
	Name  string
	Date  string
	Times int
}

func checkerror(err error, comment string) {
  if err != nil {
    log.Fatalln("[ERROR]", comment, err)
  }
}

func Write(somequintus []Quintus) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	encoder.Encode(somequintus)

	err := ioutil.WriteFile("quintus_gob", buffer.Bytes(), 0600)

	checkerror(err, "WRITE")
}

func Read() []Quintus {
	rawfile, err := ioutil.ReadFile("quintus_gob")

  checkerror(err, "READ")

	buffer := bytes.NewBuffer(rawfile)
	decoder := gob.NewDecoder(buffer)

	somequintus := make([]Quintus, 5)

	err = decoder.Decode(&somequintus)

  checkerror(err, "DECODE")
  
  return somequintus
}
