# README.md

# Book Scraper

Scraper desenvolvido em Go para coletar dados estruturados do site:

[https://books.toscrape.com](https://books.toscrape.com)

O projeto realiza scraping paginado, exporta os dados em JSON e CSV, e executa um pipeline GitLab CI/CD com:

* testes automatizados
* lint
* build Docker
* push para Container Registry
* deploy simulado

---

# Arquitetura

```txt
book-scraper/
├── cmd/
│   └── scraper/
│       └── main.go
├── internal/
│   ├── exporter/
│   ├── logger/
│   ├── models/
│   └── scraper/
├── output/
├── test/
├── Dockerfile
├── .gitlab-ci.yml
├── .golangci.yml
├── Makefile
└── README.md
```

---

# Estrutura da Fonte de Dados

O scraper utiliza como origem os elementos HTML abaixo:

```html
<article class="product_pod">
  <div class="image_container">
    <a href="...">
      <img src="..." class="thumbnail" alt="...">
    </a>
  </div>

  <p class="star-rating Three"></p>

  <h3>
    <a href="..." title="...">
      Título do Livro
    </a>
  </h3>

  <div class="product_price">
    <p class="price_color">£12.99</p>

    <p class="instock availability">
      In stock
    </p>
  </div>
</article>
```

Os dados são convertidos para uma estrutura tipada em Go:

```json
{
  "title": "A Light in the Attic",
  "price": "£51.77",
  "availability": "In stock",
  "rating": "Three",
  "product_url": "https://books.toscrape.com/catalogue/a-light-in-the-attic_1000/index.html",
  "image_url": "https://books.toscrape.com/media/cache/..."
}
```

---

# Executando Localmente

## Pré-requisitos

* Go 1.23+
* Docker Desktop
* golangci-lint

---

## Rodando sem Docker

Instale as dependências:

```bash
go mod tidy
```

Execute o scraper:

```bash
go run ./cmd/scraper
```

Os arquivos serão gerados em:

```txt
output/books.json
output/books.csv
```

---

## Rodando com Docker

Build da imagem:

```bash
docker build -t book-scraper .
```

Executar container:

```bash
docker run --rm book-scraper
```

---

# Estrutura dos Dados

## JSON

```json
{
  "title": "string",
  "price": "string",
  "availability": "string",
  "rating": "string",
  "product_url": "string",
  "image_url": "string"
}
```

---

## CSV

```csv
title,price,availability,rating,product_url,image_url
```

---

# Pipeline GitLab CI/CD

O pipeline possui 4 stages.

---

## 1. Test

Executa:

```bash
go test ./...
```

Objetivo:

* validar o funcionamento do scraper
* impedir merge de código quebrado

---

## 2. Lint

Executa:

```bash
golangci-lint run
```

Objetivo:

* garantir qualidade de código
* detectar erros comuns
* aplicar boas práticas

---

## 3. Build

Executa:

* build da imagem Docker
* push para GitLab Container Registry

Variáveis utilizadas:

* `CI_REGISTRY`
* `CI_REGISTRY_USER`
* `CI_REGISTRY_PASSWORD`
* `CI_REGISTRY_IMAGE`

---

## 4. Deploy

Executado apenas na branch `main`.

Atualmente o deploy é simulado utilizando `echo` para demonstrar o fluxo de publicação em AWS ECS.

---

# Decisões Técnicas

## Go

Go foi escolhido por:

* simplicidade
* binário único
* ótima integração com Docker
* baixo consumo de memória
* facilidade para CI/CD

---

## GoQuery

GoQuery foi utilizado para parsing HTML por oferecer:

* seletores CSS simples
* boa legibilidade
* excelente integração com HTML estático

---

## Estrutura Modular

O projeto foi separado em:

* scraper
* exporter
* logger
* models

Essa divisão facilita:

* manutenção
* testes
* reutilização
* desacoplamento

---

## Multi-stage Docker Build

O Dockerfile utiliza multi-stage build para:

* reduzir tamanho da imagem
* separar build/runtime
* aumentar segurança

---

# Melhorias Futuras

Com mais tempo seria possível implementar:

* concorrência com goroutines
* retry exponencial
* rotação de User-Agent
* rate limiting
* proxies
* persistência em PostgreSQL
* observabilidade com Prometheus
* tracing
* healthchecks
* scraping resiliente
* dashboards

---

# Uso de IA Durante o Desafio

A IA foi utilizada como ferramenta de apoio para:

* revisão arquitetural
* melhoria do Dockerfile
* refinamento do pipeline CI/CD
* validação de boas práticas Go
* ajustes de lint
* documentação técnica

A validação final das decisões e correções foi realizada manualmente.

---

# Dockerfile

```dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o scraper ./cmd/scraper

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/scraper .

RUN mkdir -p output

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./scraper"]
```

O container utiliza:

* multi-stage build
* usuário não-root
* imagem Alpine minimalista
* separação entre build e runtime