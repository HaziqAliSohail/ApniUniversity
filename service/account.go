package service

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/HaziqAliSohail/ApniUniversity/data"
	"github.com/HaziqAliSohail/ApniUniversity/models"
)

func (s *Service) AddAccount(account *models.Account) (int, error) {
	accounts, err := s.db.GetAccounts()
	if err != nil {

		return 0, err
	}

	if len(accounts) != 0 {
		account.ID = accounts[len(accounts)-1].ID + 1
	} else {
		account.ID = 1
	}

	if account.AccountType == data.TEACHER {
		var tData *models.TeacherAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &tData)

		if tData.TeacherID != 0 {
			if _, err = s.db.GetTeacher(tData.TeacherID); err != nil {

				return 0, err
			}
		}

	} else if account.AccountType == data.STUDENT {
		var sData *models.StudentAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &sData)
		account.AccountData = sData

		if sData.StudentID != 0 {
			if _, err = s.db.GetStudent(sData.StudentID); err != nil {

				return 0, err
			}
		}

	}

	return s.db.AddOrUpdateAccount(account)
}

func (s *Service) GetAccounts() ([]*models.Account, error) {

	return s.db.GetAccounts()
}

func (s *Service) UpdateAccount(account *models.Account) (int, error) {
	_, err := s.db.GetAccount(account.ID)
	if err != nil {

		return 0, err
	}

	if account.AccountType == data.TEACHER {
		var tData *models.TeacherAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &tData)
		account.AccountData = tData

		if tData.TeacherID != 0 {
			if _, err = s.db.GetTeacher(tData.TeacherID); err != nil {

				return 0, err
			}
		}

	} else if account.AccountType == data.STUDENT {
		var sData *models.StudentAccount
		dataBytes, _ := bson.Marshal(account.AccountData)
		_ = bson.Unmarshal(dataBytes, &sData)
		account.AccountData = sData

		if sData.StudentID != 0 {
			if _, err = s.db.GetStudent(sData.StudentID); err != nil {

				return 0, err
			}
		}

	}

	return s.db.AddOrUpdateAccount(account)
}

func (s *Service) GetStudentAccounts() ([]*models.Account, error) {
	accounts, _ := s.db.GetAccounts()

	var stdAccounts []*models.Account
	for _, account := range accounts {
		if account.AccountType == data.STUDENT {
			stdAccounts = append(stdAccounts, account)
		}
	}

	if len(stdAccounts) == 0 {
		return stdAccounts, errors.Errorf("No Student Account Found!")
	}

	return stdAccounts, nil
}

func (s *Service) GetDefaultedStudentAccounts() ([]*models.Account, error) {
	accounts, _ := s.db.GetAccounts()

	var stdAccounts []*models.Account
	for _, account := range accounts {
		if account.AccountType == data.STUDENT {
			var accountData *models.StudentAccount
			dBytes, _ := json.Marshal(account.AccountData)
			_ = json.Unmarshal(dBytes, &accountData)

			account.AccountData = accountData
			if accountData.IsDefaulter {
				stdAccounts = append(stdAccounts, account)
			}
		}
	}

	if len(stdAccounts) == 0 {
		return stdAccounts, errors.Errorf("No Defaulted Student Account Found!")
	}

	return stdAccounts, nil
}

func (s *Service) GetTeacherAccounts() ([]*models.Account, error) {
	accounts, _ := s.db.GetAccounts()

	var teacherAccounts []*models.Account
	for _, account := range accounts {
		if account.AccountType == data.TEACHER {
			teacherAccounts = append(teacherAccounts, account)
		}
	}

	if len(teacherAccounts) == 0 {
		return teacherAccounts, errors.Errorf("No Teacher Account Found!")
	}

	return teacherAccounts, nil
}

func (s *Service) DeleteAccount(id int) (string, error) {

	return s.db.DeleteAccount(id)
}

func (s *Service) GetAccountByID(id int) (*models.Account, error) {

	return s.db.GetAccount(id)
}
