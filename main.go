package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"gin-Home-server/controllers"
	"gin-Home-server/db"
	"gin-Home-server/forms"

	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"
	uuid "github.com/twinj/uuid"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

//RequestIDMiddleware ...
//Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

var auth = new(controllers.AuthController)

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request
//that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the default gin server
	r := gin.Default()

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis(1)

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)

		v1.GET("/user/getUserByID/:id", user.SearchUserByID)
		v1.GET("/user/getFamilyByID/:id", user.SearchFamilyById)
		v1.GET("/user/getFamilyByName/:name", user.SearchFamilyByName)
		v1.POST("/user/sendRequestToJoinFamily", user.AddFamilyUser)
		v1.GET("/user/getFamilyMembers/:familyID", user.GetFamilyMembersByFamilyID)
		/*** ADMIN***/
		v1.POST("/admin/createFamily", user.CreateFamily)
		v1.GET("/admin/receiveNewJoinFamilyReuqest/:id", user.ReceiveNewAdmin)
		v1.POST("/admin/addNewUserToFamily", user.ConfirmNewFamilyAdmin)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)

		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/***START Plan***/
		plan := new(controllers.PlanController)
		v1.POST("/plan/addPersonPlan", plan.AddPersonalPlan)
		v1.POST("/plan/addFamilyPlan", plan.AddFamilyPlan)

		v1.POST("/plan/getPersonPlanListByPersonID", plan.GetPersonalPlanListByPersonID)
		v1.POST("/plan/getFamilyPlanListByFamilyID", plan.GetFamilyPlanListByFamilyID)

		v1.POST("/plan/getPersonalPlanByPlanID", plan.GetPersonalPlanByPlanID)
		v1.POST("/plan/getFamilyPlanByPlanID", plan.GetFamilyPlanByPlanID)

		v1.DELETE("/plan/RemovePersonalPlanByPlanID", plan.RemovePersonPlanByPlanID)
		v1.DELETE("/plan/RemoveFamilyPlanByPlanID", plan.RemoveFamilyPlanByPlanID)

		v1.PUT("/plan/updatePersonalPlan/:id", plan.UpdatePersonPlanByPlanID)
		v1.PUT("/plan/updateFamilylPlan/:id", plan.UpdateFamilyPlanByPlanID)

		/***START Activity***/
		activity := new(controllers.ActivityController)
		v1.POST("/activity/addFamilyActivity", activity.CreateFamilyActivity)
		v1.POST("/activity/addPersonalActivity", activity.CreateFamilyActivity)

		v1.PUT("/activity/updateActivity/:id", activity.UpdateActivity)
		v1.POST("/activity/getFamilyActivityByDateID", activity.GetActivityByFamilyDate)
		v1.POST("/activity/getPersonActivityByDateID", activity.GetActivityByPersonDate)

		v1.GET("/activity/getActivityByID/:id", activity.GetActivityByID) //bug
		v1.DELETE("/activity/removeActivity", activity.RemoveActivityByPlanID)
		v1.POST("/activity/setFinish", activity.SetFinish)
		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", TokenAuthMiddleware(), article.Create)
		v1.GET("/articles", TokenAuthMiddleware(), article.All)
		v1.GET("/article/:id", TokenAuthMiddleware(), article.One)
		v1.PUT("/article/:id", TokenAuthMiddleware(), article.Update)
		v1.DELETE("/article/:id", TokenAuthMiddleware(), article.Delete)

	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}

}