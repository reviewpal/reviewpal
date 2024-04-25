package converters

import (
	"github.com/google/go-github/v52/github"
	"github.com/reviewpal/reviewpal/codehost/target"
)

func UserFromGithubUser(user *github.User) target.User {
	return target.User{
		ID:    user.GetID(),
		Login: user.GetLogin(),
	}
}
