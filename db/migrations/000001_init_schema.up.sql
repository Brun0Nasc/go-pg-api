CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE sexgen AS ENUM ('M', 'F');

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.data_atualizacao = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS diretores (
	"id" UUID NOT NULL DEFAULT(uuid_generate_v4()),
	"nome" VARCHAR NOT NULL,
	"sexo" sexgen NOT NULL,
	"data_criacao" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"data_atualizacao" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"data_remocao" TIMESTAMP(3),
    CONSTRAINT "diretores_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "diretores_id_key" ON "diretores"("id");

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON diretores
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();