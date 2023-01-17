package system

import "time"

type ListActiveSystem struct {
	//Success bool `xml:"success"`
	Result []struct {
		Id           int
		Name         string
		Last_checkin time.Time
		Last_boot    time.Time
		PackageList  ListLatestUpgradablePackages
		JobIDs       []UpdateJob
	}
}

type ListLatestUpgradablePackages struct {
	Result []struct {
		Name          string
		Arch          string
		From_version  string
		From_release  string
		From_epoch    string
		To_version    string
		To_release    string
		To_epoch      string
		From_arch     string
		To_arch       string
		To_package_id int
	}
}

type SystemScheduleUpdate struct {
	Sid                int
	PackageIds         []int
	EarliestOccurrence string
}

type UpdateJob struct {
	ActionId int
}
