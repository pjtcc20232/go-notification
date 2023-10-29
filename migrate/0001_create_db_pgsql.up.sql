CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE public.gurpo (
	id serial4 NOT NULL,
	grupo varchar(20) NOT NULL,
	permissoes varchar(2500) NOT NULL,
	
	CONSTRAINT grupo_pkey PRIMARY KEY (id)
);


CREATE TABLE usuarios (
    id serial4 NOT NULL,
    email varchar(255) NOT NULL, 
    senha varchar(255) NOT NULL,
    hashpass varchar(150) NOT NULL,
    grupo_id int NOT NULL,
    enable bool,
    created_at timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp without time zone NOT NULL DEFAULT NOW(),
    UNIQUE(email),
    CONSTRAINT fk_grupo_id FOREIGN KEY (grupo_id) REFERENCES public.gurpo(id)
);



CREATE TABLE public.cursos (
	id serial4 NOT NULL,
	nome varchar(200) NOT NULL,
   CONSTRAINT curso_pkey PRIMARY KEY (id)
);


CREATE TABLE public.turmas (
	id serial4 NOT NULL,
	horario varchar(20) NOT NULL,
	curso_id int NOT NULL,
	CONSTRAINT turma_pkey PRIMARY KEY (id),
	CONSTRAINT fk_curso_turma_id FOREIGN KEY (curso_id) REFERENCES public.cursos(id),
);


CREATE TABLE public.alunos (
	id serial4 NOT NULL,
    nome varchar(255) NULL,
	perido varchar(255) NULL,
	matricula varchar(255) NULL,
	usuario_id int NOT NULL,
	turma_id int NOT NULL,
	
	CONSTRAINT idalunos_pkey PRIMARY KEY (id),
	CONSTRAINT fk_usr_alunos_id FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id),
	CONSTRAINT fk_turma_alunos_id FOREIGN KEY (turma_id) REFERENCES public.trumas(id),
);



CREATE TABLE public.professores (
	id serial4 NOT NULL,
    nome varchar(255) NULL,
	cadeira varchar(255) NULL,
	usuario_id int NOT NULL,
	CONSTRAINT prof_pkey PRIMARY KEY (id),
	CONSTRAINT fk_usr_prof_id FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id),
	
);


-- public.codigos_promocionais_generatepromocionalcode definition

CREATE TABLE public.eventos (
	id serial4 NOT NULL,
	quantidade int2 NOT NULL,
	data_evento timestamptz NOT NULL,
	descricao varchar(255) NULL,
	turma_id int NOT NULL,
    professor_id int NULL,
	created_at time without time zone NOT NULL DEFAULT NOW(),
    updated_at time without time zone NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_prof_evento_id FOREIGN KEY (professor_id) REFERENCES public.professores(id),
	CONSTRAINT fk_turma_evento_id FOREIGN KEY (turma_id) REFERENCES public.trumas(id),
);
