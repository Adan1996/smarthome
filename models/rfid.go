package models

type Rfid struct {
	ID        int    `gorm:"type:primaryKey" json:"id"`
	PICC_type string `gorm:"type:varchar(50)" json:"picc_type"`
	UID       string `gorm:"type:varchar(20)" json:"uid"`
}
