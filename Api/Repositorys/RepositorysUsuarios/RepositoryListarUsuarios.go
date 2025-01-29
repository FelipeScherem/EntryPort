package repositoryUsuarios

import (
	db "projeto404/Api/Database"
	modelUsuario "projeto404/Api/Models/ModelUsers"
)

func ListarUsuarios() (*[]modelUsuario.UsuarioStruct, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var usuarios []modelUsuario.UsuarioStruct
	if err := database.Where("deleted_at IS NULL").Find(&usuarios).Error; err != nil {
		return nil, err
	}

	return &usuarios, nil
}
