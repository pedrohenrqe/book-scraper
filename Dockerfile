# =========================
# Stage 1 - Build da aplicação
# =========================

FROM golang:1.25-alpine AS builder

# Define diretório de trabalho
WORKDIR /app

# Copia arquivos de dependências primeiro
# para aproveitar cache do Docker
COPY go.mod go.sum ./

# Baixa dependências
RUN go mod download

# Copia restante do projeto
COPY . .

# Compila aplicação estática
RUN CGO_ENABLED=0 GOOS=linux go build -o scraper ./cmd/scraper

# =========================
# Stage 2 - Runtime
# =========================

FROM alpine:latest

# Instala certificados SSL
RUN apk --no-cache add ca-certificates

# Cria usuário não-root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Define diretório da aplicação
WORKDIR /app

# Copia binário compilado
COPY --from=builder /app/scraper .

# Cria diretório de saída
RUN mkdir -p output

# Ajusta permissões
RUN chown -R appuser:appgroup /app

# Troca para usuário não-root
USER appuser

# Porta exposta
# (não utilizada atualmente, mas preparada para futura API/healthcheck)
EXPOSE 8080

# Comando padrão
CMD ["./scraper"]