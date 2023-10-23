package mongo

import (
	"os"
	"reflect"
	"testing"
	"time"

	"ApniUniversity/db"
	"ApniUniversity/models"
)

func Test_client_AddOrUpdateTeacher(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		teacher *models.Teacher
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Teacher is successfully Added to Database!",
			args:    args{teacher: &models.Teacher{ID: 23, Name: "Haziq Ali", SubjectID: 2}},
			wantErr: false,
		},
		{
			name:    "Teacher is successfully Updated in Database!",
			args:    args{teacher: &models.Teacher{ID: 2, Name: "Haziq Ali", SubjectID: 2}},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			c, _ := NewClient(db.Options{})

			_, err := c.AddOrUpdateTeacher(testCase.args.teacher)
			if (err != nil) != testCase.wantErr {
				test.Errorf("AddOrUpdateTeacher() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func Test_client_GetTeachers(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	var teacherModel []*models.Teacher
	testCases := []struct {
		name    string
		want    []*models.Teacher
		wantErr bool
	}{
		{
			name:    "Teachers are successfully fetched from Database",
			want:    teacherModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetTeachers()
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetTeachers() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			test.Logf("Got %v", got)
		})
	}
}

func Test_client_GetTeacher(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	teacherModel := &models.Teacher{ID: 12, Name: "Haziq", SubjectID: 2}
	teacherID, _ := c.AddOrUpdateTeacher(teacherModel)

	type args struct {
		id int
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	teacherModel.CreatedAt = teacherModel.CreatedAt.In(utc).Round(time.Second)
	teacherModel.UpdatedAt = teacherModel.UpdatedAt.In(utc).Round(time.Second)

	testCases := []struct {
		name    string
		args    args
		want    *models.Teacher
		wantErr bool
	}{
		{
			name:    "Teacher is successfully fetched from Database",
			args:    args{id: teacherID},
			want:    teacherModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetTeacher(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetTeacher() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("GetTeacher() got = %v, want %v", got, testCase.want)
			}
		})
	}
}

func Test_client_DeleteTeacher(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})
	teacherModel := &models.Teacher{ID: 12, Name: "Hello", SubjectID: 2}
	teacherID, _ := c.AddOrUpdateTeacher(teacherModel)

	type args struct {
		id int
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Teacher is successfully Deleted from Database!",
			args:    args{id: teacherID},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if _, err := c.DeleteTeacher(testCase.args.id); (err != nil) != testCase.wantErr {
				test.Errorf("DeleteTeacher() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func Test_client_AddOrUpdateSubject(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		subject *models.Subject
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Subject is successfully Added to Database!",
			args:    args{subject: &models.Subject{ID: 19, Name: "OOP", ClassID: 2}},
			wantErr: false,
		},
		{
			name:    "Subject is successfully Updated in Database!",
			args:    args{subject: &models.Subject{ID: 1, Name: "OOP", ClassID: 2}},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			c, _ := NewClient(db.Options{})
			_, err := c.AddOrUpdateSubject(testCase.args.subject)
			if (err != nil) != testCase.wantErr {
				test.Errorf("AddSubject() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}
func Test_client_GetSubjects(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	var subjectModel []*models.Subject

	testCases := []struct {
		name    string
		want    []*models.Subject
		wantErr bool
	}{
		{
			name:    "Subjects are successfully fetched from Database",
			want:    subjectModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetSubjects()
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetSubjects() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			test.Logf("Got %v", got)
		})
	}
}
func Test_client_GetSubject(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	subjectModel := &models.Subject{ID: 12, Name: "OOP", ClassID: 2}
	subjectID, _ := c.AddOrUpdateSubject(subjectModel)

	type args struct {
		id int
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	subjectModel.CreatedAt = subjectModel.CreatedAt.In(utc).Round(time.Second)
	subjectModel.UpdatedAt = subjectModel.UpdatedAt.In(utc).Round(time.Second)

	testCases := []struct {
		name    string
		args    args
		want    *models.Subject
		wantErr bool
	}{
		{
			name:    "Subject is successfully Fetched from Database!",
			args:    args{id: subjectID},
			want:    subjectModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetSubject(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetSubject() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("GetSubject() got = %v, want %v", got, testCase.want)
			}
		})
	}
}

func Test_client_DeleteSubject(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})
	subjectModel := &models.Subject{ID: 12, Name: "OOP", ClassID: 2}
	subjectID, _ := c.AddOrUpdateSubject(subjectModel)

	type args struct {
		id int
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Subject is successfully deleted from Database!",
			args:    args{id: subjectID},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if _, err := c.DeleteSubject(testCase.args.id); (err != nil) != testCase.wantErr {
				test.Errorf("DeleteSubject() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}
