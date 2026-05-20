# Book Scraper

Scraper desenvolvido em Go para coletar dados do site:

https://books.toscrape.com

O projeto executa scraping paginado, exporta dados estruturados em JSON e CSV, e roda em um pipeline completo GitLab CI/CD com testes, lint, build Docker e deploy simulado.

## Arquitetura

O projeto foi dividido em camadas:

- `cmd/`
  - ponto de entrada da aplicação

- `internal/scraper`
  - lógica de scraping

- `internal/models`
  - estruturas de dados

- `internal/exporter`
  - exportação JSON/CSV

- `internal/logger`
  - logs centralizados

## Executando localmente (sem Docker)

### Pré-requisitos

- Go 1.23+

### Instalação

```bash
go mod tidy


---

# 4. Como rodar com Docker

```md id="prkl1y"
## Executando com Docker

### Build da imagem

```bash
docker build -t book-scraper .


---

# 5. Schema dos dados

Isso é MUITO importante.

---

## JSON schema

```md id="8ij7gj"
## Estrutura dos dados

### JSON

```json
{
  "title": "string",
  "price": "string",
  "availability": "string",
  "rating": "string",
  "product_url": "string",
  "image_url": "string"
}

---

## CSV schema

```md id="ixxlo0"
### CSV

```csv
title,price,availability,rating,product_url,image_url

---

# 6. Pipeline explicado

Muito importante.

---

```md id="l2dr7w"
## Pipeline GitLab CI/CD

O pipeline possui 4 stages:

### 1. Test

Executa:

```bash
go test ./...


---

# 7. Decisões técnicas

ESSA PARTE É OURO EM ENTREVISTA.

---

```md id="kjuz4j"
## Decisões técnicas

### Go como linguagem principal

Go foi escolhido por:
- excelente concorrência
- baixo consumo
- binário único
- ótima integração com Docker
- simplicidade para pipelines CI/CD

---

### Uso do GoQuery

GoQuery oferece:
- parsing HTML eficiente
- API semelhante ao jQuery
- boa legibilidade

---

### Exportação em JSON e CSV

JSON:
- integração com APIs
- flexibilidade

CSV:
- análise em planilhas
- interoperabilidade

---

### Estrutura modular

A divisão em:
- scraper
- exporter
- models
- logger

foi feita para:
- facilitar testes
- reduzir acoplamento
- melhorar manutenção

---

### Multi-stage Docker build

Utilizado para:
- reduzir tamanho da imagem
- separar build/runtime
- aumentar segurança
## Melhorias futuras

Com mais tempo eu implementaria:

### Anti-bot handling

- rotação de User-Agent
- proxies
- retry exponencial
- rate limiting

---

### Observabilidade

- Prometheus
- métricas
- tracing
- healthchecks
- structured logging

---

### Persistência

- PostgreSQL
- SQLite
- versionamento de scraping

---

### Concorrência

Uso de goroutines para scraping paralelo.

---

### IA

Uso de LLMs para:
- classificação automática
- enriquecimento semântico
- detecção de páginas quebradas
- parsing resiliente

## Uso de IA durante o desafio

A IA foi utilizada como ferramenta de apoio técnico para:

- revisão arquitetural
- sugestões de estrutura de projeto
- melhoria do Dockerfile
- validação de boas práticas Go
- refinamento do pipeline CI/CD

---

### Exemplos de uso

- geração inicial da estrutura modular
- identificação de melhorias de segurança
- ajustes de lint
- revisão de tratamento de erros

---

### O que funcionou bem

- aceleração da documentação
- revisão de código
- melhoria de padrões de projeto

---

### O que precisou de validação manual

- compatibilidade de versões
- ajustes do golangci-lint
- detalhes específicos do ambiente Windows/Docker