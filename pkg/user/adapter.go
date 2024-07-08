package user

import (
	"github.com/IvanDobriy/Module1/pkg/user"
)

type Adapter interface {
	NewModule1() user.Module1
}

type MyAdapter struct {
}

func (ma MyAdapter) NewModule1() user.Module1 {
	return user.Module1Lib{}
}
