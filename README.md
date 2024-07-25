# grpc-onion

## Start Devcontainer
- open command palette in vscode
- `Dev Containers: Reopen in Container`

## Docker
### Build
```bash
docker build . -t grpc-onion
```

### Run
```bash
docker run --rm grpc-onion:latest
```

## Go
### Build
```bash
CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app cmd/main.go
```

### Run
```bash
go run cmd/main.go
```
```bash
./app
```


## grpcurl
### list
```bash
grpcurl -plaintext localhost:3000 list
```

### Search
```bash
grpcurl -plaintext -d '{"id": "1"}' localhost:3000 grpc.EntryService.Search
```

### Create
```bash
grpcurl -plaintext -d '{"user": "TES00001"}' localhost:3000 grpc.EntryService.Create
```

## Test data
```bash
sqlite3 onion.db < entry.sql
```

## proto
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/entry.proto
```
