package service

import (
	"github.com/pkg/errors"

	"github.com/HaziqAliSohail/ApniUniversity/models"
)

func (s *Service) AddSubject(subject *models.Subject) (int, error) {
	subjects, err := s.db.GetSubjects()
	if err != nil {

		return 0, err
	}

	if len(subjects) != 0 {
		subject.ID = subjects[len(subjects)-1].ID + 1
	} else {
		subject.ID = 1
	}

	if subject.ClassID != 0 {
		if _, err = s.db.GetClass(subject.ClassID); err != nil {

			return 0, err
		}
	}

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

	if subject.ClassID != 0 {
		if _, err = s.db.GetClass(subject.ClassID); err != nil {

			return 0, err
		}
	}

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
	message, err := s.db.DeleteSubject(id)
	if err != nil {
		return "", err
	}

	students, err := s.db.GetStudents()
	if err != nil {
		return "", err
	}

	for _, student := range students {
		for i, subjectID := range student.Subjects {
			if subjectID == id {
				student.Subjects = append(student.Subjects[:i], student.Subjects[i+1:]...)
				_, err := s.db.AddOrUpdateStudent(student)
				if err != nil {

					return "", errors.Wrap(err, "Delete Subject: Subject not removed from the student Data!")
				}

				break
			}

		}
	}

	teachers, err := s.db.GetTeachers()
	if err != nil {
		return "", err
	}

	for _, teacher := range teachers {
		if teacher.SubjectID == id {
			teacher.SubjectID = 0
			_, err := s.db.AddOrUpdateTeacher(teacher)
			if err != nil {

				return "", errors.Wrap(err, "Delete Subject: Subject not removed from the Teacher's Data!")
			}

			break
		}

	}

	return message, nil
}

func (s *Service) GetSubjectByID(id int) (*models.Subject, error) {

	return s.db.GetSubject(id)
}
