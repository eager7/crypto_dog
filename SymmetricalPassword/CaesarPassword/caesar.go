package caesar

import (
	"errors"
)

func Encryption(offset uint8, cleartext []byte) (ciphered []byte, err error) {
	for _, b := range cleartext {
		if dest, err := code(offset, b); err != nil {
			return nil, errors.New("code err:" + err.Error())
		} else {
			ciphered = append(ciphered, dest)
		}
	}
	return ciphered, nil
}

func Decryption(offset uint8, ciphered []byte) (cleartext []byte, err error) {
	for _, b := range ciphered {
		if src, err := decode(offset, b); err != nil {
			return nil, errors.New("decode err:" + err.Error())
		} else {
			cleartext = append(cleartext, src)
		}
	}
	return cleartext, nil
}

func code(offset uint8, src byte) (dest byte, err error) {
	if src < 32 || src > 126 {
		return src, errors.New("the char must between 32-126")
	}
	dest = src + offset
	if dest > 126 {
		dest = dest - 126 + 32
	}
	return dest, nil
}

func decode(offset uint8, dest byte) (src byte, err error) {
	if dest < 32 || dest > 126 {
		return src, errors.New("the char must between 32-126")
	}
	src = dest - offset
	if src < 32 {
		src = 126 - (32 - src)
	}
	return src, nil
}
