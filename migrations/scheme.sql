
CREATE TABLE if not exists registros (
                           id SERIAL PRIMARY KEY,
                           razao_social VARCHAR(200) NOT NULL,
                           nome_fantasia VARCHAR(200),
                           cnpj VARCHAR(18) UNIQUE NOT NULL,
                           telefone VARCHAR(20),
                           email VARCHAR(100),
                           endereco TEXT,
                           criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists filiais (
                         id SERIAL PRIMARY KEY,
                         nome VARCHAR(200) NOT NULL,
                         cnpj VARCHAR(18),
                         telefone VARCHAR(20),
                         email VARCHAR(100),
                         endereco TEXT,
                         registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                         ativa BOOLEAN DEFAULT TRUE,
                         criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists clientes (
                          id SERIAL PRIMARY KEY,
                          razao_social VARCHAR(200) NOT NULL,
                          nome_fantasia VARCHAR(200),
                          cnpj VARCHAR(18) UNIQUE NOT NULL,
                          telefone VARCHAR(20),
                          email VARCHAR(100),
                          endereco TEXT,
                          registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                          ativo BOOLEAN DEFAULT TRUE,
                          criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists usuarios (
                          id SERIAL PRIMARY KEY,
                          nome VARCHAR(100) NOT NULL,
                          email VARCHAR(100) UNIQUE NOT NULL,
                          senha VARCHAR(100) NOT NULL,
                          tipo VARCHAR(20) DEFAULT 'usuario', -- admin, usuario, etc.
                          filial_id INTEGER REFERENCES filiais(id) ON DELETE SET NULL,
                          registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                          ativo BOOLEAN DEFAULT TRUE,
                          criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists obras (
                       id SERIAL PRIMARY KEY,
                       nome VARCHAR(100) NOT NULL,
                       descricao TEXT,
                       data_inicio DATE,
                       data_fim_prevista DATE,
                       orcamento_total DECIMAL(15, 2),
                       cliente_id INTEGER NOT NULL REFERENCES clientes(id) ON DELETE CASCADE,
                       registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                       filial_id INTEGER NOT NULL REFERENCES filiais(id) ON DELETE CASCADE,
                       usuario_id INTEGER NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                       ativa BOOLEAN DEFAULT TRUE,
                       motivo_inativacao TEXT,
                       inativada_em TIMESTAMP,
                       inativada_por INTEGER REFERENCES usuarios(id),
                       criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists orcamentos (
                            id SERIAL PRIMARY KEY,
                            nome VARCHAR(100) NOT NULL,
                            descricao TEXT,
                            valor_total DECIMAL(15, 2) NOT NULL,
                            cliente_id INTEGER NOT NULL REFERENCES clientes(id) ON DELETE CASCADE,
                            obra_id INTEGER REFERENCES obras(id) ON DELETE SET NULL,
                            registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                            filial_id INTEGER NOT NULL REFERENCES filiais(id) ON DELETE CASCADE,
                            usuario_id INTEGER NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                            aprovado BOOLEAN DEFAULT FALSE,
                            data_aprovacao DATE,
                            criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists categorias (
                            id SERIAL PRIMARY KEY,
                            nome VARCHAR(100) NOT NULL,
                            descricao TEXT,
                            obra_id INTEGER NOT NULL REFERENCES obras(id) ON DELETE CASCADE,
                            registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                            criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists etapas_obra (
                             id SERIAL PRIMARY KEY,
                             nome VARCHAR(100) NOT NULL,
                             descricao TEXT,
                             ordem INTEGER DEFAULT 0,
                             obra_id INTEGER NOT NULL REFERENCES obras(id) ON DELETE CASCADE,
                             registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                             data_inicio DATE,
                             data_fim_prevista DATE,
                             data_conclusao DATE,
                             concluida BOOLEAN DEFAULT FALSE,
                             criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE if not exists gastos (
                        id SERIAL PRIMARY KEY,
                        descricao VARCHAR(255) NOT NULL,
                        valor DECIMAL(10, 2) NOT NULL,
                        data DATE NOT NULL,
                        obra_id INTEGER NOT NULL REFERENCES obras(id) ON DELETE CASCADE,
                        categoria_id INTEGER REFERENCES categorias(id) ON DELETE SET NULL,
                        etapa_id INTEGER REFERENCES etapas_obra(id) ON DELETE SET NULL,
                        orcamento_id INTEGER REFERENCES orcamentos(id) ON DELETE SET NULL,
                        registro_id INTEGER NOT NULL REFERENCES registros(id) ON DELETE CASCADE,
                        usuario_id INTEGER NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                        criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);