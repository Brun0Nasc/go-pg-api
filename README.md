# **PASSOS PARA CRIAÇÃO DA API**

* Essa é uma tradução e adaptação do guia encontrado [aqui](https://codevoweb.com/golang-crud-restful-api-with-sqlc-and-postgresql/).

---

## **PASSO 1**

Inicializar o projeto:

```bash
~$ go mod init github.com/Brun0Nasc/projeto
```

Criar um arquivo `docker-compose.yml` na raiz do projeto:

```yml
version: '3'
services:
    postgres:
        image: postgres:latest
        container_name: postgres
        ports:
            - '6500:5432'
        volumes:
            - postgresDB:/var/lib/postgresql/data
        env_file:
            - ./app.env
volumes:
    postgresDB:
```

Criar um arquivo `app.env` que tenha as credenciais da imagem do Postgres que será usada na montagem do container:

```dotenv
SERVER_PORT=8000
CLIENT_PORT=8080
NODE_ENV=development

POSTGRES_DRIVER=postgres
POSTGRES_SOURCE=postgresql://user:password@host:6500/database_name?sslmode=disable

POSTGRES_HOST=127.0.0.1
POSTGRES_PORT=6500
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=database_name

ORIGIN=http://localhost:3000
```

Depois disso adicione `app.env` a um arquivo `.gitignore`.

Com tudo isso configurado, é hora de rodar o docker-compose:

```bash
~$ docker compose up -d
```

Para parar de rodar o container:

```bash
~$ docker compose down
```

## **PASSO 2**

A ferramenta de Migrations utilizada é a biblioteca `golang-migrate` [que pode ser encontrada aqui](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

Alguns comandos essenciais:

* `create` - Para criar novos arquivos de migração
* `goto V` - Para mudar o schema de migrações para uma versão específica
* `up` - Para rodar as os arquivos de up migration sequencialmente
* `down` - Para rodar os arquivos de down migration na seqência inversa

Criar um diretório para armazenar as migrations: `db/migrations` e criar novos arquivos **up/down** dentro do diretório:

```bash
~$ migrate create -ext sql -dir db/migrations -seq init_schema
```

* `-ext` - indica a estenção dos arquivos de migração **up/down**
* `-dir` - indica o diretório onde serão armazenados esses arquivos de migração
* `-seq` - manda a biblioteca `golang-migrate` gerar uma um número de versões sequencial para os arquivos de migração

Agora, no arquivo de **up** gerado, encontrado em `db/migrations/000001_init_schema.up.sql`, eu adicionei a minha primeira tabela do meu banco de dados, com suas funções e parâmetros especificados.

No arquivo **down**, os comandos devem desfazer o que é feito no arquivo de criação das tabelas, isso permite voltar a uma versão anterior do banco, desfazendo as mudanças feitas no arquivo **up** mais recente.

Para executar o script de up migration, o seguinte comando:
```bash
~$ migrate -path db/migrations -database "postgresql://user:password@host:6500database_name?sslmode=disable" -verbose up
```
