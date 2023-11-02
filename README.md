# DOCKER DO BANCO DE DADOS
 Tem que instalar docker desktop
 > https://docs.docker.com/desktop/install/windows-install/
   para iniciar o docker 
 - docker-compose up
   para apra o docker
 - docker-compose dow 

 > caso precise recriar o banco de dados:
    rode o esse comando abaixo ele vai excluir e crair um novo banco zerado
 -   Remove-Item -Recurse -Force ./tmp_data/postgres_data
 
# API notification

Sistema  API com Golang

> Requisitos do projeto:

- Go Lang >= 1.18

As demais dependências estão no arquivo go.mod e package.json

- https://go.dev/dl/

> Build do Back-End Go:
```bash
# Baixando as dependências
$ go mod tidy

# Compilar servidor HTTP
$ go build -o main cmd/product/main.go

# Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o main64.exe cmd/product/main.go

# build modo production
$ go build -ldflags "-s -w" .
# Ou
$ go build -ldflags "-s -w" cmd/product/main.go
# Ou
$ go build -ldflags "-s -w" -o main cmd/product/main.go
```
## Opções de execução roda esse comando como está no pronto do windows terminal 
 $env:SRV_PORT="8080"; $env:SRV_MODE="developer"; $env:SRV_DB_HOST="localhost"; $env:SRV_DB_DRIVE="postgres"; $env:SRV_DB_HOST= "localhost"; $env:SRV_DB_PORT= "5432"; $env:SRV_DB_USER= "postgres"; $env:SRV_DB_PASS= "supersenha"; $env:SRV_DB_NAME= "notificacao_db";
 go run cmd/api/main.go

> Exemplo de Uso:
```bash
$ ./main.exe
# Ou
$ SRV_PORT=8080 SRV_MODE=developer ./main.exe
# Ou
$ SRV_PORT=9090 SRV_MODE=production ./main.exe
```

> Acesse:
- http://localhost:8080/api/v1/products


