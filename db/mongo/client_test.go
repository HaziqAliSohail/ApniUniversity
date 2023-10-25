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

func TestClient_AddOrUpdateClass(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		class *models.Class
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Class is successfully Added to Database!",
			args:    args{class: &models.Class{ID: 23, Name: "BSE-3", Students: []int{2, 4, 5}}},
			wantErr: false,
		},
		{
			name:    "Class is successfully Updated in Database!",
			args:    args{class: &models.Class{ID: 2, Name: "BSE-8", Students: []int{1, 2, 4}}},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			c, _ := NewClient(db.Options{})

			_, err := c.AddOrUpdateClass(testCase.args.class)
			if (err != nil) != testCase.wantErr {
				test.Errorf("AddOrUpdateClass() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestClient_GetClasses(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	var classModel []*models.Class
	testCases := []struct {
		name    string
		want    []*models.Class
		wantErr bool
	}{
		{
			name:    "Classes are successfully fetched from Database",
			want:    classModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetClasses()
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetClasses() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			test.Logf("Got %v", got)
		})
	}
}

func TestClient_GetClass(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	classModel := &models.Class{ID: 12, Name: "BEE-2", Students: []int{1, 2, 3}}
	classID, _ := c.AddOrUpdateClass(classModel)

	type args struct {
		id int
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	classModel.CreatedAt = classModel.CreatedAt.In(utc).Round(time.Second)
	classModel.UpdatedAt = classModel.UpdatedAt.In(utc).Round(time.Second)

	testCases := []struct {
		name    string
		args    args
		want    *models.Class
		wantErr bool
	}{
		{
			name:    "Class is successfully fetched from Database",
			args:    args{id: classID},
			want:    classModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetClass(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetClass() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("GetClass() got = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestClient_DeleteClass(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})
	classModel := &models.Class{ID: 23, Name: "BCS-9", Students: []int{1, 3, 2}}
	classID, _ := c.AddOrUpdateClass(classModel)

	type args struct {
		id int
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Class is successfully Deleted from Database!",
			args:    args{id: classID},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if _, err := c.DeleteClass(testCase.args.id); (err != nil) != testCase.wantErr {
				test.Errorf("DeleteClass() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestClient_AddOrUpdateStudent(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		student *models.Student
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Student is successfully Added to Database!",
			args:    args{student: &models.Student{ID: 22, Name: "Haziq Ali", Subjects: []int{1, 2}}},
			wantErr: false,
		},
		{
			name:    "Student is successfully Updated in Database!",
			args:    args{student: &models.Student{ID: 22, Name: "Haziq Ali", Subjects: []int{1, 2, 3}}},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			c, _ := NewClient(db.Options{})

			_, err := c.AddOrUpdateStudent(testCase.args.student)
			if (err != nil) != testCase.wantErr {
				test.Errorf("AddOrUpdateStudent() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestClient_GetStudents(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	var studentModel []*models.Student
	testCases := []struct {
		name    string
		want    []*models.Student
		wantErr bool
	}{
		{
			name:    "Students are successfully fetched from Database",
			want:    studentModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetStudents()
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetStudents() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			test.Logf("Got %v", got)
		})
	}
}

func TestClient_GetStudent(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	studentModel := &models.Student{ID: 12, Name: "Haziq", Subjects: []int{1, 2, 4}}
	studentID, _ := c.AddOrUpdateStudent(studentModel)

	type args struct {
		id int
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	studentModel.CreatedAt = studentModel.CreatedAt.In(utc).Round(time.Second)
	studentModel.UpdatedAt = studentModel.UpdatedAt.In(utc).Round(time.Second)

	testCases := []struct {
		name    string
		args    args
		want    *models.Student
		wantErr bool
	}{
		{
			name:    "Student is successfully fetched from Database",
			args:    args{id: studentID},
			want:    studentModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetStudent(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetStudent() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("GetStudent() got = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestClient_DeleteStudent(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})
	studentModel := &models.Student{ID: 12, Name: "Hello", Subjects: []int{1, 2}}
	studentID, _ := c.AddOrUpdateStudent(studentModel)

	type args struct {
		id int
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Student is successfully Deleted from Database!",
			args:    args{id: studentID},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if _, err := c.DeleteStudent(testCase.args.id); (err != nil) != testCase.wantErr {
				test.Errorf("DeleteStudent() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestClient_AddOrUpdateAccount(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		account *models.Account
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Account is successfully Added to Database!",
			args:    args{account: &models.Account{ID: 12, Name: "Haziq Ali", AccountType: "student", AccountData: &models.StudentAccount{StudentID: 1, TotalInstallments: 1, RemainingInstallments: 1}}},
			wantErr: false,
		},
		{
			name:    "Account is successfully Updated in Database!",
			args:    args{account: &models.Account{ID: 12, Name: "Haziq Ali Sohail", AccountType: "student", AccountData: &models.StudentAccount{StudentID: 1, TotalInstallments: 1, RemainingInstallments: 1}}},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			c, _ := NewClient(db.Options{})

			_, err := c.AddOrUpdateAccount(testCase.args.account)
			if (err != nil) != testCase.wantErr {
				test.Errorf("AddOrUpdateAccount() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestClient_GetAccounts(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	var accountModel []*models.Account
	testCases := []struct {
		name    string
		want    []*models.Account
		wantErr bool
	}{
		{
			name:    "Accounts are successfully fetched from Database",
			want:    accountModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetAccounts()
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetAccounts() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			for _, get := range got {
				test.Logf("Got %v", get.AccountData)
			}
		})
	}
}

func TestClient_GetAccount(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})

	accountModel := &models.Account{ID: 23, Name: "Haziq ", AccountType: "teacher", AccountData: &models.TeacherAccount{TeacherID: 1, Salary: 1000}}
	accountID, _ := c.AddOrUpdateAccount(accountModel)

	type args struct {
		id int
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	accountModel.CreatedAt = accountModel.CreatedAt.In(utc).Round(time.Second)
	accountModel.UpdatedAt = accountModel.UpdatedAt.In(utc).Round(time.Second)

	testCases := []struct {
		name    string
		args    args
		want    *models.Account
		wantErr bool
	}{
		{
			name:    "Account is successfully fetched from Database",
			args:    args{id: accountID},
			want:    accountModel,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			got, err := c.GetAccount(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				test.Errorf("GetAccount() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("GetAccount() got = %v, want %v", got, testCase.want)
				test.Errorf("Issue Got Account Data = (%v,%T) \n Want Account Data (%v,%T)", got.AccountData, got.AccountData, testCase.want.AccountData, testCase.want.AccountData)

			}
		})
	}
}

func TestClient_DeleteAccount(test *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Options{})
	accountModel := &models.Account{ID: 23, Name: "Haziq Ali", AccountType: "teacher", AccountData: &models.TeacherAccount{TeacherID: 1, Salary: 1000}}
	accountID, _ := c.AddOrUpdateAccount(accountModel)

	type args struct {
		id int
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Account is successfully Deleted from Database!",
			args:    args{id: accountID},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if _, err := c.DeleteAccount(testCase.args.id); (err != nil) != testCase.wantErr {
				test.Errorf("DeleteAccount() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}
