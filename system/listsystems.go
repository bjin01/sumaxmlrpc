package system

import (
	"fmt"
	"log"
	"sumaxmlrpc/auth"
	"sumaxmlrpc/request"
	"time"

	"github.com/divan/gorilla-xmlrpc/xml"
)

func (l *ListActiveSystem) GetActiveSystems(Sessionkey *auth.SumaSessionKey) error {
	method := "system.listActiveSystems"
	buf, err := xml.EncodeClientRequest(method, Sessionkey)
	if err != nil {
		log.Fatalf("Encoding error: %s\n", err)
	}

	resp, err := request.MakeRequest(buf)
	if err != nil {
		log.Fatalf("GetActiveSystems API error: %s\n", err)
	}

	err = xml.DecodeClientResponse(resp.Body, l)
	if err != nil {
		log.Printf("Decode GetActiveSystems response body failed: %s\n", err)
	}
	log.Printf("Total %d active systems.\n", len(l.Result))

	return nil
}

func (l *ListActiveSystem) Print_Result() {
	for _, x := range l.Result {
		fmt.Println()
		fmt.Printf("ID: %d\nName: %s\nLast_boot: %s\nLast_checkin: %s\n", x.Id, x.Name, x.Last_checkin.Format(time.RFC3339),
			x.Last_boot.Format(time.RFC3339))
	}
}
