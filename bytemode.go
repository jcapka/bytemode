package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	c := 500000
	bytes := make([]byte, c)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

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

	for key, value := range counts {
    	fmt.Println("Key:", key, "Value:", value)
	}
	
	fmt.Println("The byte", bytes[maxSeen], "was seen", counts[bytes[maxSeen]], "times")
}
