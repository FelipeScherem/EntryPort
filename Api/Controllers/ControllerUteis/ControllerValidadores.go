package ControllerUteis

import (
	"regexp"
)

// ValidarSenha Valida se senha é segura
//
//	8 ou mais digitos
//	Ao menos 1 caracter maiusculo e minusculo
//	1 numeral
//	1 caracter especial
//
//	True se ok
//	False se há erros
func ValidarSenha(senha string) (string, bool) {

	if len(senha) < 8 {
		return "A senha deve conter ao menos 8 caracteres", false
	}

	contemMaiuscula := false
	contemMinuscula := false
	for _, char := range senha {
		if 'A' <= char && char <= 'Z' {
			contemMaiuscula = true
		}
		if 'a' <= char && char <= 'z' {
			contemMinuscula = true
		}
	}
	if !contemMaiuscula || !contemMinuscula {
		return "A senha deve conter ao menos uma letra maiuscula e uma minuscula", false
	}

	contemNumeral := false
	for _, char := range senha {
		if '0' <= char && char <= '9' {
			contemNumeral = true
		}
	}
	if !contemNumeral {
		return "A senha deve conter ao menos um numeral", false
	}

	regexEspecial := regexp.MustCompile(`[!@#$%^&*()_+{}[\]:;<>,.?/~]`)
	if !regexEspecial.MatchString(senha) {
		return "A senha deve conter ao menos um caracter especial", false
	}

	return "", true
}
