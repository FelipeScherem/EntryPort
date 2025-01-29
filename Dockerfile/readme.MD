# Passos para Configuração do Contêiner PostgreSQL

## Descrição

Com o objetivo de deixar o processo o mais simples possível, automatizei alguns processos.  
Basta seguir os comandos abaixo para ter seu contêiner rodando.

1. **Crie a imagem Docker**
    ```powershell
    docker build -t entry_project .
    ```

2. **Inicie o contêiner**
    ```powershell
    docker run -d --name entry_project -p 5432:5432 entry_project
    ```

3. **Verifique os contêineres em execução**
    ```powershell
    docker ps
    ```

4. **Entre no contêiner em modo interativo**
    ```powershell
    docker exec -it entry_project bash
    ```

5. **Acesse o PostgreSQL**

   Dentro do contêiner, acesse o PostgreSQL com o seguinte comando:
    ```bash
    psql -U under_admin -d under_hex
    ```

6. **Execute Comandos SQL**

   Agora você pode executar os comandos SQL no prompt do PostgreSQL. Alguns exemplos:

    - **Liste as tabelas no banco de dados:**
      ```sql
      \dt
      ```

    - **Veja dados da tabela `u_categoria`:**
      ```sql
      SELECT * FROM u_categoria;
      ```

    - **Veja dados da tabela `u_produto`:**
      ```sql
      SELECT * FROM u_produto;
      ```

7. **Saia do shell do PostgreSQL**
    ```sql
    \q
    ```
