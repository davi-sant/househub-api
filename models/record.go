package models

import "time"

type Record struct {
	ID           int       `json:"id"`
	RazaoSocial  string    `json:"razao_social"`
	NomeFantasia string    `json:"nome_fantasia"`
	CpfCnpj      string    `json:"cpf_cnpj"`
	Telefone     string    `json:"telefone"`
	Email        string    `json:"email"`
	Endereco     string    `json:"endereco"`
	CriadoEm     time.Time `json:"criado_em"`
}

type RecordCreate struct {
	RazaoSocial  string    `json:"razao_social" validate:"required,min=3,max=200"`
	NomeFantasia string    `json:"nome_fantasia" validate:"min=3,max=200"`
	CpfCnpj      string    `json:"cpf_cnpj" validate:"required,min=11,max=18"`
	Telefone     string    `json:"telefone" validate:"max=20"`
	Email        string    `json:"email" validate:"required,email"`
	Endereco     string    `json:"endereco" validate:"max=500"`
	CriadoEm     time.Time `json:"criado_em,omitempty"`
}

type RecordUpdate struct {
	RazaoSocial  string `json:"razao_social" validate:"required,min=3,max=200"`
	NomeFantasia string `json:"nome_fantasia" validate:"min=3,max=200"`
	Telefone     string `json:"telefone" validate:"required,min=11,max=20"`
	Email        string `json:"email" validate:"required,email"`
	Endereco     string `json:"endereco" validate:"max=500"`
}

type RecordResponse struct {
	Status   string   `json:"status"`
	Mensagem string   `json:"mensagem"`
	Dados    []Record `json:"dados,omitempty"`
	Record   *Record  `json:"record,omitempty"`
}

type ErrorItem struct {
	Erro string `json:"error"`
}
type RecordResponseError struct {
	Status   string      `json:"status"`
	Mensagem string      `json:"mensagem"`
	Erros    []ErrorItem `json:"erros,omitempty"`
}

type RecordById struct {
	ID string `json:"id" validate:"required"`
}
