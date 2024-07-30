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
docker run --rm -p 3000:3000 -v $(pwd)/onion.db:/onion.db grpc-onion:latest
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
```bash
grpcurl -plaintext localhost:3000 list entry.EntryService 
```

### Search
```bash
grpcurl -plaintext -d '{"id": "1"}' localhost:3000 entry.EntryService.Search
```

### Create
```bash
grpcurl -plaintext -d '{"user": "TES00001"}' localhost:3000 entry.EntryService.Create
```

## Test data
```bash
sqlite3 onion.db < entry.sql
```

## proto
### buf cli
```bash
buf generate
```
