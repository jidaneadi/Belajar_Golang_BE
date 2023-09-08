package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Penamaan fungsi hrs diawali Huruf Kapital
func EncryptHash(data string) string {

	//Membuat variabel untuk memilih algoritma sha256
	hasher := sha256.New()

	//Mengkonversi variabel mejadi tipe data byte dan akan dikirim ke hasher utk menghitung hash dr byte variabel
	hasher.Write([]byte(data))

	//Tambahan data
	additionalData := "Inikunciawal"

	//Menambahkan data pada hasher
	hasher.Write([]byte(additionalData))

	//Menghasilkan hash dlm btk byte
	hash := hasher.Sum(nil)

	//Mengubah hash mjd string hexadesimal
	hashString := hex.EncodeToString(hash)

	return hashString
}

// func CompareHash(data string, data2 string) string {

// 	cekPassword := data

// }

// Fungsi untuk mengenkripsi data menggunakan SHA256 dan RSA
// func encryptData(data string, publicKey *rsa.PublicKey) (string, error) {
// 	// Generate SHA256 hash dari data
// 	hash := sha256.Sum256([]byte(data))

// 	// Enkripsi hash menggunakan kunci publik RSA
// 	encryptedHash, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, hash[:])
// 	if err != nil {
// 		return "", fmt.Errorf("failed to encrypt data: %v", err)
// 	}

// 	// Convert the encrypted hash to hexadecimal string
// 	encryptedHashStr := hex.EncodeToString(encryptedHash)

// 	return encryptedHashStr, nil
// }

func EncryptRSA(data string) (string, error) {
	// Membuat kunci RSA baru
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	// Mengenkripsi data menggunakan algoritma SHA256
	hashedData := sha256.Sum256([]byte(data))

	// Mengenkripsi hashedData dengan menggunakan kunci privat RSA
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, hashedData[:], nil)
	if err != nil {
		return "", err
	}

	// Mengubah data yang dienkripsi menjadi base64 string
	encryptedDataString := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedDataString, nil
}

func DecryptRSA(encryptedData string, privateKey *rsa.PrivateKey) (string, error) {
	// Mengubah data yang dienkripsi dari base64 string menjadi byte
	encryptedDataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// Mendeskripsi data menggunakan kunci privat RSA
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedDataBytes, nil)
	if err != nil {
		return "", err
	}

	// Mengubah hasil deskripsi menjadi string
	decryptedDataString := string(decryptedData)

	return decryptedDataString, nil
}
