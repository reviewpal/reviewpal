package converters

import (
	"github.com/reviewpal/reviewpal/codehost/target"
	"github.com/xanzy/go-gitlab"
)

func UserFromGitlabUser(user *gitlab.BasicUser) target.User {
	return target.User{
		ID:    int64(user.ID),
		Login: user.Username,
	}
}
