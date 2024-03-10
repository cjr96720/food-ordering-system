FROM golang:1.21.3

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest