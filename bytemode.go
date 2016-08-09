package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	
	bytes, _ := generateBytes(500000)
	b, c := getByte(bytes)
	fmt.Println("The byte", b, "was seen", c, "times")
}

func getByte(bytes []byte) (byte, int){
	

	maxSeenCount := 0
	maxSeen := -1

	counts := make(map[byte]int)

	for i := 0; i < len(bytes); i++ {
		count , found := counts[bytes[i]]
		if found {
			counts[bytes[i]]++
		} else {
			counts[bytes[i]] = 1
		}

		if count > maxSeenCount {
			maxSeen  = i
			maxSeenCount = count
		}
	}
	return bytes[maxSeen], counts[bytes[maxSeen]]
}

func generateBytes(length int) ([]byte, error) {
	
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

