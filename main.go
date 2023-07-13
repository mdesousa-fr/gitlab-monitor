package main

import (
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

	for _, policy := range cfg.Policies {
		for _, g := range policy.Groups {
			grp, err := gitlabClient.GetGroup(g)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(grp)
			//fmt.Println(len(grp.Projects))
			for _, p := range grp.Projects {
				if p.MergeMethod != policy.MergeMethod {
					log.Printf("%s merge method is not compliant", p.Url)
				} else {
					log.Printf("%s merge method is compliant", p.Url)
				}
			}
		}
	}
}
