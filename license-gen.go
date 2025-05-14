// license-gen.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func generateLicense(deviceHash string) string {
	hash := sha256.Sum256([]byte(deviceHash))
	return hex.EncodeToString(hash[:])[:32] // ambil 32 karakter awal
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "generate" {
		fmt.Println("Usage: license-gen.go generate <device-hash>")
		os.Exit(1)
	}
	deviceHash := os.Args[2]
	license := generateLicense(deviceHash)
	filename := fmt.Sprintf("license_%s.key", deviceHash[:8])
	err := os.WriteFile(filename, []byte(license), 0644)
	if err != nil {
		fmt.Printf("Failed to write license key: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("License key generated: %s\n", filename)
}
