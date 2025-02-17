package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/sam-peets/micro/db"
	"github.com/sam-peets/micro/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Sid     string    `bson:"sid" json:"sid"`
	Uid     uint32    `bson:"uid" json:"uid"`
	Expires time.Time `bson:"expires" json:"expires"`
}

func (sess *Session) Verify() bool {
	return sess.Expires.After(time.Now())
}

func (sess *Session) GetUser() (*user.User, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	var u user.User
	err = users.FindOne(context, bson.M{"uid": sess.Uid}).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetSessionBySid(sid string) (*Session, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	sessions, err := connection.GetCollection("sessions")
	if err != nil {
		return nil, err
	}

	var sess Session
	err = sessions.FindOne(context, bson.M{"sid": sid}).Decode(&sess)
	if err != nil {
		fmt.Println("ruh roh")
		return nil, err
	}

	if sess.Verify() {
		return &sess, nil
	} else {
		return nil, errors.New("session expired")
	}
}

func NewSession(user user.User) (*Session, error) {
	now := time.Now()

	sid := make([]byte, 32)
	rand.Read(sid)
	hexSidHash := hex.EncodeToString(sid)
	sess := Session{
		Sid:     hexSidHash,
		Uid:     user.Uid,
		Expires: now.AddDate(0, 0, 1),
	}

	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	sessions, err := connection.GetCollection("sessions")
	if err != nil {
		return nil, err
	}

	sessions.InsertOne(context, sess)
	return &sess, nil
}

type UserPayload struct {
	Username string
	Password string
}

func NewUser(payload *UserPayload) (*user.User, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	res := users.FindOne(context, bson.M{"username": payload.Username})
	if !errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, errors.New("user already exists")
	}

	hashed_password := sha256.Sum256([]byte(payload.Password))

	opts := options.FindOne().SetSort(bson.M{"uid": -1})
	var lastUser user.User
	err = users.FindOne(context, bson.M{}, opts).Decode(&lastUser)
	var uid uint32
	if err != nil {
		uid = 1
	} else {
		uid = lastUser.Uid + 1
	}

	u := user.User{
		Hash:     hashed_password,
		Uid:      uid,
		Username: payload.Username,
	}

	_, err = users.InsertOne(context, u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func Auth(payload *UserPayload) (*Session, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	hashed_password := sha256.Sum256([]byte(payload.Password))
	var u user.User
	err = users.FindOne(context, bson.M{"username": payload.Username, "hash": hashed_password}).Decode(&u)
	if err != nil {
		return nil, err
	}

	if u.Hash != hashed_password {
		return nil, errors.New("incorrect password")
	}

	sessions, err := connection.GetCollection("sessions")
	if err != nil {
		return nil, err
	}

	var sess Session
	err = sessions.FindOne(context, bson.M{"uid": u.Uid}).Decode(&sess)
	if err != nil {
		nsess, err := NewSession(u)
		if err != nil {
			return nil, err
		}
		return nsess, nil
	}

	if sess.Verify() {
		return &sess, nil
	} else {
		_, err := sessions.DeleteOne(context, bson.M{"uid": u.Uid})
		if err != nil {
			return nil, fmt.Errorf("error deleting old session: %w", err)
		}

		nsess, err := NewSession(u)
		if err != nil {
			return nil, err
		}
		return nsess, nil
	}
}
