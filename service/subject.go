package service

import (
	"ApniUniversity/models"
	"github.com/pkg/errors"
)

func (s *Service) AddSubject(subject *models.Subject) (int, error) {
	subjects, err := s.db.GetSubjects()
	if err != nil {

		return 0, err
	}

	subject.ID = subjects[len(subjects)-1].ID + 1

	return s.db.AddOrUpdateSubject(subject)
}

func (s *Service) GetSubjects() ([]*models.Subject, error) {

	return s.db.GetSubjects()
}

func (s *Service) UpdateSubjectName(id int, body map[string]string) (int, error) {
	subject, err := s.db.GetSubject(id)
	if err != nil {

		return 0, err
	}
	subject.Name = body["name"]

	return s.db.AddOrUpdateSubject(subject)
}

func (s *Service) AssignClass(id int, body map[string]int) (int, error) {
	subject, err := s.db.GetSubject(id)
	if err != nil {

		return 0, err
	}

	subject.ClassID = body["classID"]

	return s.db.AddOrUpdateSubject(subject)
}

func (s *Service) GetTeacherOfSubject(id int) (*models.Teacher, error) {
	teachers, err := s.db.GetTeachers()
	if err != nil {

		return nil, err
	}
	for _, teacher := range teachers {
		if teacher.SubjectID == id {

			return teacher, nil
		}

	}

	return nil, errors.Errorf("No Teacher enrolled in this subject!")
}

func (s *Service) GetStudentsOfSubject(id int) ([]*models.Student, error) {
	subject, err := s.db.GetSubject(id)
	if err != nil {

		return nil, err
	}

	class, err := s.db.GetClass(subject.ClassID)
	if err != nil {

		return nil, err
	}

	var students []*models.Student
	for _, studentID := range class.Students {
		student, err := s.db.GetStudent(studentID)
		if err != nil {

			return nil, err
		}

		for _, sub := range student.Subjects {
			if sub == subject.ID {
				students = append(students, student)
				break
			}

		}
	}

	return students, nil
}

func (s *Service) DeleteSubject(id int) (string, error) {

	return s.db.DeleteSubject(id)
}

func (s *Service) GetSubjectByID(id int) (*models.Subject, error) {

	return s.db.GetSubject(id)
}
