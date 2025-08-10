package main

import "fmt"

func rebuildMessage(parts []string) string {
	// Map from first char to part
	firstCharMap := make(map[byte]string)
	for _, p := range parts {
		firstCharMap[p[0]] = p
	}

	// Start with the part beginning with 'A'
	message := firstCharMap['A']

	// Keep appending until we reach 'Z'
	for message[len(message)-1] != 'Z' {
		lastChar := message[len(message)-1]
		nextPart := firstCharMap[lastChar]
		// Append without duplicating the first char of the next part
		message += nextPart[1:]
	}

	return message
}

func BuildMessage() {
	fmt.Println(rebuildMessage([]string{"Ab", "bcZ"})) // AbcZ
}
