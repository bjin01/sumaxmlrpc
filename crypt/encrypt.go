package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v3"
)

func CreateSumainfo(server string, userid string, pwd string) *Sumaconf {
	var key []byte
	if os.Getenv("SUMAPROM_ENCRYPT") != "" {
		key = []byte(os.Getenv("SUMAPROM_ENCRYPT"))
	} else {
		key = []byte("A8CA6E5155A14BB0D8F4E7CE0E23B8A2")
	}
	return &Sumaconf{
		Server: server,
		Userid: userid,
		Password: func() string {
			return encrypt(key, pwd)
		}(),
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PromptUser(create_config *string) {
	fmt.Printf("Create or update config file: %s\n", *create_config)
	var pwd []byte
	var server, userid, password string
	fmt.Println("Enter SUMA Host name: ")
	fmt.Scanln(&server)
	fmt.Println("Enter SUMA api userid: ")
	fmt.Scanln(&userid)
	fmt.Println("Enter SUMA api password: ")
	pwd, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("Password typed: " + string(pwd))
	}
	password = string(pwd)

	sumainfo := CreateSumainfo(server, userid, password)
	out, err := yaml.Marshal(sumainfo)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*create_config)
	check(err)

	_, err = f.Write(out)
	check(err)

	f.Sync()
}
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

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

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}
