default:
    just --list

run: 
    air

swagger:
    swag fmt -g cmd/server/main.go
    swag init -g cmd/server/main.go
