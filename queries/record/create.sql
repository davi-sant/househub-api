INSERT INTO registros (razao_social, nome_fantasia, cnpj, telefone, email, endereco)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, razao_social, nome_fantasia, cnpj, telefone, email, endereco, criado_em;
