package repositoryUsuarios

import (
	db "projeto404/Api/Database"
	"projeto404/Api/Models/ModelUsers"
)

func VerificarSoftdelete(email string) (bool, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var resultado modelUsuario.UsuarioStruct
	database.Unscoped().
		Where("email = ?", email).
		First(&resultado)

	// Verifique se a data de deletação está válida
	if resultado.DeletedAt.Valid {
		return true, nil
	}

	// Caso contrário, retorna false
	return false, nil
}
