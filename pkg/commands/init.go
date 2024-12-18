package commands

import (
	"fmt"

	"github.com/jayakrishnanMurali/kit/pkg/repository"
)

func InitCmd(args []string) error {
	fmt.Println("Initializing repository...")
	_, err := repository.RepoCreate(args[0])

	return err
}
