package exporter

import (
	"encoding/xml"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"go_training/week3"
	"gopkg.in/yaml.v3"
)

//  To export the social media feeds to XML
func ExportJson(u week3.SocialMedia, filename string) error {
	feed := make(map[int]string)
	for index, val := range u.Feed() {
		feed[index+1] = val
	}
	bu, err := json.Marshal(feed)
	if err != nil{
		fmt.Println("An error occurred opening this file :", err.Error())
	}
	ioutil.WriteFile(filename, bu, 0644)
	return nil
}

// To export the social media feeds to XML
func ExportXml(u week3.SocialMedia, filename string) error {
	file, err := xml.MarshalIndent(u.Feed(), " ", "  ")
	if err != nil{
		fmt.Println("An error occurred opening this file :", err.Error())
	}
	ioutil.WriteFile(filename, file, 0644)
	return nil
}

// To export the social media feeds to XML
func ExportYaml(u week3.SocialMedia, filename string) error{
	out, err := yaml.Marshal(u.Feed())
	if err != nil {
		fmt.Println("An error occurred opening this file", err.Error())
	}
	ioutil.WriteFile(filename, out, 0644)
	return nil
}

