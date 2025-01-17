package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateSubaccountRequest struct {
	// Nome da subconta (REQUIRED)
	Name string `json:"name,omitempty"`
	// E-mail da subconta (REQUIRED)
	Email string `json:"email,omitempty"`
	// E-mail para login da subconta, caso não informado será utilizado o email da subconta
	LoginEmail string `json:"loginEmail,omitempty"`
	// CPF ou CNPJ do proprietário da subconta
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Data de nascimento (somente quando Pessoa Física)
	BirthDate Date `json:"birthDate,omitempty"`
	// Tipo da empresa (somente quando Pessoa Jurídica)
	CompanyType CompanyType `json:"companyType,omitempty"`
	// Fone fixo
	Phone string `json:"phone,omitempty"`
	// Fone celular (REQUIRED)
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Enviar URL referente ao site da conta filha
	Site string `json:"site,omitempty"`
	// Logradouro (REQUIRED)
	Address string `json:"address,omitempty"`
	// Número do endereço (REQUIRED)
	AddressNumber string `json:"addressNumber,omitempty"`
	// Complemento do endereço
	Complement string `json:"complement,omitempty"`
	// Bairro (REQUIRED)
	Province string `json:"province,omitempty"`
	// CEP do endereço (REQUIRED)
	PostalCode string `json:"postalCode,omitempty"`

	// Valor do rendimento
	IncomeValue float64 `json:"incomeValue,omitempty"`

	// Array com as configurações de webhooks desejadas
	Webhooks []CreateSubaccountWebhookRequest `json:"webhooks,omitempty"`
}

type GetAllSubaccountsRequest struct {
	// Filtrar pelo cpf ou cnpj da subconta
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Filtrar pelo email da subconta
	Email string `json:"email,omitempty"`
	// Filtrar pelo nome da subconta
	Name string `json:"name,omitempty"`
	// Filtrar pelo walletId da subconta
	WalletId string `json:"walletId,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type CreateSubaccountWebhookRequest struct {
	// URL que receberá as informações de sincronização (REQUIRED)
	Url string `json:"url,omitempty"`
	// E-mail para receber as notificações em caso de erros na fila (REQUIRED)
	Email string `json:"email,omitempty"`
	// Versão utilizada da API. Utilize "3" para a versão v3 (REQUIRED)
	ApiVersion string `json:"apiVersion,omitempty"`
	// Habilitar ou não o webhook
	Enabled bool `json:"enabled"`
	// Situação da fila de sincronização
	Interrupted bool `json:"interrupted"`
	// Token de autenticação
	AuthToken string `json:"authToken,omitempty"`
	// Tipo de webhook
	Type WebhookType `json:"type,omitempty"`
}

type SendWhiteLabelDocumentRequest struct {
	// Tipo de documento (REQUIRED)
	Type SubaccountDocumentType `json:"type,omitempty"`
	// Arquivo (REQUIRED)
	DocumentFile *os.File `json:"documentFile,omitempty"`
}

type UpdateWhiteLabelDocumentSentRequest struct {
	// Arquivo (REQUIRED)
	DocumentFile *os.File `json:"documentFile,omitempty"`
}

type SubaccountResponse struct {
	Id            string                  `json:"id,omitempty"`
	Name          string                  `json:"name,omitempty"`
	PersonType    PersonType              `json:"personType,omitempty"`
	Email         string                  `json:"email,omitempty"`
	LoginEmail    string                  `json:"loginEmail,omitempty"`
	CpfCnpj       string                  `json:"cpfCnpj,omitempty"`
	BirthDate     Date                    `json:"birthDate,omitempty"`
	CompanyType   CompanyType             `json:"companyType,omitempty"`
	Phone         string                  `json:"phone,omitempty"`
	MobilePhone   string                  `json:"mobilePhone,omitempty"`
	Site          string                  `json:"site,omitempty"`
	Address       string                  `json:"address,omitempty"`
	AddressNumber string                  `json:"addressNumber,omitempty"`
	Complement    string                  `json:"complement,omitempty"`
	Province      string                  `json:"province,omitempty"`
	PostalCode    string                  `json:"postalCode,omitempty"`
	City          int                     `json:"city,omitempty"`
	Country       string                  `json:"country,omitempty"`
	ApiKey        string                  `json:"apiKey,omitempty"`
	WalletId      string                  `json:"walletId,omitempty"`
	AccountNumber AccountBankInfoResponse `json:"accountNumber,omitempty"`
	Errors        []ErrorResponse         `json:"errors,omitempty"`
}

type SubaccountDocumentSentResponse struct {
	Id     string                   `json:"id,omitempty"`
	Status SubaccountDocumentStatus `json:"status,omitempty"`
	Errors []ErrorResponse          `json:"errors,omitempty"`
}

type SubaccountDocumentsResponse struct {
	RejectReasons string                       `json:"rejectReasons,omitempty"`
	Data          []SubaccountDocumentResponse `json:"data,omitempty"`
	Errors        []ErrorResponse              `json:"errors,omitempty"`
}

type SubaccountDocumentResponse struct {
	Id          string                      `json:"id,omitempty"`
	Status      SubaccountDocumentStatus    `json:"status,omitempty"`
	Type        SubaccountDocumentType      `json:"type,omitempty"`
	Title       string                      `json:"title,omitempty"`
	Description string                      `json:"description,omitempty"`
	Responsible DocumentResponsibleResponse `json:"responsible,omitempty"`
	Documents   []any                       `json:"documents,omitempty"`
}

type DocumentResponsibleResponse struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type subaccount struct {
	env         Env
	accessToken string
}

type Subaccount interface {
	// Create (Criar subconta)
	//
	// O objeto de retorno da API conterá a chave de API da subconta criada (SubaccountResponse.ApiKey) além do
	// SubaccountResponse.WalletId para Split de Cobranças ou Transferências.
	//
	// A chave de API (SubaccountResponse.ApiKey) será devolvida uma única vez, na resposta da chamada de criação
	// da subconta Asaas, portanto, assegure-se de gravar a informação nesse momento.
	// Caso não tenha realizado o armazenamento, entre em contato com nosso Suporte Técnico.
	//
	// # Resposta: 200
	//
	// SubaccountResponse = not nil
	//
	// Error = nil
	//
	// SubaccountResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// SubaccountResponse = not nil
	//
	// Error = nil
	//
	// SubaccountResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Criar subconta: https://docs.asaas.com/reference/criar-subconta
	Create(ctx context.Context, body CreateSubaccountRequest) (*SubaccountResponse, error)
	// SendWhiteLabelDocument (Enviar documentos via API)
	//
	// Quando houver o atributo onboardingUrl no objeto do documento, ele deverá ser enviado via link externo.
	// Não será aceito o envio via POST nesses casos.
	//
	// # Resposta: 200
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountDocumentSentResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountDocumentSentResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Enviar documentos via API: https://docs.asaas.com/reference/enviar-documentos-via-api
	SendWhiteLabelDocument(ctx context.Context, documentId string, body SendWhiteLabelDocumentRequest) (
		*SubaccountDocumentSentResponse, error)
	// UpdateWhiteLabelDocumentSentById (Atualizar documento enviado)
	//
	// # Resposta: 200
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountDocumentSentResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountDocumentSentResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Atualizar documento enviado: https://docs.asaas.com/reference/atualizar-documento-enviado
	UpdateWhiteLabelDocumentSentById(ctx context.Context, documentSentId string, body UpdateWhiteLabelDocumentSentRequest) (
		*SubaccountDocumentSentResponse, error)
	// DeleteWhiteLabelDocumentSentById (Remover documento enviado)
	//
	// # Resposta: 200
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// Se DeleteResponse.IsSuccess() for true quer dizer que foi excluída.
	//
	// Se caso DeleteResponse.IsFailure() for true quer dizer que não foi excluída.
	//
	// # Resposta: 404
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo DeleteResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// DeleteResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Remover documento enviado: https://docs.asaas.com/reference/remover-documento-enviado
	DeleteWhiteLabelDocumentSentById(ctx context.Context, documentSentId string) (*DeleteResponse, error)
	// GetById (Recuperar uma única subconta)
	//
	// # Resposta: 200
	//
	// SubaccountResponse = not nil
	//
	// Error = nil
	//
	// SubaccountResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// SubaccountResponse = not nil
	//
	// Error = nil
	//
	// SubaccountResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// SubaccountResponse = not nil
	//
	// Error = nil
	//
	// SubaccountResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Recuperar uma única subconta: https://docs.asaas.com/reference/recuperar-uma-unica-subconta
	GetById(ctx context.Context, subaccountId string) (*SubaccountResponse, error)
	// GetDocumentSentById (Visualizar documento enviado)
	//
	// # Resposta: 200
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// SubaccountDocumentSentResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentSentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountDocumentSentResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountDocumentSentResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Visualizar documento enviado: https://docs.asaas.com/reference/visualizar-documento-enviado
	GetDocumentSentById(ctx context.Context, documentSentId string) (*SubaccountDocumentSentResponse, error)
	// GetAll (Listar subcontas)
	//
	// # Resposta: 200
	//
	// Pageable(SubaccountResponse) = not nil
	//
	// Error = nil
	//
	// Se Pageable.IsSuccess() for true quer dizer que retornaram os dados conforme a documentação.
	// Se Pageable.IsNoContent() for true quer dizer que retornou os dados vazio.
	//
	// Error = nil
	//
	// Pageable.IsNoContent() = true
	//
	// Pageable.Data retornou vazio.
	//
	// # Resposta: 401/500
	//
	// Pageable(SubaccountResponse) = not nil
	//
	// Error = nil
	//
	// Pageable.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo Pageable.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// Pageable(SubaccountResponse) = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Listar subcontas: https://docs.asaas.com/reference/listar-subcontas
	GetAll(ctx context.Context, filter GetAllSubaccountsRequest) (*Pageable[SubaccountResponse], error)
	// GetPendingDocuments (Verificar documentos pendentes)
	//
	// Para recuperar os documentos pendentes e ter acesso ao onboardingUrl dos mesmos.
	//
	// # Resposta: 200
	//
	// SubaccountDocumentsResponse = not nil
	//
	// Error = nil
	//
	// Se SubaccountDocumentsResponse.IsSuccess() for true quer dizer que retornaram os dados conforme a documentação.
	// Se SubaccountDocumentsResponse.IsNoContent() for true quer dizer que retornou os dados vazio.
	//
	// Error = nil
	//
	// SubaccountDocumentsResponse.IsNoContent() = true
	//
	// SubaccountDocumentsResponse.Data retornou vazio.
	//
	// # Resposta: 401/500
	//
	// SubaccountDocumentsResponse = not nil
	//
	// Error = nil
	//
	// SubaccountDocumentsResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubaccountDocumentsResponse.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubaccountDocumentsResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/jmjp/go-asaas
	//
	// # DOCS
	//
	// Verificar documentos pendentes: https://docs.asaas.com/reference/verificar-documentos-pendentes
	GetPendingDocuments(ctx context.Context) (*SubaccountDocumentsResponse, error)
}

func NewSubaccount(env Env, accessToken string) Subaccount {
	logWarning("Subaccount service running on", env.String())
	return subaccount{
		env:         env,
		accessToken: accessToken,
	}
}

func (s subaccount) Create(ctx context.Context, body CreateSubaccountRequest) (*SubaccountResponse, error) {
	req := NewRequest[SubaccountResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, "/v3/accounts", body)
}

func (s subaccount) SendWhiteLabelDocument(ctx context.Context, documentId string, body SendWhiteLabelDocumentRequest) (
	*SubaccountDocumentSentResponse, error) {
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/myAccount/documents/%s", documentId), body)
}

func (s subaccount) UpdateWhiteLabelDocumentSentById(
	ctx context.Context,
	documentSentId string,
	body UpdateWhiteLabelDocumentSentRequest,
) (*SubaccountDocumentSentResponse, error) {
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.makeMultipartForm(http.MethodPut, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), body)
}

func (s subaccount) DeleteWhiteLabelDocumentSentById(ctx context.Context, documentSentId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), nil)
}

func (s subaccount) GetById(ctx context.Context, subaccountId string) (*SubaccountResponse, error) {
	req := NewRequest[SubaccountResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/accounts/%s", subaccountId), nil)
}

func (s subaccount) GetDocumentSentById(ctx context.Context, documentSentId string) (*SubaccountDocumentSentResponse,
	error) {
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), nil)
}

func (s subaccount) GetAll(ctx context.Context, filter GetAllSubaccountsRequest) (*Pageable[SubaccountResponse], error) {
	req := NewRequest[Pageable[SubaccountResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/accounts", filter)
}

func (s subaccount) GetPendingDocuments(ctx context.Context) (*SubaccountDocumentsResponse, error) {
	req := NewRequest[SubaccountDocumentsResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/documents", nil)
}
