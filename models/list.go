package models

type List struct {
	Id      int     `gorm:"primaryKey"`
	Title   string  `json:"Title" gorm:"size:100"`
	SubList SubList `gorm:"foreignKey:IdList; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SubList struct {
	Id              uint   `gorm:"primaryKey"`
	IdList          int    `json:"id_list"`
	Description     string `json:"description"`
	FileDescription []byte `json:"file_description"`
}

type List1 struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Title string `json:"Title" gorm:"size:100"`
}

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
