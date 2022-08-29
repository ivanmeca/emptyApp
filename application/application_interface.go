package application

import "context"

type IApplication interface {
	Init(ctx context.Context) error
	Run(ctx context.Context) error
}
