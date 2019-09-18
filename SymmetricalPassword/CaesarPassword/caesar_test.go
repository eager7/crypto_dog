package caesar

import (
	"fmt"
	"testing"
)

func TestEngine(t *testing.T) {
	fmt.Println('a', int('a'))
	fmt.Println(code(1, 'a'))
	fmt.Println(code(10, 'z'))

	fmt.Println(decode(1, 'b'))
	fmt.Println(decode(10, 38))
}

func TestCaesar(t *testing.T) {
	cleartext := "I love you, Chain!"
	ciphered, err := Encryption(30, []byte(cleartext))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ciphered:", string(ciphered))
	text, err := Decryption(30, ciphered)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("text:", string(text))
}
