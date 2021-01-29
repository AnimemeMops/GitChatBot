package actions

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type User struct {
	ID      string
	Name    string
	Token   string
	Client  github.Client
	Context context.Context
	Repos   github.Repository
}

var Users map[string]*User = make(map[string]*User)

func (user *User) SetToken(token string) {
	user.Token = token
}
func (user *User) SetName(name string) {
	user.Name = name
}
func (user *User) Initialization() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: user.Token},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	var client *github.Client = github.NewClient(tokenClient)
	user.Client = *client
	user.Context = context
	Users[user.ID] = user

}
