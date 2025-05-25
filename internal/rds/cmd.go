package rds

import (
	"context"
	"fmt"
)

type Cmd struct {
	GenerateTokenArgs
}

func (c *Cmd) Run(ctx context.Context) error {
	svc := &Service{}

	token, err := svc.GenerateToken(ctx, c.GenerateTokenArgs)
	if err != nil {
		return err
	}

	fmt.Println(token)

	// TODO: implement testing for postgres
	// TODO: implement testing for mysql

	return nil
}
