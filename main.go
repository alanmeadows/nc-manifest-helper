package main

import (

	"fmt"
	"os"
	"os/user"
	"flag"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"

)

type ControlPlaneAddresses struct {
	Schema   string `json:"schema"`
	Metadata struct {
		Schema             string `json:"schema"`
		Name               string `json:"name"`
		LayeringDefinition struct {
			Abstract bool   `json:"abstract"`
			Layer    string `json:"layer"`
		} `json:"layeringDefinition"`
		StoragePolicy string `json:"storagePolicy"`
	} `json:"metadata"`
	Data struct {
		Genesis struct {
			Hostname string `json:"hostname"`
			IP       struct {
				Oam string `json:"oam"`
				Ksn string `json:"ksn"`
			} `json:"ip"`
		} `json:"genesis"`
		Masters []struct {
			Hostname string `json:"hostname"`
			IP       struct {
				Oam string `json:"oam"`
				Ksn string `json:"ksn"`
			} `json:"ip"`
		} `json:"masters"`
	} `json:"data"`
}

func main() {


	// get user homedir - this makes sense for alan and probably noone else
	usr, _ := user.Current()
	usrdir := usr.HomeDir
	defaultSiteDirectory := filepath.Join(usrdir, "Workbench", "aic-clcp-site-manifests")

	// commands we implement
	getIpCommand := flag.NewFlagSet("getip", flag.ExitOnError)
	// ....

	/* 

	// getIpCommand

	*/


	// common
	siteDirectory := getIpCommand.String("sites-dir", defaultSiteDirectory, "Path to aic-clcp-sites-manifests")
	siteName := getIpCommand.String("site", "", "The site name (e.g. mtn52a)")

	// specific to getIpCommand
	getIpCommandHostFlag := getIpCommand.String("host", "", "The hostname of the host or genesis")

	// default help

	if len(os.Args) == 1 {
		fmt.Println("usage: nc-manifest-helper <command> [<args>]")
		fmt.Println(" getip		Get the IP Address of a host in the site")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "getip":
		getIpCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid command.\n", os.Args[1])
		os.Exit(1)

	}

	// handle getIpCommand
	if getIpCommand.Parsed() {

		hostIp := getHostIp(*siteDirectory, *siteName, *getIpCommandHostFlag)
		fmt.Println(hostIp)

	}

}

func getHostIp(siteDirectory, siteName, hostName string) string {

	// DEBUG
	// fmt.Printf("siteDirectory -> %v, siteName -> %v, hostName -> %v\n", siteDirectory, siteName, hostName)

	if hostName == "genesis" {

		filename  := filepath.Join(siteDirectory, "site", siteName, "network", "control-plane-addresses.yaml")
		yamlFile, err := ioutil.ReadFile(filename)

	    if err != nil {
	        panic(err)
	    }

	    var controlPlaneConfig ControlPlaneAddresses

	    err = yaml.Unmarshal(yamlFile, &controlPlaneConfig)
	    if err != nil {
	    	panic(err)
	    }

	    return controlPlaneConfig.Data.Genesis.IP.Oam

	} else {

		// TODO
		return ""
	}
}