package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/Pursuit92/github"
)

func main() {
	addr := flag.String("addr", ":8080", "Address to listen for hooks")
	site := flag.String("site", "", "Site location")
	flag.Parse()
	if *site == "" {
		log.Fatal("Must supply site directory")
	}
	ch, errch := github.ReceiveHooks(*addr)
	for {
		select {
		case <-ch:
			gitPull := exec.Command("/usr/bin/git", "pull")
			hugo := exec.Command("/usr/local/bin/hugo", "-s", *site)
			err := gitPull.Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Git ran successfully!")
			err = hugo.Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Hugo ran successfully!")
		case err := <-errch:
			log.Fatal(err)
		}
	}
}
