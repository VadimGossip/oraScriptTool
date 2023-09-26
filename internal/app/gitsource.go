package app

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

type DBAdapter struct {
	//	gitWalkerRepo gitwalker.Repository
	//	cfg           *domain.Config
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{}
}

func (d *DBAdapter) Connect() error {

	_, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://git.etservice.net/telejet/tcs_db/tj-oracle",
	})
	if err != nil {
		return fmt.Errorf("error while opening git repo %s", err)
	}
	//	d.gitWalkerRepo = gitwalker.NewRepository(d.cfg.Path.RootDir, gitRepo)
	return nil
}
