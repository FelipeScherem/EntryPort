package util

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Define os erros comuns do Gorm
var (
	ErrRecordNotFound = errors.New("registro não encontrado")
	ErrDuplicateEntry = errors.New("entrada duplicada")
	ErrInvalidData    = errors.New("dados inválidos")
)

// MapGormError mapeia erros do Gorm para erros definidos pelo usuário
func MapGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	// Exemplo: Verificar erro de entrada duplicada usando código SQLState
	// Isso pode variar dependendo do banco de dados que você está usando
	sqlErr := fmt.Sprintf("%v", err)
	if sqlErr == "Error 1062: Duplicate entry" {
		return ErrDuplicateEntry
	}

	return err
}

// HandleError formata uma mensagem de erro para o usuário
func HandleError(err error) string {
	switch err {
	case ErrRecordNotFound:
		return "Erro: O registro solicitado não foi encontrado."
	case ErrDuplicateEntry:
		return "Erro: A entrada já existe no banco de dados."
	case ErrInvalidData:
		return "Erro: Os dados fornecidos são inválidos."
	default:
		return fmt.Sprintf("Erro desconhecido: %v", err)
	}
}
