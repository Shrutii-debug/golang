package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shrutii-debug/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sruti-agarwal:NEWPASSWORD@golang-project.favmqw4.mongodb.net/"

const dbName = "netflix"
const colName = "watchlist"

//most important

var collection *mongo.Collection

//connect with mongodb

func init(){
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("mongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready")
}

//MONGODB helpers

//insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in database with id: ", inserted.InsertedID)

}

//update 1 record
func updateOneMovie(movieId string){
	//primitive.ObjectIDFromHex() - converts string into id acceptable by mongodb
	id, _ := primitive.ObjectIDFromHex(movieId)

	//bson.M is used if you want shorter or clearer results and order of elements doesnt matter
	//bson.D is used if order of elements matter
	//majorly we will be using bson.M
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:",result.ModifiedCount)

}

//delete 1 record
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got deleted with id: ", deleteCount)
}
//delete all records from mongoDB

func deleteAllMovies() int64 {
	//filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("no of mocies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
//read documentation
func getAllMovies() []primitive.M{
	//assuming no error happens, a cursor is returned
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		 err := cur.Decode(&movie)
		 if err != nil {
			log.Fatal(err)
		 }
		 movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
	//we dont get returned the actual values, we get return of a cursor which
	//is a mongoDB object and you have to loop through and pull out the values
}

//actual controllers - separate files

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}