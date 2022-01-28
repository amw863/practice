package convert

import "fmt"
import "github.com/spaolacci/murmur3"

const base62Str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func convertBase8(i uint64) []uint64 {
	var c []rune
	for i != 0 {
		c = append(c, rune(i%8))
		i = i / 8
	}

	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	fmt.Println(string(c))
	return nil
}

// 10 进制转62进制
func convertBase62(i uint64) []byte {
	var char []byte
	for i != 0 {
		char = append(char, base62Str[i%62])
		i = i / 62
	}

	// 转换算法
	for i, j := 0, len(char)-1; i < j; i, j = i+1, j-1 {
		char[i], char[j] = char[j], char[i]
	}

	//fmt.Println(string(char))
	return char
}

func HashMurmur3() {
	h := murmur3.New32()
	h.Write([]byte("1111111"))
	fmt.Println(h.Sum32())
	mu := murmur3.Sum32([]byte("1111111"))
	fmt.Println(mu)
}
