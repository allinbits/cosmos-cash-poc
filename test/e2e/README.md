### BDD e2e tests for cosmos cash

#### How to run

- `go test`

go test -v --godog.random --godog.tags=wip
go test -v --godog.format=pretty --godog.random -race -coverprofile=coverage.txt -covermode=atomic
