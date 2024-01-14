package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	dirPath := "fake_drive"
	aesKey := []byte("u@QCuh@TFE~2L'Gz") // 16 bytes * 8 = 128 bit
	startTime := time.Now()

	fmt.Printf("Decrypting '%s' with AES-128...\n", dirPath)

	// Walk through all files and folders inside provided dirPath
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// Check if it is not a directory (meaning it is a file)
		// and check if it has a target file extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".enc") {
			fmt.Printf("[!] Decrypting '%s'... ", info.Name())

			// Read data from a file
			data, _ := os.ReadFile(path)

			// Create new AES cipher in GCM mode
			aesKeyCypher, _ := aes.NewCipher(aesKey)
			gcm, _ := cipher.NewGCM(aesKeyCypher)

			// Separate nonce and cipherText
			nonce, encryptedData := data[:gcm.NonceSize()], data[gcm.NonceSize():]

			// Decrypt data
			originalData, _ := gcm.Open(nil, nonce, encryptedData, nil)

			// Recreate original file where decrypted data will be stored
			newFileName := strings.TrimSuffix(path, ".enc")
			ogFile, _ := os.Create(newFileName)
			defer ogFile.Close()

			// Write decrypted data to a file
			ogFile.Write(originalData)

			// Remove encrypted file
			os.Remove(path)

			fmt.Println("DONE")
		}

		return err
	})

	if err != nil {
		fmt.Println("Error walking through the directory path: ", err)
		return
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("------------------------------------------")
	fmt.Println("Decryption finished successfully!")
	fmt.Println("Duration:", duration)
}
