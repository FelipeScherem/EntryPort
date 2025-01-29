package modelUsuario

import (
	"gorm.io/gorm"
	"time"
)

// UsuarioStruct Struct com os dados de usuarios
type UsuarioStruct struct {
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	Nome             string         `gorm:"column:nome"`
	Email            string         `gorm:"column:email;unique"`
	Telefone         string         `gorm:"column:telefone;unique"`
	Senha            string         `gorm:"column:senha"`
	DataDeNascimento time.Time      `gorm:"column:data_de_nascimento"`
	Foto             string         `gorm:"column:foto"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName define o nome da tabela no banco de dados
func (UsuarioStruct) TableName() string {
	return "usuarios"
}
