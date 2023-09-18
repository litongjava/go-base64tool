package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func encodeToBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func decodeFromBase64(encodedData string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go-base64tool encode <data>")
		fmt.Println("  go-base64tool decode <encoded_data>")
		return
	}

	action := os.Args[1]
	data := os.Args[2]

	switch action {
	case "encode":
		fmt.Println(encodeToBase64(data))
	case "decode":
		decodedData, err := decodeFromBase64(data)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return
		}
		fmt.Println(decodedData)
	default:
		fmt.Println("Invalid action. Use 'encode' or 'decode'.")
	}
}
