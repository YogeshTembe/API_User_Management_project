package main

import (
	"net/http"

	"github.com/YogeshTembe/golang_project/controller"
	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/store"
	"github.com/gorilla/mux"
)

func main() {

	var StandardLogger = li.NewLogger()

	db := controller.ConnectToDB(StandardLogger)
	userStore := store.NewStore(db, StandardLogger)
	userStore.Logger.Info("server started")

	router := mux.NewRouter()
	router.HandleFunc("/api", controller.PostFilepath(userStore)).Methods("POST")
	router.HandleFunc("/api/users", controller.GetUsers(userStore)).Methods("GET")

	StandardLogger.Fatal(http.ListenAndServe(":8000", router).Error())
}
