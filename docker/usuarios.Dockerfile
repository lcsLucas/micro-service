# Imagem padrão
FROM golang 

# Diretório de trabalho
WORKDIR /app/src/usuarios

#aponta a váriavel GOPATH para o diretório app
ENV GOPATH=/app

# copia os arquivos de usuários do projeto para o workdir do container
COPY /usuarios /app/src/usuarios
COPY /cmd/usuarios /app/cmd/usuarios

#executa o main.go de usuários e baixa as dependências
RUN go build app/cmd/usuarios/main.go

# Comando para rodar o executável
ENTRYPOINT ["./main"]

# expo~e a porta 8080
EXPOSE 8080