package RepositoryUsuarios

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

func ListarUsuarios() (*[]modelUsuario.UsuarioStruct, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var usuarios []modelUsuario.UsuarioStruct
	if err := database.Where("deleted_at IS NULL").Find(&usuarios).Error; err != nil {
		return nil, err
	}

	return &usuarios, nil
}

func BuscarUsuarioLogin(emailDoUsuario string) (*modelUsuario.UsuarioStruct, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var usuario modelUsuario.UsuarioStruct

	if err := database.Where("deleted_at IS NULL AND email = ?", emailDoUsuario).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
