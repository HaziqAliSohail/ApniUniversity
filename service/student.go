package service

import (
	"encoding/json"

	"github.com/pkg/errors"

	"ApniUniversity/data"
	"ApniUniversity/models"
)

func (s *Service) AddStudent(student *models.Student) (int, error) {
	students, err := s.db.GetStudents()
	if err != nil {

		return 0, err
	}

	student.ID = students[len(students)-1].ID + 1

	if len(student.Subjects) > 0 {
		for _, subjectID := range student.Subjects {

			if _, err = s.db.GetSubject(subjectID); err != nil {

				return 0, err
			}
		}
	}

	return s.db.AddOrUpdateStudent(student)
}

func (s *Service) GetStudents() ([]*models.Student, error) {

	return s.db.GetStudents()
}

func (s *Service) UpdateStudentName(id int, body map[string]string) (int, error) {
	student, err := s.db.GetStudent(id)
	if err != nil {

		return 0, err
	}
	student.Name = body["name"]

	return s.db.AddOrUpdateStudent(student)
}

func (s *Service) UpdateGPA(id int, body map[string]float64) (int, error) {
	student, err := s.db.GetStudent(id)
	if err != nil {

		return 0, err
	}
	student.GPA = body["gpa"]

	return s.db.AddOrUpdateStudent(student)
}

func (s *Service) AssignSubjectToStudent(id int, body map[string]interface{}) (int, error) {
	student, err := s.db.GetStudent(id)
	if err != nil {

		return 0, err
	}

	var subjectID int
	sDataBytes, _ := json.Marshal(body["subjectID"])
	_ = json.Unmarshal(sDataBytes, &subjectID)

	var assign bool
	dataBytes, _ := json.Marshal(body["assign"])
	_ = json.Unmarshal(dataBytes, &assign)
	if assign {
		added := false
		for _, sid := range student.Subjects {
			if sid == subjectID {
				return 0, errors.Errorf("Assign Subject: Subject already assigned to the Student!")
			}
		}

		subject, err := s.db.GetSubject(subjectID)
		if err != nil {
			return 0, nil
		}

		class, err := s.db.GetClass(subject.ClassID)
		if err != nil {
			return 0, nil
		}

		for _, studentID := range class.Students {
			if studentID == id {
				student.Subjects = append(student.Subjects, subjectID)
				added = true
				break
			}
		}

		if added == false {
			return 0, errors.Errorf("Student must be already enrolled in the class, where the subject is being taught!")
		}

	} else {
		removed := false
		for i, sid := range student.Subjects {
			if sid == subjectID {
				student.Subjects = append(student.Subjects[:i], student.Subjects[i+1:]...)
				removed = true
				break
			}
		}

		if removed == false {
			return 0, errors.Errorf("Student not enrolled in this subject")
		}
	}

	return s.db.AddOrUpdateStudent(student)
}

func (s *Service) GetClassesOfStudent(id int) ([]*models.Class, error) {
	classes, err := s.db.GetClasses()
	if err != nil {

		return nil, err
	}
	var stdClasses []*models.Class
	for _, class := range classes {
		for _, studentID := range class.Students {
			if studentID == id {
				stdClasses = append(stdClasses, class)
			}
		}
	}

	return stdClasses, nil
}

func (s *Service) GetSubjectsOfStudent(id int) ([]*models.Subject, error) {
	student, err := s.db.GetStudent(id)
	if err != nil {

		return nil, err
	}

	if len(student.Subjects) == 0 {
		return nil, errors.Errorf("Student is not enrolled in any subject!")
	}
	var subjects []*models.Subject
	for _, subjectID := range student.Subjects {
		subject, err := s.db.GetSubject(subjectID)
		if err != nil {

			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (s *Service) DeleteStudent(id int) (string, error) {
	message, err := s.db.DeleteStudent(id)
	if err != nil {
		return "", err
	}

	classes, err := s.db.GetClasses()
	if err != nil {
		return "", err
	}

	for _, class := range classes {
		for i, studentID := range class.Students {
			if studentID == id {
				class.Students = append(class.Students[:i], class.Students[i+1:]...)
				_, err := s.db.AddOrUpdateClass(class)
				if err != nil {

					return "", errors.Wrap(err, "Delete Student: Student not removed from the class!")
				}

				break
			}

		}
	}

	accounts, err := s.db.GetAccounts()
	if err != nil {
		return "", err
	}

	for _, account := range accounts {
		if account.AccountType == data.STUDENT {
			var accountData *models.StudentAccount
			dBytes, _ := json.Marshal(account.AccountData)
			_ = json.Unmarshal(dBytes, &accountData)

			if accountData.StudentID == id {
				accountData.StudentID = 0
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

func (s *Service) GetStudentByID(id int) (*models.Student, error) {

	return s.db.GetStudent(id)
}
