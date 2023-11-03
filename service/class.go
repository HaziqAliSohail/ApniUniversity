package service

import (
	"ApniUniversity/models"
	"encoding/json"
	"github.com/pkg/errors"
)

func (s *Service) AddClass(class *models.Class) (int, error) {
	classes, err := s.db.GetClasses()
	if err != nil {

		return 0, err
	}

	class.ID = classes[len(classes)-1].ID + 1

	return s.db.AddOrUpdateClass(class)
}

func (s *Service) GetClasses() ([]*models.Class, error) {

	return s.db.GetClasses()
}

func (s *Service) UpdateClassName(id int, body map[string]string) (int, error) {
	class, err := s.db.GetClass(id)
	if err != nil {

		return 0, err
	}
	class.Name = body["name"]

	return s.db.AddOrUpdateClass(class)
}

func (s *Service) AddOrRemoveStudent(id int, body map[string]interface{}) (int, error) {
	class, err := s.db.GetClass(id)
	if err != nil {

		return 0, err
	}

	var studentID int
	sDataBytes, _ := json.Marshal(body["studentID"])
	_ = json.Unmarshal(sDataBytes, &studentID)

	var add bool
	dataBytes, _ := json.Marshal(body["add"])
	_ = json.Unmarshal(dataBytes, &add)
	if add {
		for _, sid := range class.Students {
			if sid == studentID {
				return 0, errors.Errorf("Add Student: Student Already enrolled in the class!")
			}
		}

		class.Students = append(class.Students, studentID)
	} else {
		removed := false
		for i, sid := range class.Students {
			if sid == studentID {
				student, err := s.db.GetStudent(studentID)
				if err != nil {
					return 0, err
				}

				var updatedSubjects []int
				for _, subjectID := range student.Subjects {
					subject, err := s.db.GetSubject(subjectID)
					if err != nil {
						return 0, err
					}

					if subject.ClassID != id {
						updatedSubjects = append(updatedSubjects, subject.ID)
					}

				}

				class.Students = append(class.Students[:i], class.Students[i+1:]...)
				removed = true
			}
		}

		if removed == false {
			return 0, errors.Errorf("Remove Student: Student not enrolled in this class!")
		}
	}

	classID, err := s.db.AddOrUpdateClass(class)
	if err != nil {
		if add {

			return 0, errors.Wrap(err, "Add Student: Student not enrolled in the class!")
		} else {

			return 0, errors.Wrap(err, "Remove Student: Student not removed from the class!")
		}
	}

	return classID, nil
}

func (s *Service) GetSubjectsOfClass(id int) ([]*models.Subject, error) {
	subjects, err := s.db.GetSubjects()
	if err != nil {

		return nil, err
	}

	var classSubjects []*models.Subject
	for _, subject := range subjects {
		if subject.ClassID == id {
			classSubjects = append(classSubjects, subject)
		}
	}

	if len(subjects) == 0 {
		return nil, errors.Errorf("No subjects registered in this class!")
	}

	return subjects, nil
}

func (s *Service) GetTeachersOfClass(id int) ([]*models.Teacher, error) {
	teachers, err := s.db.GetTeachers()
	if err != nil {

		return nil, err
	}

	subjects, err := s.db.GetSubjects()
	if err != nil {

		return nil, err
	}

	var classTeachers []*models.Teacher
	for _, subject := range subjects {
		if subject.ClassID == id {
			for _, teacher := range teachers {
				if teacher.SubjectID == subject.ID {
					classTeachers = append(classTeachers, teacher)
				}
			}
		}
	}

	return classTeachers, nil
}

func (s *Service) GetStudentsOfClass(id int) ([]*models.Student, error) {
	class, err := s.db.GetClass(id)
	if err != nil {

		return nil, err
	}
	var students []*models.Student
	for _, studentID := range class.Students {
		student, err := s.db.GetStudent(studentID)
		if err != nil {

			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *Service) DeleteClass(id int) (string, error) {

	return s.db.DeleteClass(id)
}

func (s *Service) GetClassByID(id int) (*models.Class, error) {

	return s.db.GetClass(id)
}
