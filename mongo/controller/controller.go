package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karmsetu/maongo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchList"

// !imp
var collection *mongo.Collection

// connect with mongoDB
func init(){
	// client option 
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to DB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance is ready")


}

// mongo helpers - file

// insert one record
func insertOneMovie(movie model.Netflix)  {
	inserted, err := collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in DB with id:", inserted.InsertedID)
}

// update 1 record 
func updateOneMovie(movieId string)  {
	id, err := primitive.ObjectIDFromHex(movieId)
	CheckIsError(err)
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	CheckIsError(err)
	fmt.Println("modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	CheckIsError(err)
	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	CheckIsError(err)
	fmt.Println("movie got deleted with delete count:", deleteCount)
}

// delete all records from mongo db
func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter ,nil)
	CheckIsError(err)
	fmt.Println("number of movies deleted::", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all collection
func getAllMovies() []primitive.M  {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	CheckIsError(err)

	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		CheckIsError(err)
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

// actual controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	CheckIsError(err)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

	
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()

	json.NewEncoder(w).Encode(count)

	
}

func CheckIsError(err error){
if err != nil {
		log.Fatal(err)
	}
}