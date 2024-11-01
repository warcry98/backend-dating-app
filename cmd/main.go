package main

import (
	"backend-dating-app/api/profile"
	"backend-dating-app/api/swipe"
	"backend-dating-app/api/user"
	"backend-dating-app/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	migrateDB(dbConn)

	r := mux.NewRouter()

	userHandler := user.NewUserHandler(dbConn)
	profileHandler := profile.NewProfileHandler(dbConn)
	swipeHandler := swipe.NewSwipeHandler(dbConn)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	auth.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(user.AuthenticationMiddleware)

	protected.HandleFunc("/swipe", swipeHandler.RecordSwipe).Methods("POST")
	protected.HandleFunc("/profile", profileHandler.GetOwnProfile).Methods("GET")
	protected.HandleFunc("/profiles", profileHandler.GetProfiles).Methods("GET")

	create := protected.PathPrefix("/create").Subrouter()
	create.HandleFunc("/profile", profileHandler.CreateProfile).Methods("POST")

	update := protected.PathPrefix("/update").Subrouter()
	update.HandleFunc("/profile", profileHandler.UpdateProfile).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &profile.Profile{}, &swipe.Swipe{})
}
