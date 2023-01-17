package updategpgssl

import (
	"fmt"
	"io/ioutil"
	"log"
	"sumaxmlrpc/auth"
	"sumaxmlrpc/request"

	"github.com/divan/gorilla-xmlrpc/xml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (u *Keys) CreateKeys(SessionKey auth.SumaSessionKey, description string, mType string, filepath string) error {
	u.SessionKey = fmt.Sprintf("%s", SessionKey.Sessionkey)
	u.Description = description
	u.Type = mType

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	u.Content = string(content)
	//fmt.Printf("file content is %s\n", u.Content)

	return nil
}

func (u *Keys) UpdateGPGSSL() error {
	method := "kickstart.keys.update"

	buf, err := xml.EncodeClientRequest(method, u)
	if err != nil {
		log.Fatalf("Encoding error: %s\n", err)
	}

	resp, err := request.MakeRequest(buf)
	if err != nil {
		log.Fatalf("UpdateGPGSSL API error: %s\n", err)
	}

	if resp.StatusCode == 200 {
		log.Println("UpdateGPGSSL updated successfully.")
	}

	return nil
}
