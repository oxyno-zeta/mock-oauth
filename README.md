## Mock OAuth server

Implements the following endpoints with minimal validations
- /auth endpoint
- /token endpoint
- /userinfo endpoint

## How to run

| #       | Description           | Command  |
| :------------- |:-------------| :-----|
| 1      | Run | `cd ${current.project.path} && go get -d && go run main.go` |

## Userinfo for a specific username
Set the env var `username=some_username`
