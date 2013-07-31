package quintus

import (
	"fmt"
	"quintus"
	"testing"
)

func TestWrite(t *testing.T) {
	someq := make([]quintus.Quintus, 1, 50)

	someq[0].Name = "Dude"
	someq[0].Date = "Abides"
	someq[0].Times = 42

	quintus.Write(someq)

	some_result := quintus.Read()
	fmt.Println("[INFO]", some_result)

	testq := quintus.Quintus{"Dude", "Abides", 42}

	if someq[0] != testq {
		t.Error("[ERROR]", "EQUALITY", someq[0], testq)
	}
}
