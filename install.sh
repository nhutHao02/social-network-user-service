go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
export PATH="$PATH:$(go env GOPATH)/bin"
go install github.com/google/wire/cmd/wire@latest