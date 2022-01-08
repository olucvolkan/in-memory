go build -o case .
go vet
go fmt

export $(cat .env | xargs)
./case