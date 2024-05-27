package entity

type Role struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RoleUpdate struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
