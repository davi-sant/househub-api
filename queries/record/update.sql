UPDATE registros
SET	razao_social = $1,
	nome_fantasia = $2,
	telefone = $3,
	email = $4,
	endereco = $5
WHERE id = $6
RETURNING id, razao_social, nome_fantasia, cnpj, telefone, email, endereco, criado_em;
