package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/CherkashinAV/finance_app/initializers"
	"github.com/CherkashinAV/finance_app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	var body struct{
		Name string
		Surname string
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"Wrong body message!",
		})
		return
	}

	//hash the pass
	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"Password hash error!",
		})
		return
	}

	//creating the user	
	var user models.User
	var result *gorm.DB

	//trying to find someone with such email already registered
	result = initializers.DB.Where("email = ?", body.Email).First(&user);

	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"User with same email already registered!",
		})
		return
	}

	user = models.User{Name: body.Name, Surname: body.Surname, Email: body.Email, Password: string(hashPass), Available_funds: 0, TotalFunds: 0, Lang: "ru", Theme: "light"}
	result = initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"User creation failed!",
		})
		return
	}

	//respond with statusOK
	c.JSON(http.StatusOK, gin.H {
		"user_id":user.ID,
	})

	return
}

func Login(c *gin.Context) {
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"Wrong body message!",
		})
		return
	}

	var user models.User
	var result *gorm.DB

	result = initializers.DB.Where("email = ?", body.Email).First(&user);
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"No users with such email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"Passwords don't match!",
		})
		return
	}

	//gen jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_CODE")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":"Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

	return
}

func UpdateUserInfo(c *gin.Context) {
}


func CheckIsAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"auth":true})
}

func GetUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

