package models

type Application struct {
	ApplicationID int64
	Name          string `sql:"size:255"`
	Description   string `sql:"size:255"`
}
