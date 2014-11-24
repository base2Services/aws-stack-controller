package main

import (
	"net/http"
	"flag"
	"log"
	b2aws "github.com/base2Services/go-b2aws"
	awsstackcontrol "github.com/base2Services/go-aws-stack-control"
)

var stack string
var environment string
var publicKey string
var secretKey string
var action string
var force bool

const profileName = "CLI"

func init() {
	flag.StringVar(&stack, "s", "", "Name of stack to shutdown or startup")
	flag.StringVar(&environment, "e", "", "Name of environment to shutdown or startup")
	flag.StringVar(&publicKey, "p", "", "aws iam publicKey")
	flag.StringVar(&secretKey, "S", "", "aws iam secretKey")
	flag.StringVar(&action, "a", "", "'startup' OR 'shutdown' without this, this does nothing")
	flag.BoolVar(&force, "f", false, "Force production actions on production tagged servers")
}

func main() {
	flag.Parse()
	switch (action) {
	case "shutdown":
		shutdownGroup()
	case "startup":
		startupGroup()
	default:
		log.Println("No action selected")
	}
}

func regionsAndInstances(client *http.Client) (regionMap map[string]string, instances []b2aws.Instance, err error) {
	log.Println("Getting regions")
	regionMap, err = b2aws.GetRegions(publicKey, secretKey, client)
	if err != nil { return }
	for region, regionUrl := range regionMap {
		log.Println("Getting instances for region " + region)
		instanceList, err := b2aws.GetInstances(publicKey, secretKey, regionUrl, client)
		if err != nil { return regionMap, instances, err }
		for _, instance := range instanceList {
			instance.ProfileName = profileName
			instance.Region = region
			instances = append(instances, instance)
		}
	}
	return
}

func shutdownGroup() {
	log.Printf("shutdown group %s %s\n", environment, stack)

	if environment == "Production" && !force {
		log.Fatalf("Won't shutdown a production environment.")
		return
	}

	client := http.DefaultClient
	regionMap, instances, err := regionsAndInstances(client)
	if err != nil {
		log.Panicln("Error: " + err.Error())
		return
	}

	log.Printf("task func shutdown group %s %s\n", environment, stack)
	callback := createCallback(client, regionMap, instances)
	awsstackcontrol.ShutdownEnvironment(client, regionMap, stack, environment, profileName, publicKey, secretKey, callback)
}

func startupGroup() {
	log.Printf("startup group %s %s\n", environment, stack)

	if environment == "Production" && !force {
		log.Fatalf("Won't shutdown a production environment.")
		return
	}

	client := http.DefaultClient
	regionMap, instances, err := regionsAndInstances(client)
	if err != nil {
		log.Panicln("Error: " + err.Error())
		return
	}

	log.Printf("task func startup group %s %s\n", environment, stack)
	callback := createCallback(client, regionMap, instances)
	awsstackcontrol.StartupEnvironment(client, regionMap, stack, environment, profileName, publicKey, secretKey, callback)
}

