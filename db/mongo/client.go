package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/HaziqAliSohail/ApniUniversity/config"
	"github.com/HaziqAliSohail/ApniUniversity/data"
	"github.com/HaziqAliSohail/ApniUniversity/db"
	"github.com/HaziqAliSohail/ApniUniversity/models"
)

const (
	teacherCollection = "Teacher"
	subjectCollection = "Subject"
	classCollection   = "Class"
	studentCollection = "Student"
	accountCollection = "Account"
)

// structure containing mongodb client
type client struct {
	connection *mongo.Client
}

var (
	mongoClient *client
	mongoOnce   sync.Once
)

// NewClient Initializing the Database Client using constructor method
func NewClient(_ db.Options) (db.DataStore, error) {
	mongoOnce.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
		log().Infof("Initializing MongoDB At: %s", uri)
		clientOptions := options.Client().ApplyURI(uri)
		cli, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully Connected to mongodb!")
		mongoClient = &client{connection: cli}
	})

	return mongoClient, nil
}

// AddOrUpdateTeacher Adding a teacher to database or updating the already present teacher
func (c *client) AddOrUpdateTeacher(teacher *models.Teacher) (int, error) {
	if teacher == nil {
		return 0, errors.Errorf("Data is empty!")
	}

	update := false
	if tData, err := c.GetTeacher(teacher.ID); err == nil {
		update = true
		teacher.CreatedAt = tData.CreatedAt
	}

	if update {
		teacher.UpdatedAt = time.Now()
	} else {
		teacher.CreatedAt = time.Now()
		teacher.UpdatedAt = time.Now()
	}

	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)

	_, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: teacher.ID}}, bson.D{primitive.E{Key: "$set", Value: teacher}}, options.Update().SetUpsert(true))
	if err != nil {
		return 0, errors.Wrap(err, "Teacher not Added/Updated!")
	}

	return teacher.ID, nil
}

// GetTeacher Fetching a teacher's data from database based on ID
func (c *client) GetTeacher(id int) (*models.Teacher, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)
	var teacher *models.Teacher

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&teacher)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "Teacher not Found!")
		}

		return nil, err
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	teacher.CreatedAt = teacher.CreatedAt.In(utc).Round(time.Second)
	teacher.UpdatedAt = teacher.UpdatedAt.In(utc).Round(time.Second)

	return teacher, nil
}

// GetTeachers Fetching the data of all teachers present in the database
func (c *client) GetTeachers() ([]*models.Teacher, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)
	var teachers []*models.Teacher

	teachersCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "No Teachers' Data Found!")
		}

		return nil, err
	}

	if err := teachersCursor.All(context.TODO(), &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

// DeleteTeacher Deleting a teacher from database based on ID
func (c *client) DeleteTeacher(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Teacher not Found!")
		}

		return "", err
	}

	return "Teacher Deleted Successfully!", nil
}

// AddOrUpdateSubject Adding a subject to database or updating the already present subject
func (c *client) AddOrUpdateSubject(subject *models.Subject) (int, error) {
	if subject == nil {
		return 0, errors.Errorf("Data is Empty!")
	}

	update := false
	if sData, err := c.GetTeacher(subject.ID); err == nil {
		update = true
		subject.CreatedAt = sData.CreatedAt
	}

	if update {
		subject.UpdatedAt = time.Now()
	} else {
		subject.CreatedAt = time.Now()
		subject.UpdatedAt = time.Now()
	}

	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)

	_, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: subject.ID}}, bson.D{primitive.E{Key: "$set", Value: subject}}, options.Update().SetUpsert(true))

	if err != nil {
		return 0, errors.Wrap(err, "Subject not added/updated!")
	}

	return subject.ID, nil
}

// GetSubject Fetching a subject's data from database based on ID
func (c *client) GetSubject(id int) (*models.Subject, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)
	var subject *models.Subject

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&subject)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "Subject not Found!")
		}

		return nil, err
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	subject.CreatedAt = subject.CreatedAt.In(utc).Round(time.Second)
	subject.UpdatedAt = subject.UpdatedAt.In(utc).Round(time.Second)

	return subject, nil
}

// GetSubjects Fetching the data of all subjects present in the database
func (c *client) GetSubjects() ([]*models.Subject, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)
	var subjects []*models.Subject

	subjectsCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "No Subjects' Data Found!")
		}

		return nil, err
	}

	if err := subjectsCursor.All(context.TODO(), &subjects); err != nil {
		return nil, err
	}

	return subjects, nil
}

// DeleteSubject Deleting a subject from database based on ID
func (c *client) DeleteSubject(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Subject not Found!")
		}

		return "", err
	}

	return "Subject Deleted Successfully!", nil
}

// AddOrUpdateClass Adding a class to database or updating the already present class
func (c *client) AddOrUpdateClass(class *models.Class) (int, error) {
	if class == nil {
		return 0, errors.Errorf("Data is empty!")
	}

	update := false
	if cData, err := c.GetClass(class.ID); err == nil {
		update = true
		class.CreatedAt = cData.CreatedAt
	}

	if update {
		class.UpdatedAt = time.Now()
	} else {
		class.CreatedAt = time.Now()
		class.UpdatedAt = time.Now()
	}

	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(classCollection)

	_, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: class.ID}}, bson.D{primitive.E{Key: "$set", Value: class}}, options.Update().SetUpsert(true))
	if err != nil {
		return 0, errors.Wrap(err, "Class not Added/Updated!")
	}

	return class.ID, nil
}

// GetClass Fetching a class' data from database based on ID
func (c *client) GetClass(id int) (*models.Class, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(classCollection)
	var class *models.Class

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&class)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "Class not Found!")
		}

		return nil, err
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	class.CreatedAt = class.CreatedAt.In(utc).Round(time.Second)
	class.UpdatedAt = class.UpdatedAt.In(utc).Round(time.Second)

	return class, nil
}

// GetClasses Fetching the data of all classes present in the database
func (c *client) GetClasses() ([]*models.Class, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(classCollection)
	var classes []*models.Class

	classesCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "No Classes' Data Found!")
		}

		return nil, err
	}

	if err := classesCursor.All(context.TODO(), &classes); err != nil {
		return nil, err
	}

	return classes, nil
}

// DeleteClass Deleting a class from database based on ID
func (c *client) DeleteClass(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(classCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Class not Found!")
		}

		return "", err
	}

	return "Class Deleted Successfully!", nil
}

// AddOrUpdateStudent Adding a student to database or updating the already present student
func (c *client) AddOrUpdateStudent(student *models.Student) (int, error) {
	if student == nil {
		return 0, errors.Errorf("Data is empty!")
	}

	update := false
	if sData, err := c.GetStudent(student.ID); err == nil {
		update = true
		student.CreatedAt = sData.CreatedAt
	}

	if update {
		student.UpdatedAt = time.Now()
	} else {
		student.CreatedAt = time.Now()
		student.UpdatedAt = time.Now()
	}

	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(studentCollection)

	_, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: student.ID}}, bson.D{primitive.E{Key: "$set", Value: student}}, options.Update().SetUpsert(true))
	if err != nil {
		return 0, errors.Wrap(err, "Student not Added/Updated!")
	}

	return student.ID, nil
}

// GetStudent Fetching a student's data from database based on ID
func (c *client) GetStudent(id int) (*models.Student, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(studentCollection)
	var student *models.Student

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "Student not Found!")
		}

		return nil, err
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	student.CreatedAt = student.CreatedAt.In(utc).Round(time.Second)
	student.UpdatedAt = student.UpdatedAt.In(utc).Round(time.Second)

	return student, nil
}

// GetStudents Fetching the data of all students present in the database
func (c *client) GetStudents() ([]*models.Student, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(studentCollection)
	var students []*models.Student

	studentsCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "No Students' Data Found!")
		}

		return nil, err
	}

	if err := studentsCursor.All(context.TODO(), &students); err != nil {

		return nil, err
	}

	return students, nil
}

// DeleteStudent Deleting a student from database based on ID
func (c *client) DeleteStudent(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(studentCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Student not Found!")
		}

		return "", err
	}

	return "Student Deleted Successfully!", nil
}

// AddOrUpdateAccount Adding an account to database or updating the already present account
func (c *client) AddOrUpdateAccount(account *models.Account) (int, error) {
	if account == nil {
		return 0, errors.Errorf("Data is empty!")
	}

	update := false
	if aData, err := c.GetAccount(account.ID); err == nil {
		update = true
		account.CreatedAt = aData.CreatedAt
	}

	if update {
		account.UpdatedAt = time.Now()
	} else {
		account.CreatedAt = time.Now()
		account.UpdatedAt = time.Now()
	}

	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(accountCollection)

	_, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: account.ID}}, bson.D{primitive.E{Key: "$set", Value: account}}, options.Update().SetUpsert(true))
	if err != nil {
		return 0, errors.Wrap(err, "Account not Added/Updated!")
	}

	return account.ID, nil
}

// GetAccount Fetching an account's data from database based on ID
func (c *client) GetAccount(id int) (*models.Account, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(accountCollection)
	var account *models.Account

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "Account not Found!")
		}

		return nil, err
	}

	utc, _ := time.LoadLocation("Asia/Karachi")
	account.CreatedAt = account.CreatedAt.In(utc).Round(time.Second)
	account.UpdatedAt = account.UpdatedAt.In(utc).Round(time.Second)

	if account.AccountType == data.TEACHER {
		var tData *models.TeacherAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &tData)
		account.AccountData = tData
	} else if account.AccountType == data.STUDENT {
		var sData *models.StudentAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &sData)
		account.AccountData = sData
	}

	return account, nil
}

// GetAccounts Fetching the data of all accounts present in the database
func (c *client) GetAccounts() ([]*models.Account, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(accountCollection)
	var accounts []*models.Account

	accountsCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "No Accounts' Data Found!")
		}

		return nil, err
	}

	if err := accountsCursor.All(context.TODO(), &accounts); err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.AccountType == data.TEACHER {
			var tData *models.TeacherAccount
			dataBytes, _ := bson.Marshal(account.AccountData)
			_ = bson.Unmarshal(dataBytes, &tData)
			account.AccountData = tData
		} else if account.AccountType == data.STUDENT {
			var sData *models.StudentAccount
			dataBytes, _ := bson.Marshal(account.AccountData)
			_ = bson.Unmarshal(dataBytes, &sData)
			account.AccountData = sData
		}
	}

	return accounts, nil
}

// DeleteAccount Deleting an account from database based on ID
func (c *client) DeleteAccount(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(accountCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Account not Found!")
		}

		return "", err
	}

	return "Account Deleted Successfully!", nil
}
