package main

import "fmt"

func add(a, b byte) byte {
	sum := a + (b - 'A')
	if sum > 'Z' {
		return sum - 26
	}
	return sum
}

func sub(a, b byte) byte {
	difference := a - (b - 'A')
	if difference < 'A' {
		return difference + 26
	}
	return difference
}

func encryptWithConstitution(m []byte) {
	text := "WETHEPEOPLEOFTHE"
	out := make([]byte, len(m))
	for i := range m {
		out[i] = add(m[i], text[i])
	}
	fmt.Println(string(out))
}

func decryptWithConstitution(m []byte) []byte {
	text := "WETHEPEOPLEOFTHE"
	out := make([]byte, len(m))
	for i := range m {
		out[i] = sub(m[i], text[i])
	}
	fmt.Println(string(out))
	return out
}

func main() {
	// encryptWithConstitution([]byte("OMARCOMIN"))
	decryptWithConstitution([]byte("KQTYGDQWC"))
}
