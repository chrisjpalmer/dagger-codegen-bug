// Provides utilities for interacting with git

package main

import (
	"context"
	"dagger/git-repo/internal/dagger"

	"github.com/go-git/go-git/v5"
)

const ghHost = "github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl"

type RegistryAuth interface {
	DaggerObject
	From(address string) *dagger.Container
}

type GitRepo struct {
	// +private
	DefaultBranch string
	// +private
	SSH *dagger.Socket
	// +private
	Src *dagger.Directory
}

func New(
	// +default="main"
	defaultBranch string,
	// the source containing the git repository
	src *dagger.Directory,
	// ssh socket used for cloning the source
	ssh *dagger.Socket,

) *GitRepo {
	return &GitRepo{
		DefaultBranch: defaultBranch,
		SSH:           ssh,
		Src:           src,
	}
}

// GitContainer returns a container with Git and SSH configured.
func (g *GitRepo) GitContainer(ctx context.Context) (*dagger.Container, error) {
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

func repository(ctx context.Context, src *dagger.Directory) (*git.Repository, error) {
	//
	//
	//
	//
	//
	//
}
