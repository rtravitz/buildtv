package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

func makeJenkinsRequest(reqURL string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.SetBasicAuth(username, token)
	res, err := client.Do(req)
	return res, err
}

func getJobs() ([]Job, error) {
	reqURL := fmt.Sprintf("%s/job/%s/api/json", baseURL, team)
	res, err := makeJenkinsRequest(reqURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	j := struct {
		Jobs []Job
	}{}

	json.NewDecoder(res.Body).Decode(&j)

	return j.Jobs, nil
}

func getHealth(pipeline string) (Health, error) {
	reqURL := fmt.Sprintf("%s/job/%s/job/%s/api/json", baseURL, team, pipeline)
	res, err := makeJenkinsRequest(reqURL)
	if err != nil {
		return Health{}, err
	}
	defer res.Body.Close()

	h := struct {
		HealthReport []Health `json:"healthReport"`
	}{}

	json.NewDecoder(res.Body).Decode(&h)

	return h.HealthReport[0], nil
}
