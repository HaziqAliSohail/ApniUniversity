package service

import (
	"ApniUniversity/models"
	"encoding/json"
	"github.com/pkg/errors"
)

func (s *Service) AddStudent(student *models.Student) (int, error) {
	students, err := s.db.GetStudents()
	if err != nil {

		return 0, err
	}

	student.ID = students[len(students)-1].ID + 1

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

	return s.db.DeleteStudent(id)
}

func (s *Service) GetStudentByID(id int) (*models.Student, error) {

	return s.db.GetStudent(id)
}
