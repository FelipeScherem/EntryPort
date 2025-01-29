-- Criar a extensão para permitir conexões de superusuário de ferramentas externas
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Configurar o timezone para UTC-3
SET timezone = 'UTC-3';

-- Criar a tabela de usuários
CREATE TABLE IF NOT EXISTS usuarios (
                                        id SERIAL PRIMARY KEY,
                                        nome VARCHAR(150) NOT NULL,
                                        email VARCHAR(150) NOT NULL UNIQUE,
                                        telefone VARCHAR(15) NOT NULL UNIQUE,
                                        senha VARCHAR(255) NOT NULL,
                                        data_de_nascimento DATE NOT NULL,
                                        foto VARCHAR(255),
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        deleted_at TIMESTAMP
);

-- Criar a tabela de mensagens
CREATE TABLE IF NOT EXISTS mensagens (
                                         id SERIAL PRIMARY KEY,
                                         remetente_id INT NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                                         destinatario_id INT NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                                         conteudo TEXT NOT NULL,
                                         data_envio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         deleted_at TIMESTAMP
);

-- Criar a tabela de conversas (opcional, para gerenciar interações entre dois usuários)
CREATE TABLE IF NOT EXISTS conversas (
                                         id SERIAL PRIMARY KEY,
                                         usuario1_id INT NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                                         usuario2_id INT NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
                                         ultima_mensagem TEXT,
                                         ultima_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         deleted_at TIMESTAMP
);

-- Adicionar exemplo de índice (opcional) para melhorar a busca por e-mail
CREATE INDEX IF NOT EXISTS idx_email ON usuarios (email);
