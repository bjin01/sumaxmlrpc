package system

import (
	"fmt"
	"log"
	"sumaxmlrpc/auth"
	"sumaxmlrpc/request"
	"time"

	"github.com/divan/gorilla-xmlrpc/xml"
)

func (l *ListActiveSystem) ScheduleUpgradePkgs(Sessionkey *auth.SumaSessionKey, delay *int) error {
	method := "system.schedulePackageInstall"
	type InputParams struct {
		Sessionkey         string
		Sid                int
		PackageIds         []int
		EarliestOccurrence time.Time
	}

	for a, b := range l.Result {
		var pkglist []int
		for _, d := range b.PackageList.Result {
			//fmt.Printf("pkg: %s\n", d.Name)
			//pkgid, _ := strconv.Atoi(d.To_package_id)
			pkglist = append(pkglist, d.To_package_id)
		}

		startime := time.Now().Local().Add(time.Hour * time.Duration(*delay))
		//fmt.Printf("%d: job time schedule: %s\n", b.Id, startime.String())
		inputsparams := InputParams{
			Sessionkey.Sessionkey,
			b.Id,
			pkglist,
			startime,
		}
		buf, err := xml.EncodeClientRequest(method, &inputsparams)
		if err != nil {
			log.Fatalf("Encoding error: %s\n", err)
		}
		resp, err := request.MakeRequest(buf)
		if err != nil {
			log.Fatalf("GetUpgPkgs API error: %s\n", err)
		}

		/* bodyB, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Raw logout xml: %s\n", bodyB) */

		jobid := new(UpdateJob)
		err = xml.DecodeClientResponse(resp.Body, jobid)
		if err != nil {
			log.Printf("Decode ScheduleUpgradePkgs response body failed: %s\n", err)
		}
		fmt.Printf("Upgrade Job %d for %s created.\n", jobid.ActionId, b.Name)
		l.Result[a].JobIDs = append(l.Result[a].JobIDs, *jobid)
	}
	return nil
}
