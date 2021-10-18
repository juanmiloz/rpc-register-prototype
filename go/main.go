package main

import (
	//"fmt"
	//"log"
	"github.com/gin-gonic/gin"
	"net/http"
	//"net"
)


func main(){
	router := gin.Default()
	router.LoadHTMLFiles("../login.html", "../recoverPassword.html", "../homePage.html", "../profilePage.html", "../competitionPage.html")
	router.Static("/css", "../css")
	router.Static("/img", "../img")
	router.Static("/js", "../js")
	router.Static("/icons", "../icons")
	router.GET("/", loadMainPage)
	router.POST("/register", registerNewUser)
	router.POST("/login", loginProcessHandler)
	router.GET("/userIndex", loadIndexPage)
	router.GET("/userProfile", loadProfilePage)
	router.GET("/userCompetitions", loadCompetitiosPage)
	router.POST("/logout", logoutHandler)

	router.Run("localhost:8080")
}

func loadMainPage(c *gin.Context) {
	if(checkUserLogged()) {
		c.HTML(http.StatusOK, "homePage.html", gin.H{
			"message": " ",
			"user_logged": user_logged,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func loadIndexPage(c *gin.Context) {
	if(checkUserLogged()) {
		c.HTML(http.StatusOK, "homePage.html", gin.H{
			"message": " ",
			"user_logged": user_logged,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func registerNewUser(c *gin.Context) {
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	country := c.PostForm("country")
	new_user := User{Username: username, Email: email, Firstname: firstname, Lastname: lastname, Password: password, Country: country}
	_, response := createNewUser(new_user)
	if(response == 1) {
		c.String(http.StatusOK, "1")
	} else if (response == 2){
		c.String(http.StatusOK, "2")
	} else if (response == 3){
		c.String(http.StatusOK, "3")
	}
}

func loginProcessHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	responseLogin := loginProcess(email, password)
	if(responseLogin == 1) {
		c.String(http.StatusOK, "1")
	} else if (responseLogin == 2){
		c.String(http.StatusOK, "2")
	} else if (responseLogin == 3){
		c.String(http.StatusOK, "3")
	}
}

func loadProfilePage(c *gin.Context) {
	if(checkUserLogged()) {
		c.HTML(http.StatusOK, "profilePage.html", gin.H{
			"message": " ",
			"user_logged": user_logged,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func loadCompetitiosPage(c *gin.Context) {
	if(checkUserLogged()) {
		c.HTML(http.StatusOK, "competitionPage.html", gin.H{
			"message": " ",
			"user_logged": user_logged,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func logoutHandler(c *gin.Context) {
	if(logoutUser() == 1) {
		c.String(http.StatusOK, "1")
	} else {
		c.String(http.StatusOK, "0")
	}
}