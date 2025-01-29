package repositoryUsuarios

import (
	db "projeto404/Api/Database"
	modelUsuario "projeto404/Api/Models/ModelUsers"
)

func BuscarUsuario(idDoUsuario int) (*modelUsuario.UsuarioStruct, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var usuario modelUsuario.UsuarioStruct

	if err := database.Where("deleted_at IS NULL").First(&usuario, idDoUsuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
