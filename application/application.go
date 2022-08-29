package application

import (
	"context"
	"fmt"
)

type app struct {
	appCtx context.Context
}

func NewApp() IApplication {
	return &app{}
}

func (a *app) Init(ctx context.Context) error {
	a.appCtx = ctx
	return nil
}

func (a *app) Run(ctx context.Context) error {
	fmt.Printf("App is running")
	return nil
}
