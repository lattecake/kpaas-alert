package db

import "github.com/jinzhu/gorm"

type Alert struct {
	gorm.Model
	Namespace string
	Name      string
	Body      string
}

func Create(alert *Alert) bool {
	return GetMysqlSession().NewRecord(alert)
}

func Get(id int64) (alert *Alert, err error) {

	db := GetMysqlSession().First(&alert, id)
	if db.Error != nil {
		return nil, db.Error
	}

	return
}

func Update() {

}

func Delete() {

}
