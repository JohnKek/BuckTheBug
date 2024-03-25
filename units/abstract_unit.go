package units

import "github.com/google/uuid"

type Unit struct {
	ID                  uuid.UUID
	X                   float64
	Y                   float64
	SpriteName          string
	Action              string
	Frame               int
	HorizontalDirection int
	HealPoint           float64
}
