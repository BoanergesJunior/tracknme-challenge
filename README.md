# TrackNMe Challenge

## 🚀 Execução Local

Para executar o projeto localmente, você precisa ter o Docker instalado em sua máquina. Para parar todos os containers em execução, utilize o comando:

```bash
docker compose up -d --build
```

### Configuração do Ambiente

1. Na raiz do projeto, você encontrará um arquivo `env.example`
2. Copie o arquivo `env.example` para `.env`:
```bash
cp env.example .env
```
3. Ajuste as variáveis de ambiente no arquivo `.env`, especialmente a variável `GOMODCACHE` para apontar para o Go da sua máquina local

## ☁️ Ambiente em Produção

O projeto está disponível em produção através do seguinte endpoint:

```
https://tracknme-779596190711.us-central1.run.app
```

### Infraestrutura

O projeto utiliza os seguintes serviços em produção:

- **Cloud Run**: Hospedagem da aplicação
- **Cloud SQL (PostgreSQL)**: Banco de dados relacional para persistência de dados
- **Upstash**: Serviço de Redis para cache

### CI/CD

O projeto utiliza GitHub Workflows para integração e entrega contínua. Cada push para a branch `master` aciona automaticamente o processo de deploy, atualizando a aplicação no Cloud Run.

## 🛠️ Testando a API

Para testar os endpoints da API, você pode utilizar o arquivo de coleção do Insomnia disponível na raiz do projeto:

- Arquivo: `Insomnia_2025-05-16-tracknme`

Este arquivo contém todas as configurações necessárias para testar os endpoints da API localmente ou em produção.
