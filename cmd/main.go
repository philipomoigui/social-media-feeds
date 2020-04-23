package main

import (
	"go_training/week3/exporter"
	"go_training/week3/facebook"
	"go_training/week3/twitter"
	"go_training/week3/linkedin"
) 


func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)

	// to return Social media feeds in Json
	err := exporter.ExportJson(twtr, "twtrdata.json")
	err = exporter.ExportJson(fb, "fbdata.json")
	err =  exporter.ExportJson(lnkdin, "lnkdindata.json")

	// to return Social media feeds in XML
	err =  exporter.ExportXml(fb, "fbdata.xml")
	err =  exporter.ExportXml(twtr, "twtrdata.xml")
	err =  exporter.ExportXml(lnkdin, "lnkdindata.xml")

	// to return Social media feeds in YAML
	err =  exporter.ExportYaml(fb, "fbdata.yaml")
	err =  exporter.ExportYaml(twtr, "twtrdata.yaml")
	err =  exporter.ExportYaml(lnkdin, "lnkdindata.yaml")
	if err != nil {
		panic(err)
	}
}
