// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func InitializeShop(name,age string) (*Service,error) {
	wire.Build(NewS,NewBiz,MockUserRepoSet)
	return &Service{},nil
}