package main

import (
	"fmt"
	"log"

	"github.com/mdesousa-fr/gitlab-monitor/internal/config"
	"github.com/mdesousa-fr/gitlab-monitor/internal/gitlab"
)

func main() {
	cfg, err := config.ReadConfig("example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	gitlabClient := gitlab.NewClient("https://gitlab.com/api/v4", cfg.App.Token)
	err = gitlabClient.Auth()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range cfg.Policies {
		for _, g := range p.Groups {
			grp, err := gitlabClient.GetGroup(g)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(grp)
			fmt.Println(len(grp.Projects))
		}
	}
}
