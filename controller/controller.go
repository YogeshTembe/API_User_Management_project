package controller

import (
	"encoding/json"
	"net/http"

	"github.com/YogeshTembe/golang_project/file_handler"
	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	"github.com/YogeshTembe/golang_project/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser     = "root"
	dbPassword = "Panda@19"
	dbName     = "users"
	dbHost     = "127.0.0.1:3306"
)

func CreateUsers(store *store.Store, users []model.User) {
	for _, user := range users {
		if store.Db.Model(&user).Where("id = ?", user.Id).Updates(&user).RowsAffected == 0 {
			store.Db.Create(&user)
		}
	}
	store.Logger.Info("Succesfully added users data in database.")
}

func ConnectToDB(logger *li.StandardLogger) *gorm.DB {
	var err error
	var db *gorm.DB
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database-" + err.Error())
	} else {
		logger.Info("connected to database")
	}
	return db
}

func PostFilepath(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var file model.File
		json.NewDecoder(r.Body).Decode(&file)
		csvFile := file_handler.OpenCSVFile(file.FilePath, store.Logger)
		defer csvFile.Close()
		users := file_handler.ReadCSVFile(csvFile, store.Logger)
		file_handler.WriteJSONFile("users.json", users, store.Logger)
		CreateUsers(store, users)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Succesfully inserted in DB.")
	}
}

func GetUsers(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []model.User
		store.Db.Find(&users)
		json.NewEncoder(w).Encode(users)
	}
}
