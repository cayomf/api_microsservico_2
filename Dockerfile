# Imagem base com o Go instalado
FROM golang:latest

# Define o diretório de trabalho
WORKDIR /app

# Copia o código-fonte para o diretório de trabalho
COPY . .

# Compila a aplicação
RUN go build -o api_microsservico_2

# Define a porta em que a aplicação escuta
EXPOSE 8081

# Comando para executar a aplicação
CMD ["./api_microsservico_2"]
