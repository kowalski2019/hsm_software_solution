package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	aux "ssm/auxiliary"
)

var orig_key = aux.GetFromConfig("constants.crypto_key")
var auth_key = aux.GetFromConfig("constants.auth_key")

func Encryption(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("auth_key") == auth_key {
		if r.FormValue("crypto_type") == "encode" {
			Encode(w, r, r.FormValue("text"))
		} else if r.FormValue("crypto_type") == "decode" {
			Decode(w, r, r.FormValue("text"))
		} else {
			// nothing to do StatusBadRequest
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Bad Request"}`))

		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Unauthorized"}`))

	}

}

func Encode(w http.ResponseWriter, r *http.Request, text string) {
	encrypted, err := encrypt(text, orig_key, w)
	//fmt.Printf("encrypted : %s\n", encrypted)
	if err != nil {
		errorDuringEncryption(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"encrypted": "` + encrypted + `"}`))
}
func Decode(w http.ResponseWriter, r *http.Request, text string) {
	decrypted, err := decrypt(text, orig_key, w)
	//fmt.Printf("decrypted : %s\n", decrypted)
	if err != nil {
		errorDuringEncryption(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"decrypted": "` + decrypted + `"}`))
}

func errorDuringEncryption(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "` + t + `"}`))
}
func encrypt(stringToEncrypt string, keyString string, w http.ResponseWriter) (encryptedString string, er error) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("[Encoding Error] %s", err)
	}
	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("[Encoding Error] %s", err)

	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("[Encoding Error] %s", err)
	}
	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix
	//to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func decrypt(encryptedString string, keyString string, w http.ResponseWriter) (decryptedString string, er error) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("[Decoding Error] %s", err)
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("[Decoding Error] %s", err)
	}
	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("[Decoding Error] %s", err)
	}

	return fmt.Sprintf("%s", plaintext), nil
}
