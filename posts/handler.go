package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/sam-peets/micro/auth"
)

type GetPostPayload struct {
	Sid string `json:"sid"`
	Pid uint32 `json:"pid"`
}

func HandleGetPost(c *gin.Context) {
	var payload GetPostPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	_, err = auth.GetSessionBySid(payload.Sid)
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}

	p, err := GetPost(payload.Pid)
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, p)
}

type NewPostPayload struct {
	Sid     string `json:"sid"`
	Content string `json:"content"`
}

func HandleNewPost(c *gin.Context) {
	var payload NewPostPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	sess, err := auth.GetSessionBySid(payload.Sid)
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}

	u, err := sess.GetUser()
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}

	p, err := NewPost(*u, payload.Content)
	if err != nil {
		c.JSON(401, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, p)
}

type RecentPostsPayload struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

func HandleRecentPosts(c *gin.Context) {
	var payload RecentPostsPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	posts, err := GetRecent(payload.Limit, payload.Skip)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, posts)
}
