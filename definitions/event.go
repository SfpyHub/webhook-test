package definitions

import "time"

type Event struct {
	Token            string        `json:"token,omitempty"`
	ClientID         string        `json:"client_id"`
	Type             string        `json:"type,omitempty"`
	Status           string        `json:"status,omitempty"`
	Endpoint         string        `json:"endpoint,omitempty"`
	Data             *Notification `json:"data,omitempty"`
	DeliveryAttempts int           `json:"delivery_attempts,omitempty"`
	Resource         string        `json:"resource,omitempty"`
	ResourcePath     string        `json:"resource_path,omitempty"`
	NextAttemptAt    time.Time     `json:"next_attempt_at,omitempty"`
	CreatedAt        time.Time     `json:"created_at,omitempty"`
	UpdatedAt        time.Time     `json:"updated_at,omitempty"`
}
