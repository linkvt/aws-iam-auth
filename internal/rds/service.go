package rds

import (
	"context"

	"github.com/linkvt/aws-iam-auth/internal/auth"
)

type Service struct {
	Signer *auth.Signer
}

type GenerateTokenArgs struct {
	Endpoint string `required:"" short:"c" help:"Name of the RDS or Aurora endpoint."`
	Region   string `            short:"r" help:"AWS region of the RDS or Aurora instance."`
	Username string `required:"" short:"u" help:"Name of the database user."`
}

// GenerateToken generates a password used for IAM authentication to the Aurora/RDS instance or cluster.
func (s *Service) GenerateToken(ctx context.Context, args GenerateTokenArgs) (string, error) {
	authParams := auth.TokenParams{
		Host:     args.Endpoint,
		Service:  "rds-db",
		Username: args.Username,
	}

	// generate a token with our signer instead of the one in the SDK to keep errors consistent
	return s.Signer.GenerateAuthToken(ctx, authParams)
}
