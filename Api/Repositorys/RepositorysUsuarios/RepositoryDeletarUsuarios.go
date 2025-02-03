package RepositoryUsuarios

import (
	db "projeto404/Api/Database"
	"projeto404/Api/Models/ModelUsers"
)

func DeletarUsuario(usuarioStruct modelUsuario.UsuarioStruct, id int) (string, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	if err := database.
		Delete(&usuarioStruct, id).
		Error; err != nil {
		return "Erro ao excluir o usuário:", err
	} else {
		return "Usuario excluido com sucesso", nil
	}
}
