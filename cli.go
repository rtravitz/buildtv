package main

import (
	"bufio"
	"fmt"
	"os"

	"log"

	"github.com/fatih/color"
)

type User struct {
	name, token string
	pipelines   []Pipeline
	teams       []string
}

type Pipeline struct {
	team    string
	project string
	Health
}

func runCli() *User {
	u := getUser()
	teams, err := getTeams(u)
	if err != nil {
		log.Fatalln("Could not get teams: ", err)
	}
	addTeamsToUser(u, teams)
	err = choosePipeline(u)
	if err != nil {
		log.Fatalln("Could not add pipeline: ", err)
	}
	return u
}

func getUser() *User {
	c := color.New(color.FgCyan)
	c.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	c.Print("Api token: ")
	token, _ := reader.ReadString('\n')
	return &User{name: name, token: token}
}

func choosePipeline(u *User) error {
	c := color.New(color.FgCyan)
	c.Println("Select a team:")

	for idx, team := range u.teams {
		fmt.Printf("%d: %s\n", idx+1, team)
	}

	c.Print("> ")

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return err
	}

	team := u.teams[i-1]

	pipelines, err := getPipelines(u, team)
	if err != nil {
		return err
	}

	c.Println("Choose a pipeline: ")
	for idx, pipe := range pipelines {
		fmt.Printf("%d: %s\n", idx+1, pipe.Name)
	}
	c.Print("> ")

	_, err = fmt.Scanf("%d", &i)
	if err != nil {
		return err
	}

	pipelineName := pipelines[i-1].Name

	health, err := getHealth(u, team, pipelineName)
	if err != nil {
		return err
	}

	pipe := Pipeline{team, pipelines[i].Name, health}
	u.pipelines = append(u.pipelines, pipe)
	fmt.Println(*u)

	return nil
}
