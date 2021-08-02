package domain

import "time"

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Comment     *string    `json:"comment,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (p *Product) SetCreateTimestamps() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = nil
}

func (p *Product) SetUpdateTimestamps() {
	now := time.Now()
	p.UpdatedAt = &now
}
