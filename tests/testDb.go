package main

import (
	"fmt"
	"log"
	"os"

	// "gin-Home-server/db"
	// "gin-Home-server/forms"
	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
)

func main() {

	fmt.Println("TestDb")
	// //Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	fmt.Println("DB_PASS", os.Getenv("DB_PASS"))

	// getDb := db.GetDB()

	// var email string = "Runoob1"
	// getDb.Query("INSERT INTO public.user(email, password, name) VALUES($1, $2, $3) RETURNING id", email, email, email)

	// var registerForm forms.RegisterForm

	// registerForm.Name = "testing"
	// registerForm.Email = "testEmail"
	// registerForm.Password = "testPassword"

	// data, _ := json.Marshal(registerForm)
	// fmt.Println("-------data rigister", data)
	// req, err := http.NewRequest("POST", "http://localhost:9000/v1/user/register", bytes.NewBufferString(string(data)))
	// req.Header.Set("Content-Type", "application/json")

	// fmt.Println("------err",err)

	// fmt.Println("---req",req)

	// clt := http.Client{}
	// clt.Do(req)

	// resp := httptest.NewRecorder()

	// testRouter.ServeHTTP(resp, req)

}
