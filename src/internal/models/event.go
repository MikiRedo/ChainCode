package models

// Event representa la estructura de un evento
type Event struct {
	CorrelationID string   `json:"correlationId" validate:"required"`
	ID            []string `json:"id" validate:"required"`
	Type          string   `json:"type" validate:"required"`
	Payload       []byte   `json:"payload" validate:"required"`
}

//he intentado añadair el otro struct aquí, pero me da error el go mod