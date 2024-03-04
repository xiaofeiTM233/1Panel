package model

type Firewall struct {
	BaseModel

	Type        string `gorm:"type:varchar(64);not null" json:"type"`
	Port        string `gorm:"type:varchar(64);not null" json:"port"`
	Protocol    string `gorm:"type:varchar(64);not null" json:"protocol"`
	Address     string `gorm:"type:varchar(64);not null" json:"address"`
	Strategy    string `gorm:"type:varchar(64);not null" json:"strategy"`
	Description string `gorm:"type:varchar(64);not null" json:"description"`
}
