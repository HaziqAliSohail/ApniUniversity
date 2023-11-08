package service

import (
	"encoding/json"

	"ApniUniversity/data"
	"ApniUniversity/models"
)

func (s *Service) AddTeacher(teacher *models.Teacher) (int, error) {
	teachers, err := s.db.GetTeachers()
	if err != nil {

		return 0, err
	}

	teacher.ID = teachers[len(teachers)-1].ID + 1

	if teacher.SubjectID != 0 {
		if _, err = s.db.GetSubject(teacher.SubjectID); err != nil {

			return 0, err
		}
	}
	return s.db.AddOrUpdateTeacher(teacher)
}

func (s *Service) GetTeachers() ([]*models.Teacher, error) {

	return s.db.GetTeachers()
}

func (s *Service) UpdateTeacherName(id int, body map[string]string) (int, error) {
	teacher, err := s.db.GetTeacher(id)
	if err != nil {

		return 0, err
	}
	teacher.Name = body["name"]

	return s.db.AddOrUpdateTeacher(teacher)
}

func (s *Service) AssignSubjectToTeacher(id int, body map[string]int) (int, error) {
	teacher, err := s.db.GetTeacher(id)
	if err != nil {

		return 0, err
	}

	teacher.SubjectID = body["subjectID"]

	if teacher.SubjectID != 0 {
		if _, err = s.db.GetSubject(teacher.SubjectID); err != nil {

			return 0, err
		}
	}

	return s.db.AddOrUpdateTeacher(teacher)
}

func (s *Service) GetClassOfTeacher(id int) (*models.Class, error) {
	teacher, err := s.db.GetTeacher(id)
	if err != nil {

		return nil, err
	}

	subject, err := s.db.GetSubject(teacher.SubjectID)
	if err != nil {

		return nil, err
	}

	class, err := s.db.GetClass(subject.ClassID)
	if err != nil {

		return nil, err
	}

	return class, nil
}

func (s *Service) GetStudentsOfTeacher(id int) ([]*models.Student, error) {
	teacher, err := s.db.GetTeacher(id)
	if err != nil {

		return nil, err
	}

	subject, err := s.db.GetSubject(teacher.SubjectID)
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

func (s *Service) DeleteTeacher(id int) (string, error) {
	message, err := s.db.DeleteTeacher(id)
	if err != nil {

		return "", err
	}

	accounts, err := s.db.GetAccounts()
	if err != nil {
		return "", err
	}

	for _, account := range accounts {
		if account.AccountType == data.TEACHER {
			var accountData *models.TeacherAccount
			dBytes, _ := json.Marshal(account.AccountData)
			_ = json.Unmarshal(dBytes, &accountData)

			if accountData.TeacherID == id {
				accountData.TeacherID = 0
				account.AccountData = accountData

				_, err = s.db.AddOrUpdateAccount(account)
				if err != nil {

					return "", err
				}

				break
			}

		}
	}

	return message, nil
}

func (s *Service) GetTeacherByID(id int) (*models.Teacher, error) {

	return s.db.GetTeacher(id)
}
