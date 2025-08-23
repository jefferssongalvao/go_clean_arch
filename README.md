# Go Clean Architecture API

[![Build](https://img.shields.io/github/actions/workflow/status/SEU_USUARIO/SEU_REPO/ci.yml?branch=main)](https://github.com/SEU_USUARIO/SEU_REPO/actions)
[![Go](https://img.shields.io/badge/Go-1.24.2-blue)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-OK-blue)](https://www.docker.com/)

API em **Go** seguindo os princípios de Clean Architecture, com **PostgreSQL** como banco de dados. Inclui suporte a migrações e seed de dados.

---

## Tecnologias

* Go 1.24.2
* PostgreSQL 17
* Docker / Docker Compose
* Makefile para facilitar comandos

---

## Pré-requisitos

* Docker e Docker Compose instalados
* `make` disponível no sistema
* Git (opcional, para clonar o projeto)

---

## Configuração

```bash
make init
```

Este comando irá:

1. Criar os containers da API e do PostgreSQL
2. Aguardar o banco estar pronto
3. Rodar migrações (`migrate`)
4. Rodar seed inicial de dados (`seed`)

### Subir containers

```bash
make up
```

### Parar e remover containers

```bash
make down
```

### Ver logs da API

```bash
make logs
```

### Rodar migrações manualmente

```bash
make migrate
```

---

## Estrutura da API

* `cmd/app/main.go` → Entrypoint principal da API
* `cmd/migrate/main.go` → Entrypoint das migrações e seed
* `internal/` → Camadas internas (domain, repository, usecase, handler)
* `pkg/` → Pacotes públicos auxiliares

---

## Endpoints principais

| Método | Endpoint        | Descrição             |
| ------ | --------------- | --------------------- |
| GET    | `/students`     | Lista todos os alunos |
| POST   | `/students`     | Cria um novo aluno    |
| GET    | `/students/:id` | Detalha um aluno      |
| PUT    | `/students/:id` | Atualiza um aluno     |
| DELETE | `/students/:id` | Remove um aluno       |

> Substitua `students` pelo recurso que sua API gerencia.

---

## Observações

* Certifique-se de que a porta definida em `.env` (`API_PORT`) não está em uso.
* O banco será inicializado automaticamente com os dados do seed.
* Para produção, recomenda-se usar uma imagem Go baseada em Debian (`bullseye` ou `bookworm`) para reduzir vulnerabilidades.

---
