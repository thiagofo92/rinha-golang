//go:build windows
// +build windows

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"thiagofo92/api-test/mongo/connetion"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	POST_PEOPLE      = "/pessoas"
	GET_PERSON       = "/pessoas"
	GET_PEOPLES_TERM = "/pessoas"
	GET_COUNT        = "contagem-pessoas"
)

//	type Person struct {
//		NickName string   `json:"apelido" bson:"apelido"`
//		Name     string   `json:"nome" bson:"nome"`
//		BirthDay string   `json:"nascimento" bson:"nascimento"`
//		Stack    []string `json:"stack" bson:"stack"`
//	}
type Person struct {
	NickName string   `json:"apelido" bson:"nickname"`
	Name     string   `json:"nome" bson:"name"`
	BirthDay string   `json:"nascimento" bson:"birthday"`
	Stack    []string `json:"stack" bson:"stack"`
}

type People struct {
	conn *mongo.Database
}

func main() {
	route := chi.NewRouter()
	conn, err := connetion.NewConnectionMongo()

	if err != nil {
		panic(err)
	}
	people := People{}

	people.conn = conn.Database("rinha")

	route.Post("/pessoas", people.AddPerson)
	route.Get("/pessoas", people.GetPerson)
	route.Get("/pessoas/[:t]", people.GetPersonByTerm)
	route.Get("/contagem-pessoas", people.CountPerson)

	http.ListenAndServe(":3500", route)
}

func (p People) GetPerson(w http.ResponseWriter, r *http.Request) {

}

func (p People) AddPerson(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var data Person

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Printf("Error to unmarshalL: %v", err)
		w.WriteHeader(400)
		return
	}

	// birthday := strings.Split(data.BirthDay, "-")

	filter := bson.D{
		{
			Key:   "nickname",
			Value: data.NickName,
		},
	}

	mongoData := p.conn.Collection("people").FindOne(ctx, filter)

	var exist Person

	mongoData.Decode(&exist)

	if !reflect.DeepEqual(exist, Person{}) {
		w.WriteHeader(422)
		return
	}

	result, err := p.conn.Collection("people").InsertOne(ctx, data)

	if err != nil {
		panic(err)
	}

	location := "/pessoas/" + result.InsertedID.(primitive.ObjectID).Hex()

	w.Header().Set("Location", location)
	w.WriteHeader(201)
}

func (p People) GetPersonByTerm(w http.ResponseWriter, r *http.Request) {

}
func (p People) CountPerson(w http.ResponseWriter, r *http.Request) {

}
