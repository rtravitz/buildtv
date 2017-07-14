package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

//Job is a Jenkins Job Response
type Job struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//Health is a Job Health Response
type Health struct {
	Description string `json:"description"`
	Score       int    `json:"score"`
}

func makeJenkinsRequest(u *User, reqURL string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.SetBasicAuth(u.name, u.token)
	res, err := client.Do(req)
	return res, err
}

func getPipelines(u *User, team string) ([]Job, error) {
	reqURL := fmt.Sprintf("%s/job/%s/api/json", baseURL, team)
	fmt.Println(reqURL)
	res, err := makeJenkinsRequest(u, reqURL)
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

func getHealth(u *User, team, project string) (Health, error) {
	reqURL := fmt.Sprintf("%s/job/%s/job/%s/api/json", baseURL, team, project)
	res, err := makeJenkinsRequest(u, reqURL)
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

func getTeams(u *User) ([]Job, error) {
	reqURL := fmt.Sprintf("%s/api/json", baseURL)
	res, err := makeJenkinsRequest(u, reqURL)
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
