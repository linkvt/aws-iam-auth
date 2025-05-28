# aws-iam-auth

`aws-iam-auth` generates IAM authentication tokens for various AWS services with additional features to validate the connection.

## Installation

### Go

```console
go install github.com/linkvt/aws-iam-auth@latest
```

## Example Usage

### Generate a token for ElastiCache and test it

The generated token is printed to Stdout while logging messages are printed to Stderr, so that you can e.g. pipe the generate the token to other commands without logs interfering.

```console
❯ aws-iam-auth elasticache --username user --cache-name elasticache-name --test --test-endpoint test.localhost 
elasticache-name/?Action=&User=user&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIAZ3XJZLZ5FKCSTITK%2F20250528%2Feu-west-1%2Felasticache%2Faws4_request&X-Amz-Date=20250528T142015Z&X-Amz-Expires=900&X-Amz-
Security-Token=IQoJb3JpZ2luX2VjEK%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCWV1LXdlc3QtMSJIMEYCIQDv7x3xXLa26CN2mllwOQVUYIhUE6Gi8Z3E6RGhnOo60wIhAOM%2FBjgKHDihtqxksaCuG3tWkfZusyBU1faOclWEnGmJKvgCCHcQARoMNjc4MDIxMzg5OTQ2
...
2OkyZGUFnTzkb6lGnFw0qZjFDvn%2BtAbiXWwZ7FuF2hoIyaZs6AV4rMDH0yPVmeZ4kyA2wQI5mdYLla9wnNQz7TLL5Ba6uuDgXnOYGjcmlUOwEApmw%3D%3D&X-Amz-SignedHeaders=host&X-Amz-Signature=7119a39a2565d3213249eaa2c650baf9ffaa84cc88863b8b
5430c141f28e708e
2025/05/28 16:20:15 INFO Testing ElastiCache IAM authentication host=elasticache-name username=user
2025/05/28 16:20:20 INFO Failed to authenticate with token: dial tcp: lookup elasticache-name: no such host host=elasticache-name username=user
```

### Help

```console
❯ aws-iam-auth --help
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
