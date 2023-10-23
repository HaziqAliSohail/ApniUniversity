package mongo

import (
	"ApniUniversity/config"
	"ApniUniversity/db"
	"ApniUniversity/models"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	teacherCollection = "Teacher"
	subjectCollection = "Subject"
)

type client struct {
	connection *mongo.Client
}

func NewClient(_ db.Options) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("Initializing MongoDB At: %s", uri)
	clientOptions := options.Client().ApplyURI(uri)
	cli, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected to mongodb!")

	return &client{connection: cli}, nil
}

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

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", teacher.ID}}, bson.D{{"$set", teacher}}, options.Update().SetUpsert(true))
	if err != nil {
		return 0, errors.Wrap(err, "Teacher not Added/Updated!")
	}

	return teacher.ID, nil
}

func (c *client) GetTeacher(id int) (*models.Teacher, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)
	var teacher *models.Teacher

	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&teacher)
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

func (c *client) GetTeachers() ([]*models.Teacher, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)
	var teachers []*models.Teacher

	teachersCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
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

func (c *client) DeleteTeacher(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(teacherCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Teacher not Found!")
		}

		return "", err
	}

	return fmt.Sprintf("Teacher Deleted Successfully!"), nil
}

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

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", subject.ID}}, bson.D{{"$set", subject}}, options.Update().SetUpsert(true))

	if err != nil {
		return 0, errors.Wrap(err, "Subject not added/updated!")
	}

	return subject.ID, nil
}

func (c *client) GetSubject(id int) (*models.Subject, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)
	var subject *models.Subject

	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&subject)
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

func (c *client) GetSubjects() ([]*models.Subject, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)
	var subjects []*models.Subject

	subjectsCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
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

func (c *client) DeleteSubject(id int) (string, error) {
	collection := c.connection.Database(viper.GetString(config.DbName)).Collection(subjectCollection)

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.Wrap(err, "Subject not Found!")
		}

		return "", err
	}

	return fmt.Sprintf("Subject Deleted Successfully!"), nil
}
