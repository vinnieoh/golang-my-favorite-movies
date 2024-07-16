# Utiliza uma imagem base com Go instalada
FROM golang:1.18-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências necessárias
RUN go mod download

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila a aplicação Go
RUN go build -o main .

# Copia o script wait-for-database.sh para o diretório binário do sistema e dá permissão de execução
COPY ./scripts/wait-for-database.sh /usr/local/bin/wait-for-database.sh
RUN chmod +x /usr/local/bin/wait-for-database.sh

# Define build arguments para passar as variáveis de ambiente
ARG DB_URL
ARG JWT_SECRET
ARG ALGORITHM
ARG POSTGRES_HOST
ARG POSTGRES_PORT

# Define as variáveis de ambiente dentro do container
ENV DB_URL=${DB_URL}
ENV JWT_SECRET=${JWT_SECRET}
ENV ALGORITHM=${ALGORITHM}
ENV POSTGRES_HOST=${POSTGRES_HOST}
ENV POSTGRES_PORT=${POSTGRES_PORT}

# Define a porta que será exposta pelo container
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/usr/local/bin/wait-for-database.sh", "&&", "./main"]
