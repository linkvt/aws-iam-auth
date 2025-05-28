package elasticache

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/linkvt/aws-iam-auth/internal/auth"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Signer *auth.Signer
}

type GenerateTokenArgs struct {
	// TODO: replace CacheName with endpoint, extract cache name and reuse endpoint for testing?
	CacheName  string `required:"" short:"c" help:"Name of the ElastiCache."`
	Region     string `            short:"r" help:"AWS region of the ElastiCache instance."`
	Serverless bool   `            short:"s" help:"Create a token for a serverless ElastiCache."`
	Username   string `required:"" short:"u" help:"Name of the database user."`
}

type TestTokenArgs struct {
	Test         bool   `group:"test" short:"t" help:"Tests the token by sending a PING to the ElastiCache."`
	TestEndpoint string `group:"test"           help:"Endpoint used for testing the connection."`
	TestPort     int    `group:"test"           help:"Port used for testing the connection."                 default:"6379"`
	TestUseTls   bool   `group:"test"           help:"Use TLS for connecting to ElastiCache."`
}

// GenerateToken generates a password used for IAM authentication to the ElastiCache endpoint.
func (s *Service) GenerateToken(ctx context.Context, args GenerateTokenArgs) (string, error) {
	authParams := auth.TokenParams{
		Host:     args.CacheName,
		Service:  "elasticache",
		Username: args.Username,
	}
	if args.Serverless {
		authParams.ExtraParams["ResourceType"] = "Serverless"
	}

	return s.Signer.GenerateAuthToken(ctx, authParams)
}

// TestToken verifies a generated IAM auth token by pinging the ElastiCache endpoint.
func (s *Service) TestToken(
	ctx context.Context,
	token string,
	tokenArgs GenerateTokenArgs,
	testArgs TestTokenArgs,
) bool {
	logger := slog.With(
		slog.String("host", tokenArgs.CacheName),
		slog.String("username", tokenArgs.Username),
	)
	logger.Info("Testing ElastiCache IAM authentication")

	addr := fmt.Sprintf("%s:%d", tokenArgs.CacheName, testArgs.TestPort)
	opt := &redis.Options{
		Addr:     addr,
		Username: tokenArgs.Username,
		Password: token,
	}
	if testArgs.TestUseTls {
		opt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS13,
		}
	}

	// TODO: make code more testable
	client := redis.NewClient(opt)
	defer client.Close()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Info(fmt.Sprintf("Failed to authenticate with token: %v", err))
		return false
	}

	logger.Info(("Successfully authenticated with token"))
	return true
}
