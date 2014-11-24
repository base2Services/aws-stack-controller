package main

import (
	"log"
	b2aws "github.com/base2Services/go-b2aws"
	"net/http"
)

func createCallback(client *http.Client, regionMap map[string]string, instances []b2aws.Instance) MyChannelCallback {
	return MyChannelCallback {
		RegionMap: regionMap,
		Client: client,
		Instances: instances,
	}
}

type MyChannelCallback struct {
	Instances	[]b2aws.Instance
	RegionMap	map[string]string
	Client	*http.Client
}

func (mccb MyChannelCallback) NoSuchEnvironment() {
	log.Printf("Can not startup or shutdown environment.\nEnsure all tags are in use.")}

func (mccb MyChannelCallback) MisingOrderTags() {
	log.Printf("No startup or shutdown tags filled, or all are below 1. Or couldn't find stack")}

func (mccb MyChannelCallback) TierShutdown() {
	log.Printf("Teir Shutdown")
}

func (mccb MyChannelCallback) StackShutdown() {
	log.Printf("Stack Shutdown")
}

func (mccb MyChannelCallback) TierStartedup() {
	log.Printf("Teir startup")
}

func (mccb MyChannelCallback) StackStartedup() {
	log.Printf("Stack startup")
}

func (mccb MyChannelCallback) TierTakingTooLong() {
	log.Printf("Teir taking to long skipping to next")
}

func (mccb MyChannelCallback) GetAllInstances() []b2aws.Instance {
	return mccb.Instances
}

func (mccb MyChannelCallback) Infof(format string, args ...interface{}){
	log.Printf(format,args...)
}

func (mccb MyChannelCallback) Warningf(format string, args ...interface{}){
	log.Printf(format,args...)
}

func (mccb MyChannelCallback) Errorf(format string, args ...interface{}){
	log.Fatalf(format,args...)
}

