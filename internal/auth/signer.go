package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	signer "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

const emptyBodySha256 = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

type Signer struct{}

func NewSigner() *Signer {
	return &Signer{}
}

type TokenParams struct {
	Action      string
	ExtraParams map[string]string
	Host        string
	Region      string
	Service     string
	Username    string
}

func (s *Signer) GenerateAuthToken(
	ctx context.Context,
	params TokenParams,
) (string, error) {
	// TODO: pass aws credentials to this function
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve AWS credentials: %w", err)
	}

	query := url.Values{
		"Action":        {params.Action},
		"User":          {params.Username},
		"X-Amz-Expires": {"900"},
	}

	signURL := url.URL{
		Scheme:   "http",
		Host:     params.Host,
		Path:     "/",
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, signURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create signing request: %w", err)
	}

	region := params.Region
	if region == "" {
		region = cfg.Region
	}

	signer := signer.NewSigner()
	signedURI, _, err := signer.PresignHTTP(
		ctx,
		credentials,
		req,
		emptyBodySha256,
		params.Service,
		region,
		time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("failed to sign uri: %w", err)
	}

	u, err := url.Parse(signedURI)
	if err != nil {
		return "", fmt.Errorf("failed to parse signed URI: %w", err)
	}

	token := strings.TrimPrefix(u.String(), "http://")

	return token, nil
}
