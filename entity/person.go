package entity

const PersonTableName = "person_new"

type Person struct {
	BaseModel
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
	City string `json:"city" gorm:"type:varchar(100);"`
	Age  int    `json:"age" gorm:"type:int(10)"`
}

func (e *Person) TableName() string {
	return PersonTableName
}
