package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sumaxmlrpc/auth"
	"sumaxmlrpc/crypt"
	"sumaxmlrpc/request"
	"sumaxmlrpc/schedules"
	"sumaxmlrpc/system"
	"sumaxmlrpc/updategpgssl"
)

const (
	usage = `usage: %s -sumaconf <sumaconf.yaml> subcommand

Subcommands:
  listsystems
  listupgrades
  listjobs
  schedule_upgrades
    -schedule (provide number of hours to start the job from now)

  gpgsslupdate
		-file (the path to file that should be imported)
		-description (the name of the GPG or SSL definition)
		-type ("GPG" or "SSL")

e.g.:
  <program> -sumaconf sumaconfg.yaml schedule_upgrades -schedule 3

Options:
`
	listsystems_usage       = "usage: %s listsystems\n"
	listupgrades_usage      = "usage: %s listupgrades\n"
	schedule_upgrades_usage = "usage: %s schedule_upgrades -schedule <delay hours>\n"
)

type Sumaconf struct {
	Server   string `yaml:"server"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var (
	sumaconf                                                                           crypt.Sumaconf
	jobstart                                                                           *int
	gpgssl_file, gpgssl_description, gpgssl_type                                       *string
	Listsystems_yes, Listpkgs_yes, Scheduleupgpkgs_yes, ListJobs_yes, Gpgsslupdate_yes bool
	cmdlistsystems, cmdlistpkgs, cmdscheduleupgpkgs, cmdlistjobs, gpgsslupdate         *flag.FlagSet
	args                                                                               []string
	step                                                                               string
)

func init() {
	Listsystems_yes = false
	Listpkgs_yes = false
	Scheduleupgpkgs_yes = false
	ListJobs_yes = false
	Gpgsslupdate_yes = false

	var conf_file = flag.String("sumaconf", "", "provide the suma conf file with login data.")
	var create_config = flag.String("create-sumaconf", "", "Create a config file with login data.")
	flag.Usage = func() { // [1]
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	cmdlistsystems = flag.NewFlagSet("listsystems", flag.ExitOnError)
	cmdlistpkgs = flag.NewFlagSet("listupgrades", flag.ExitOnError)
	cmdscheduleupgpkgs = flag.NewFlagSet("schedule_upgrades", flag.ExitOnError)
	cmdlistjobs = flag.NewFlagSet("listjobs", flag.ExitOnError)
	jobstart = cmdscheduleupgpkgs.Int("schedule", -1, "provide number of hours to start the job from now.")
	gpgsslupdate = flag.NewFlagSet("gpgsslupdate", flag.ExitOnError)
	gpgssl_file = gpgsslupdate.String("file", "", "Enter ssl or gpg file path.")
	gpgssl_description = gpgsslupdate.String("description", "", "Enter the name of GPG or SSL definition.")
	gpgssl_type = gpgsslupdate.String("type", "", "Enter the type of GPG or SSL definition.")
	args = flag.Args()

	if len(*conf_file) == 0 && len(*create_config) == 0 {
		log.Fatal("sumaconf not provided or create a new sumaconf. Exit")
		step = "exit"
	} else if len(*conf_file) == 0 && len(*create_config) > 0 {
		crypt.PromptUser(create_config)
		step = "create_config"
	} else {
		sumaconf = sumaconf.Decrypt_Sumaconf(conf_file)
		step = "start_run"
	}

	/* if len(args) == 0 {
		log.Fatal("Please specify a subcommand.")
	} */
	//cmd, args := args[0], args[1:]
	cmdscheduleupgpkgs.Parse(args)

}
func main() {
	SessionKey := new(auth.SumaSessionKey)
	var err error
	if step == "start_run" {
		MysumaLogin := auth.Sumalogin{Login: sumaconf.Userid, Passwd: sumaconf.Password}
		request.Sumahost = &sumaconf.Server
		*SessionKey, err = auth.Login("auth.login", MysumaLogin)
		if err != nil {
			log.Fatal(err)
		}
	}

	switch args[0] {

	case "listsystems":
		cmdlistsystems.Parse(args[1:])
		Listsystems_yes = true
		fmt.Println("subcommand 'listsystems'")
	case "listupgrades":
		cmdlistpkgs.Parse(args[1:])
		Listsystems_yes = true
		Listpkgs_yes = true
	case "schedule_upgrades":
		cmdscheduleupgpkgs.Parse(args[1:])
		Listsystems_yes = true
		Listpkgs_yes = true
		Scheduleupgpkgs_yes = true
	case "listjobs":
		cmdlistjobs.Parse(args[1:])
		ListJobs_yes = true
	case "gpgsslupdate":
		gpgsslupdate.Parse(args[1:])
		log.Printf("get file path %s\n", *gpgssl_file)
		log.Printf("get description name %s\n", *gpgssl_description)
		Gpgsslupdate_yes = true
	default:
		fmt.Println("expected a subcommands")
		os.Exit(1)
	}

	if Listsystems_yes {
		listactivesystems := new(system.ListActiveSystem)
		_ = listactivesystems.GetActiveSystems(SessionKey)
		if Listpkgs_yes != true {
			listactivesystems.Print_Result()
		}

		if Listpkgs_yes {
			_ = listactivesystems.GetUpgPkgs(SessionKey)
			fmt.Printf("\n")
			if Scheduleupgpkgs_yes {
				_ = listactivesystems.ScheduleUpgradePkgs(SessionKey, jobstart)
			}
		}
	}

	if ListJobs_yes {
		listjobs := new(schedules.ListJobs)
		_ = listjobs.GetCompletedJobs(SessionKey)
		_ = listjobs.GetFailedJobs(SessionKey)
		_ = listjobs.GetPendingjobs(SessionKey)
	}

	if Gpgsslupdate_yes {
		updgpgssl := new(updategpgssl.Keys)
		err = updgpgssl.CreateKeys(*SessionKey, *gpgssl_description, *gpgssl_type, *gpgssl_file)
		err = updgpgssl.UpdateGPGSSL()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("\n")
	err = auth.Logout("auth.logout", *SessionKey)
	if err != nil {
		log.Fatal(err)
	}

}
