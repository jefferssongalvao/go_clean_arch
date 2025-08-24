package entities

import valueobjects "github.com/jefferssongalvao/go_clean_arch/internal/domain/value_objects"

type User struct {
	ID       uint                  `json:"id"`
	Username string                `json:"username"`
	Password valueobjects.Password `json:"-"`
}
