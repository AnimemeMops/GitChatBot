package actions

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetAllRepos(user *User, token string) (string, error) {
	str := ""
	repos, _, err := user.Client.Repositories.List(user.Context, "", nil)
	if err != nil {
		return str, err
	}

	for _, e := range repos {
		str += github.Stringify(e.Name) + "\n"
	}
	return str, err
}
func SetRepos(user *User, rep string) string {
	msg := "Doesn't exist"
	repos, _, err := user.Client.Repositories.List(user.Context, "", nil)
	if err != nil {
		msg = fmt.Sprintf("Problem in getting repository information %v\n", err)
		return msg
	}
	for _, e := range repos {
		if *e.Name == rep {
			user.Repos = *e
			msg = "Success"
		}
	}

	return msg
}
func GetAllCommits(user *User) {
	commitInfo, _, err := user.Client.Repositories.ListCommits(user.Context, "AnimemeMops", *user.Repos.Name, nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
	}

	fmt.Printf("%+v\n", github.Stringify(commitInfo[0].Commit.Message)) // Last commit information
}
