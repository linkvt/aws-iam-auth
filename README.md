# aws-iam-auth

`aws-iam-auth` generates IAM authentication tokens for various AWS services with additional features to validate the connection.

## Installation

### Go

```console
go install github.com/linkvt/aws-iam-auth@latest
```


## Example Usage

```console
❯ go run main.go --help                                                                         
Usage: aws-iam-auth <command>

A tool that generates tokens for IAM authentication to various AWS services.

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  elasticache --cache-name=STRING --username=STRING [flags]
    Generate a token for ElastiCache

  rds --endpoint=STRING --username=STRING [flags]
    Generate a token for RDS

Run "aws-iam-auth <command> --help" for more information on a command.
```

## Ideas

- Make library for generating tokens public
- Implement providers that cache tokens and regenerate them when needed
- Brew Tap with goreleaser
