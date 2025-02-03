package RepositoryUsuarios

import (
	"github.com/jackc/pgx/v5/pgconn"
	db "projeto404/Api/Database"
	"projeto404/Api/Models/ModelUsers"
	"strings"
)

// CriarUsuario insere os dados do cliente no banco
func CriarUsuario(usuarioStruct modelUsuario.UsuarioStruct) (string, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	// Tenta fazer a inserção, se houver dados repetidos, ele retorna com os campos repetidos
	if err := database.
		Create(&usuarioStruct).
		Error; err != nil {

		camposRepetidos := "Campos repetidos: "
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 É o código de erro padrão do postgresql unique_violation
			if strings.Contains(pgErr.Detail, "email") {
				camposRepetidos += "email, "
			}
			if strings.Contains(pgErr.Detail, "telefone") {
				camposRepetidos += "telefone, "
			}
			if strings.Contains(pgErr.Detail, "cpf") {
				camposRepetidos += "cpf, "
			}
			if strings.Contains(pgErr.Detail, "cnpj") {
				camposRepetidos += "cnpj, "
			}
			// Remova a vírgula extra no final, se houver
			camposRepetidos = camposRepetidos[:len(camposRepetidos)-2]

		}
		return camposRepetidos, err
	}

	// Se esta tudo ok, retorna
	return "", nil
}
