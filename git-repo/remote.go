package main

import (
	"context"
	"regexp"
)

// Remote - returns the git remote
func (g *GitRepo) Remote(ctx context.Context) (*Remote, error) { //
}

// Remote - the git remote of the repo
type Remote struct {
	// Owner - the organisation/owner of the repository
	Owner string
	// Repo - the name of the repository
	Repo string
}

var sshMatch = regexp.MustCompile(`git@github\.com:(\S*)\/(\S*)`)
var httpsMatch = regexp.MustCompile(`https:\/\/github\.com\/(\S*)\/(\S*)`)

func parseRemote(remote string) (*Remote, error) {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
}
