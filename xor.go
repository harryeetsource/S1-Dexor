package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func xor(data, key []byte) []byte {
	decrypted := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		decrypted[i] = data[i] ^ key[i%len(key)]
	}
	return decrypted
}

func decryptFile(path string, key []byte) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	decrypted := xor(data, key)

	newPath := path + "_decrypted"
	err = ioutil.WriteFile(newPath, decrypted, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	hexKey := "FF" // replace with your hex-encoded key
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Fatalf("Invalid key: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := filepath.Join(dir, file.Name())
		err := decryptFile(filename, key)
		if err != nil {
			log.Printf("Error decrypting file %s: %v", filename, err)
		} else {
			log.Printf("Successfully decrypted file %s", filename)
		}
	}
}
