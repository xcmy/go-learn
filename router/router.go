package router

import (
	"go-learn/component"
	"fmt"
)

func PrintName()  {
	component.PrintUrl()
}

type Router struct {
	Name string
}

func (router Router) GetName()  {
	fmt.Println(router.Name)
}

