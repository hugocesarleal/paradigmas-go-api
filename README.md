# API Go - Benchmark de Requisições de Ações (Sequencial vs Paralelo)

Este projeto é uma API REST desenvolvida em **Go** que consome dados do mercado de ações através da [BRAPI](https://brapi.dev/). O objetivo principal é demonstrar e comparar o desempenho ao realizar múltiplas requisições HTTP de forma **sequencial** versus **paralela** (utilizando *Goroutines*).

Além das respostas em JSON, a API inclui um *endpoint* que gera um gráfico visual comparando os tempos de execução.

## 🚀 Tecnologias Utilizadas

* **[Go 1.22.2](https://go.dev/)** - Linguagem de programação
* **[Gin](https://gin-gonic.com/)** - *Framework* web para roteamento HTTP
* **Goroutines & `sync.WaitGroup`** - Processamento concorrente/paralelo
* **[go-echarts](https://github.com/go-echarts/go-echarts)** - Geração de gráficos em HTML
* **[godotenv](https://github.com/joho/godotenv)** - Carregamento de variáveis de ambiente

## ⚙️ Funcionalidades

* Busca de cotações de múltiplas ações
* Processamento **Sequencial** - Uma requisição por vez
* Processamento **Paralelo** - Requisições simultâneas com *Goroutines*
* **Benchmark** - Comparação de tempos de execução
* **Visualização** - Gráfico de barras em HTML

## 📋 Pré-requisitos

* [Go](https://go.dev/dl/) 1.22+
* Chave de API da [BRAPI](https://brapi.dev/)

## 🛠️ Configuração e Instalação

1. Clone o repositório:
```bash
cd paradigmas-go-main
```

2. Crie um arquivo `.env` na raiz do projeto:
```
BRAPI_URL=https://brapi.dev/api/quote/
BRAPI_API_KEY=SUA_CHAVE_AQUI
STOCKS=PETR4,VALE3,ITUB4,BBDC4,BBAS3,MGLU3,WEGE3,ABEV3,SUZB3,GGBR4,RAIL3,LREN3,B3SA3,LAME4,BRFS3,EMBR3,BRKM5,JBSS3,CIEL3,RADL3,USIM5
PORT=8080
```

3. Instale as dependências:
```bash
go mod tidy
```

4. Execute o servidor:
```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

## 📡 Endpoints da API

### GET /api/sequential
Busca sequencial de ações (uma por vez).

**Retorno:**
```json
{
    "mode": "sequential",
    "duration": 3.456,
    "data": [...]
}
```

### GET /api/parallel
Busca paralela de ações (todas simultâneas).

**Retorno:**
```json
{
    "mode": "parallel",
    "duration": 0.412,
    "data": [...]
}
```

### GET /api/benchmark
Executa ambas as buscas e retorna apenas os tempos.

**Retorno:**
```json
{
    "sequential": 3.456,
    "parallel": 0.412
}
```

### GET /api/graph
Gera um gráfico HTML interativo comparando os tempos.

## 📁 Estrutura do Projeto

```
.
├── main.go              # Ponto de entrada e configuração
├── handlers/            # Controladores da API
│   ├── benchmark.go
│   └── graph.go
├── services/            # Lógica de negócio
│   └── brapi.go
├── models/              # Estruturas de dados
│   └── stock.go
└── .env                 # Variáveis de ambiente
```