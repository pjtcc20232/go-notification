
-- Inserir o registro para 'aluno' com permissão 'alunos'
INSERT INTO public.grupo (grupo, permissoes)
VALUES ('aluno', 'alunos');

-- Inserir o registro para 'professores' com permissão 'professores'
INSERT INTO public.grupo (grupo, permissoes)
VALUES ('professores', 'professores');

-- Inserir o registro para 'adm' com permissão 'adimin'
INSERT INTO public.grupo (grupo, permissoes)
VALUES ('adm', 'adimin');