package elasticache

import (
	"context"
	"fmt"
)

type Cmd struct {
	GenerateTokenArgs
	TestTokenArgs
}

func (c *Cmd) Validate() error {
	if c.Test {
		if c.TestEndpoint == "" {
			return fmt.Errorf("test endpoint has to be specified")
		}
		if c.TestPort < 0 || c.TestPort > 65_535 {
			return fmt.Errorf("test port is out of range: %d", c.TestPort)
		}
	}

	return nil
}

func (c *Cmd) Run(ctx context.Context) error {
	svc := &Service{}

	token, err := svc.GenerateToken(ctx, c.GenerateTokenArgs)
	if err != nil {
		return err
	}

	fmt.Println(token)

	if c.Test {
		svc.TestToken(ctx, token, c.GenerateTokenArgs, c.TestTokenArgs)
	}

	return nil
}
