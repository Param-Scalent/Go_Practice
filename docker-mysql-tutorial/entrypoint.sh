wait-for "${DATABASE_PORT}:${DATABASE_PORT}" -- "$@"

# Watch for .go file changes
CompileDaemon --build="go build -o main.go" ---command=./main