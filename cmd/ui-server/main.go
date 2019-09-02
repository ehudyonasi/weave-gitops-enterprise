package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type paramSet struct {
	port            int
	privKeyFile     string
	gitURL          string
	gitPollInterval time.Duration
}

func main() {
	var params paramSet
	flag.IntVar(&params.port, "port", 80, "")
	flag.StringVar(&params.privKeyFile, "git-private-key-file", "private-key", "")
	flag.StringVar(&params.gitURL, "git-url", "git@github.com:fbarl/test-wkp.git", "")
	flag.DurationVar(&params.gitPollInterval, "git-poll-interval", 15*time.Second, "")
	flag.Parse()

	// Serve up the branch list!
	handleBranchesRequest, pollGitBranches := BranchesRequestHandler(params.gitURL, params.privKeyFile, params.gitPollInterval)
	if params.gitURL != "" {
		go pollGitBranches()
	}
	http.HandleFunc("/api/repo/branches", handleBranchesRequest)

	// Serve the UI
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Println("Listening on", params.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", params.port), nil))
}
