package system

import (
	"fmt"
	"log"
	"sumaxmlrpc/auth"
	"sumaxmlrpc/request"
	"time"

	"github.com/divan/gorilla-xmlrpc/xml"
)

func (l *ListActiveSystem) GetUpgPkgs(Sessionkey *auth.SumaSessionKey) error {
	method := "system.listLatestUpgradablePackages"
	type InputParams struct {
		Sessionkey string
		Sid        int
	}

	for a, k := range l.Result {
		inputsparams := InputParams{Sessionkey.Sessionkey, k.Id}

		buf, err := xml.EncodeClientRequest(method, &inputsparams)
		if err != nil {
			log.Fatalf("Encoding error: %s\n", err)
		}
		resp, err := request.MakeRequest(buf)
		if err != nil {
			log.Fatalf("GetUpgPkgs API error: %s\n", err)
		}

		host_pkglist := new(ListLatestUpgradablePackages)
		err = xml.DecodeClientResponse(resp.Body, host_pkglist)
		if err != nil {
			log.Printf("Decode GetUpgPkgs response body failed: %s\n", err)
		}
		l.Result[a].PackageList = *host_pkglist

	}

	/* bodyB, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw logout xml: %s\n", bodyB) */

	for _, x := range l.Result {
		fmt.Println()
		fmt.Printf("ID: %d\nName: %s\nLast_boot: %s\nLast_checkin: %s\nNo of Upgradeable Pkgs: %d\n", x.Id, x.Name, x.Last_checkin.Format(time.RFC3339),
			x.Last_boot.Format(time.RFC3339), len(x.PackageList.Result))
		/* for _, c := range x.PackageList.Result {
			fmt.Printf("%#v\n", c)
		}
		fmt.Println() */
	}

	return nil
}
