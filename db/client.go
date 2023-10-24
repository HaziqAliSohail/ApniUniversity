package db

import (
	"ApniUniversity/models"
)

type DataStore interface {
	AddOrUpdateTeacher(teacher *models.Teacher) (int, error)
	GetTeacher(id int) (*models.Teacher, error)
	GetTeachers() ([]*models.Teacher, error)
	DeleteTeacher(id int) (string, error)
	AddOrUpdateSubject(subject *models.Subject) (int, error)
	GetSubject(id int) (*models.Subject, error)
	GetSubjects() ([]*models.Subject, error)
	DeleteSubject(id int) (string, error)
}
type Options struct {
	TestMode bool
}
