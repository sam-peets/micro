package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sam-peets/micro/auth"
	"github.com/sam-peets/micro/db"
	"github.com/sam-peets/micro/posts"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	client := db.GetClient()
	defer client.Close()
	context, cancel := client.Context()
	defer cancel()
	client.Client.Ping(context, readpref.Primary())

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/auth", auth.HandleAuth)
		api.POST("/auth/new", auth.HandleAuthNew)
		api.POST("/auth/validate", auth.HandleAuthValidate)

		api.POST("/posts", posts.HandleGetPost)
		api.POST("/posts/new", posts.HandleNewPost)
		api.POST("/posts/recent", posts.HandleRecentPosts)

		api.POST("/users", auth.HandleGetUser)
	}

	r.LoadHTMLGlob("template/*.html")

	bind := func(path, file string, obj any) {
		r.GET(path, func(c *gin.Context) {
			c.HTML(http.StatusOK, file, obj)
		})
	}

	bind("/", "index.html", gin.H{
		"title": "Micro",
	})

	bind("/login", "login.html", gin.H{
		"title": "Login",
	})

	bind("/signup", "signup.html", gin.H{
		"title": "Sign Up",
	})

	r.StaticFS("/static", http.Dir("static"))

	// user, err := auth.NewUser(&auth.UserPayload{
	// Username: "tester2",
	// Password: "asdf",
	// })
	// fmt.Println(user, err)

	r.Run()
}
