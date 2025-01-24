package posts

import (
	"math/rand"
	"time"

	"github.com/sam-peets/micro/db"
	"github.com/sam-peets/micro/user"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Uid       uint32    `bson:"uid"`
	PostId    uint32    `bson:"postid"`
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"timestamp"`
}

func (p *Post) User() (*user.User, error) {
	return user.GetUser(p.Uid)
}

func NewPost(u user.User, content string) (*Post, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	posts, err := connection.GetCollection("posts")
	if err != nil {
		return nil, err
	}

	p := Post{
		Uid:       u.Uid,
		PostId:    rand.Uint32(),
		Content:   content,
		Timestamp: time.Now(),
	}

	_, err = posts.InsertOne(context, p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func GetPost(id uint32) (*Post, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	posts, err := connection.GetCollection("posts")
	if err != nil {
		return nil, err
	}

	var p Post
	err = posts.FindOne(context, bson.M{"postid": id}).Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func GetRecent(limit int) ([]Post, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	posts, err := connection.GetCollection("posts")
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetSort(bson.M{"_id": -1}).SetLimit(int64(limit))

	cur, err := posts.Find(context, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var recent_posts []Post
	err = cur.All(context, &recent_posts)
	if err != nil {
		return nil, err
	}
	return recent_posts, nil
}
