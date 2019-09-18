package replacement

func NewReplaceMap() map[byte]byte {
	var set = []byte(` !"#$%^&*()-+,.?/\|}{][~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ`)
	m, ms, list := make(map[byte]byte, len(set)), make(map[uint8]uint8), make([]uint8, 0)
	for _, b := range set {
		m[b] = b
	}
	for k := range m { //利用map取值是随机的这个特性打乱顺序
		list = append(list, k)
	}
	for k, v := range list {
		ms[set[k]] = v
	}
	return ms
}

func Encryption(rMap map[byte]byte, cleartext []byte) (ciphered []byte, err error) {
	for _, b := range cleartext {
		ciphered = append(ciphered, rMap[b])
	}
	return ciphered, nil
}

func Decryption(rMap map[byte]byte, ciphered []byte) (cleartext []byte, err error) {
	dMap := make(map[byte]byte, len(rMap))
	for k, v := range rMap {
		dMap[v] = k
	}
	for _, b := range ciphered {
		cleartext = append(cleartext, dMap[b])
	}
	return cleartext, nil
}
