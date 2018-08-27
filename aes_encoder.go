package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AesNewCBCDecrypter(enctext, akey string) string {
	key := []byte(akey)
	//ciphertext, _ := hex.DecodeString(enctext)

	ciphertext, _ := base64.StdEncoding.DecodeString(enctext)

	m := len(key)
	if m < 16 {
		key = append(key, make([]byte, 16-m)...)
	}
	//fmt.Printf("KEY:%x\n", key)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		//panic("ciphertext too short")
		return "ciphertext too short"
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.

	if len(ciphertext)%aes.BlockSize != 0 {
		//panic("ciphertext is not a multiple of the block size")
		return "ciphertext is not a multiple of the block size"
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.
	//fmt.Printf("TXT:%x\n", ciphertext)
	//fmt.Printf("%s\n", ciphertext)

	return string(ciphertext)
}
func AesNewCBCEncrypter(text, akey string) string {
	key := []byte(akey)
	plaintext := []byte(text)

	m := len(key)
	if m < 16 {
		key = append(key, make([]byte, 16-m)...)

	}

	m = len(plaintext) % aes.BlockSize
	if m != 0 {
		//x :=
		plaintext = append(plaintext, make([]byte, aes.BlockSize-m)...)

	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
 

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

//    for i:=0; i<16 ;i++ {
//		iv[i]=key[i]
//	}
 
  
 
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("KEY:%x\n", key)
	fmt.Printf("IV :%x\n", iv)
	fmt.Printf("TXT:%x\n", plaintext)
	fmt.Printf("ENC:%x\n", ciphertext)
	fmt.Printf("B64:%s\n", base64.StdEncoding.EncodeToString(ciphertext))
	return base64.StdEncoding.EncodeToString(ciphertext)

}

 
func  _main() {

 
	key := "1234567890123456"
	s := "ThisThisThisThisThisThis"
	s="Edit1"
	es := AesNewCBCEncrypter(s, key)
 

	fmt.Println(AesNewCBCDecrypter(es, key))

}
