package definitions

import (
	"encoding/json"
	"time"
)

type Notification struct {
	Token        string          `json:"token,omitempty"`
	EventID      string          `json:"event_id,omitempty"`
	Notification json.RawMessage `json:"notification,omitempty"`
	Resource     string          `json:"resource,omitempty"`
	ResourcePath string          `json:"resource_path,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty"`
}
