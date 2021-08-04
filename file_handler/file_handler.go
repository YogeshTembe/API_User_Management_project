package file_handler

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	"github.com/YogeshTembe/golang_project/validation"
	uuid "github.com/satori/go.uuid"
)

func OpenCSVFile(fileDir string, logger *li.StandardLogger) *os.File {
	csvFile, err := os.Open(fileDir)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Successfully Opened CSV file")
	return csvFile
}

func ReadCSVFile(csvFile *os.File, logger *li.StandardLogger) []model.User {
	var users []model.User

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		logger.Fatal("Unable to read the CSV file-" + err.Error())
	}

	for _, line := range csvLines {
		id, _ := uuid.FromString(line[0])
		phoneNo, _ := strconv.Atoi(line[3])
		isActive, _ := strconv.ParseBool(line[4])

		user := model.User{
			Id:          id,
			Name:        line[1],
			Email:       line[2],
			PhoneNumber: phoneNo,
			IsActive:    isActive,
		}
		isValid := validation.Validate(&user, logger)

		if isValid {
			users = append(users, user)
			validation.UserIds[user.Id.String()] = struct{}{}
		}
	}
	logger.Info("CSV file reading and data validation is done.")
	return users
}

func WriteJSONFile(fileDir string, users []model.User, logger *li.StandardLogger) {
	file, _ := json.MarshalIndent(users, "", " ")
	err := ioutil.WriteFile(fileDir, file, 0644)
	if err != nil {
		logger.Fatal("Unable to write to JSON file-" + err.Error())
	}
	logger.Info("JSON file writing is done.")
}
