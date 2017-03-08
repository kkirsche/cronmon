package model

import (
	"net"
	"time"
)

// Host represents a network host
type Host struct {
	ID           int64
	Name         string    `json:"name"`
	IP           net.IP    `json:"ip_address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	LastUpdateBy string    `json:"last_updated_by"`
	Tasks        []Task    `json:"tasks"`
}
