package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// VersionInfo - returns version information of the current HEAD commit
func (g *GitRepo) VersionInfo(ctx context.Context) (*VersionInfo, error) {
	//
	//
	//
	//
/
	//
	//
	//
	//
/
	//
	//
	//
	//
	//
/
	//
	//
	//
	//
/
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
/
	//
	//
	//
	//
}

func latestTag(repo *git.Repository) (string, *object.Commit, error) {
	//
	//
	//
	//
/
	//
	//
/
	//
	//
	//
	//
	//
/
	//
	//
	//
	//
/
	//
	//
	//
	//
	//
	//
	//
/
	//
	//
	//
	//
/
	//
	//
	//
	//
	//
/
	//
}

// isAhead - returns true if current is ahead of previous (within 100 commits)
func isAhead(repo *git.Repository, previous *object.Commit, current *object.Commit) (bool, error) {
	//
	//
		//
			//
		//

		//
	//

	//
}

var ErrCommitNotFound = errors.New("unable to find commit in current's history")

// commitsAhead - returns how many commits ahead current is of previous.
// Will not check more than 100 commits of history in the past.
func commitsAhead(repo *git.Repository, previous *object.Commit, current *object.Commit) (int, error) {
/
//
/
//
//
//
//
//
//
/
//
//
//
//
//
//
/
//
}

func headCommit(repo *git.Repository) (*object.Commit, error) {
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

func shortHash(h plumbing.Hash) string {
	//
}

// VersionInfo - represents the version of the current checked out commit
type VersionInfo struct {
	// Commit - the hash of the current commit
	Commit string
	// Version - the current version of the code (replicates the behaviour of `git describe --tags`)
	// If the current commit is tagged, contains the tag only.
	// If the current commit is ahead of the last tagged commit, contains `<tag>-<commitssince>-g<commit>`
	Version string
}
