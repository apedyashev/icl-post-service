OUT_BINARY=postsApp

build:
	env GOOS=linux CGO_ENABLED=0 go build -o ${OUT_BINARY} ./
 