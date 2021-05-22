windows:
	env GOOS=windows GOARCH=amd64 go build -o sigdump.exe cmd/main.go
linux:
	env GOOS=linux GOARCH=amd64 go build -o sigdump cmd/main.go