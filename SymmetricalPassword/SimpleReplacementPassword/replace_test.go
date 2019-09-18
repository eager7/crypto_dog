package replacement

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	cleartext := "I love you, Chain!"
	rMap := NewReplaceMap()
	fmt.Println(rMap)
	ciphered, err := Encryption(rMap, []byte(cleartext))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ciphered:", string(ciphered))
	text, err := Decryption(rMap, ciphered)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("text:", string(text))
}
