package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getInstanceType() string {
	// Retrieve the TOKEN
	req, err := http.NewRequest("PUT", "http://169.254.169.254/latest/api/token", nil)
	if err != nil {
		return ""
	}
	req.Header.Set("X-aws-ec2-metadata-token-ttl-seconds", "21600")
	tokenResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer tokenResp.Body.Close()

	tokenContent, err := ioutil.ReadAll(tokenResp.Body)
	if err != nil {
		return ""
	}
	token := string(tokenContent)

	// Make request to get instance type with TOKEN
	httpClient := &http.Client{
		Timeout: time.Second,
	}

	link := "http://169.254.169.254/latest/meta-data/instance-type"
	req, err = http.NewRequest("GET", link, nil)
	if err != nil {
		return ""
	}

	req.Header.Set("X-aws-ec2-metadata-token", token)
	response, err := httpClient.Do(req)
	if err != nil {
		return ""
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	return string(content)
}

func main() {
	instanceType := getInstanceType()
	if instanceType != "" {
		fmt.Println("Instance type:", instanceType)
	} else {
		fmt.Println("Failed to retrieve instance type.")
	}
}

