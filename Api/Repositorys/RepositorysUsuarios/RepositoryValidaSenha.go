package RepositoryUsuarios

import (
	db "projeto404/Api/Database"
	"projeto404/Api/Models/ModelUsers"
)

// ValidarSenhaAntiga verifica se a senha antiga está correta para o usuário
//
//	True se estiver ok
//	False se estiver errada
func ValidarSenha(idUsuario int, senhaAtual string) (bool, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var senhaNoBanco string
	query := database.Model(&modelUsuario.UsuarioStruct{}).
		Select("senha").
		Where("id = ?", idUsuario).
		Scan(&senhaNoBanco)
	// Verifica se houve erro na query
	if query.Error != nil {
		return false, query.Error
	}
	//Verifica se senha informada é igual a senha do banco
	if senhaAtual != senhaNoBanco {
		return false, nil
	} else {
		return true, nil
	}
}
