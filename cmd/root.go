package cmd

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/linkvt/aws-iam-auth/internal/elasticache"
	"github.com/linkvt/aws-iam-auth/internal/rds"
)

var cli struct {
	ElastiCache elasticache.Cmd `cmd:"" name:"elasticache" help:"Generate a token for ElastiCache"`
	Rds         rds.Cmd         `cmd:"" help:"Generate a token for RDS"`
}

func Execute() {
	kctx := kong.Parse(
		&cli,
		kong.Name("aws-iam-auth"),
		kong.Description(
			"A tool that generates tokens for IAM authentication to various AWS services.",
		),
		kong.UsageOnError(),
		kong.Groups{
			"test": "Connection Test:",
		},
	)

	ctx := context.Background()
	kctx.BindTo(ctx, (*context.Context)(nil))

	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}
