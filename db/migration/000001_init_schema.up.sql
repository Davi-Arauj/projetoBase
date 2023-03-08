-- -- public.t_produtos definition

-- -- Drop table

-- -- DROP TABLE public.t_produtos;

CREATE TABLE IF NOT EXISTS "t_produto" (
	"id" bigserial NOT NULL PRIMARY KEY,
	"codigo_barras" bigint NOT NULL,
	"nome" varchar NOT NULL,
	"endereco_foto" varchar NOT NULL,
	"valor_pago" float NOT NULL,
	"valor_venda" float NOT NULL,
	"quantidade" bigint NOT NULL,
	"data_criacao" timestamptz NOT NULL DEFAULT now(),
	"data_atualizacao" timestamptz NULL
);

CREATE TABLE IF NOT EXISTS "t_cliente" (
	"id" bigserial PRIMARY KEY,
	"nome" varchar NOT NULL,
	"email" varchar NOT NULL,
	"cpf" varchar NOT NULL,
	"fone" bigint NOT NULL,
	"foto" varchar NOT NULL,
	"sexo" varchar NOT NULL,
	"data_nascimento" timestamp NOT NULL,
	"data_criacao" timestamptz NOT NULL DEFAULT (now()),
	"data_atualizacao" timestamptz NULL
);

CREATE TABLE IF NOT EXISTS "t_usuario" (
	"id" bigserial PRIMARY KEY,
	"nome" varchar NOT NULL,
	"email" varchar NOT NULL,
    "senha" varchar NOT NULL,
    "hash" varchar NOT NULL,
	"cpf" varchar NOT NULL,
	"fone" bigint NOT NULL,
	"foto" varchar NOT NULL,
	"sexo" varchar NOT NULL,
	"data_nascimento" timestamp NOT NULL,
	"data_criacao" timestamptz NOT NULL DEFAULT (now()),
	"data_atualizacao" timestamptz NULL
);



-- Table Triggers
-- create trigger updated_produto before
-- update
--     on
--     public.t_produto for each row execute function updated_datetime();
    
--    create trigger updated_cliente before
-- update
--     on
--     public.t_cliente for each row execute function updated_datetime();

