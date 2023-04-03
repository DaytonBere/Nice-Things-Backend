package models

import "gorm.io/gorm"

type NiceThing struct {
	gorm.Model
	Sender int
	Receiver int
	Body string
}