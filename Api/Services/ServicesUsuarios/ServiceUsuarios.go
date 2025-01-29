package serviceUsuarios

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ConsultarDadosCNPJ Espera receber o CNPJ para verificar
//
//	Retorna o nomeFantasia e razãoSocial
func ConsultarDadosCNPJ(cnpj string) (string, string, error, string) {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	consultaURL := fmt.Sprintf("https://receitaws.com.br/v1/cnpj/%s", cnpj)

	resp, err := http.Get(consultaURL)
	if err != nil {
		return "", "", err, "Erro ao consultar os dados do CNPJ"
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", "", err, "Erro ao consultar os dados do CNPJ"
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", "", err, "Erro ao consultar os dados do CNPJ"
	}

	// Verifica se os campos estão presentes antes de acessá-los
	fantasia, existeNomeFantasia := data["fantasia"].(string)
	razaoSocial, existerazaoSocial := data["nome"].(string)

	if !existeNomeFantasia {
		return "", "", nil, "Campo 'fantasia' não encontrado na resposta"
	}
	if !existerazaoSocial {
		return "", "", nil, "Campo 'nome' não encontrado na resposta"
	}

	return fantasia, razaoSocial, nil, ""
}
