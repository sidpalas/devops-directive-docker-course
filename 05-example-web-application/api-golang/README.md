```bash
mkdir go-workspace
export GOPATH=$PWD/go-workspace
go mod download
DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres go run main.go
```