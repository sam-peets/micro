package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sam-peets/micro/user"
)

// POST /auth
func HandleAuth(c *gin.Context) {
	var payload UserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, err)
		return
	}

	sess, err := Auth(&payload)
	if err != nil {
		c.JSON(401, err)
		return
	}

	c.JSON(200, sess)
}

// POST /auth/new
func HandleAuthNew(c *gin.Context) {
	var payload UserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, err)
		return
	}

	user, err := NewUser(&payload)
	if err != nil {
		c.JSON(400, err)
		return
	}

	sess, err := NewSession(*user)
	if err != nil {
		c.JSON(401, err)
		return
	}
	c.JSON(200, sess)
}

type AuthSessionPayload struct {
	Sid string `json:"sid"`
}

func HandleAuthValidate(c *gin.Context) {
	var payload AuthSessionPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	fmt.Println(payload.Sid)
	sess, err := GetSessionBySid(payload.Sid)
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}

	u, err := sess.GetUser()
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, u)
}

type GetUserPayload struct {
	Uid uint32 `json:"uid"`
}

func HandleGetUser(c *gin.Context) {
	var payload GetUserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	u, err := user.GetUser(payload.Uid)
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, u)

}
