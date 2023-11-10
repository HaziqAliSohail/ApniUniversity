package db

import (
	"github.com/HaziqAliSohail/ApniUniversity/models"
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

	AddOrUpdateClass(class *models.Class) (int, error)
	GetClass(id int) (*models.Class, error)
	GetClasses() ([]*models.Class, error)
	DeleteClass(id int) (string, error)

	AddOrUpdateStudent(student *models.Student) (int, error)
	GetStudent(id int) (*models.Student, error)
	GetStudents() ([]*models.Student, error)
	DeleteStudent(id int) (string, error)

	AddOrUpdateAccount(account *models.Account) (int, error)
	GetAccount(id int) (*models.Account, error)
	GetAccounts() ([]*models.Account, error)
	DeleteAccount(id int) (string, error)
}
type Options struct {
	TestMode bool
}
