package middleware

import (
	// Manipular o contexto
	"context"
	"encoding/json"
	"fmt"
	"log"

	//Servidor
	"net/http"
	//Sistema operacional + var de ambiente
	"os"
	"todo-project/models"

	// Gorilla pra facilitar o roteamento
	"github.com/gorilla/mux"
	// MongoDB
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// Variável de ambiente
	"github.com/joho/godotenv"
)

// Coleção do Mongo
var collection *mongo.Collection

func start() {
	loadEnv()
	createDBInstance()
}

func loadEnv() {
	err := godotenv.Load(".env")
	// Tratando erro
	if err != nil {
		log.Fatal("Erro no dot env")
	}
}

func createDBInstance() {
	// a := os.Getenv("DB_URI")
	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	dbCollection := os.Getenv("DB_COLLECTION")
	// voltar aqui depois
	clientOptions := options.Client().ApplyURI(dbUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// Conexão deu certo
	fmt.Println("Conectado ao Mongo")
	collection = client.Database(dbName).Collection(dbCollection)
	fmt.Println("collection criada")
}

func GetAllAnimes(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	payload := getAllAnimes()

	// collection.Find()
	json.NewEncoder(writer).Encode(payload)

}

func CreateAnime(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Acess-Control-Allow-Methods", "POST")
	writer.Header().Set("Acess-Control-Allow-Header", "Content-Type")
	writer.Header().Set("Acess-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version, Authorization")

	// Var do tipo struct Watchlist
	var anime models.WatchList
	// Decodificando o elemento pego pelo POST (json e bson)
	json.NewDecoder(route.Body).Decode(&anime)
	// Inserindo o anime
	insertOneAnime(anime)
	json.NewEncoder(writer).Encode(anime)

}

func AnimeFinished(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Acess-Control-Allow-Methods", "PUT")
	writer.Header().Set("Acess-Control-Allow-Header", "Content-Type")
	writer.Header().Set("Acess-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version, Authorization")

	params := mux.Vars(route)

	// VER ISSO DEPOIS
	animeFinished(params["id"])
	json.NewEncoder(writer).Encode(params["id"])
}

func UndoAnime(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Acess-Control-Allow-Methods", "DELETE")
	writer.Header().Set("Acess-Control-Allow-Header", "Content-Type")
	writer.Header().Set("Acess-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version, Authorization")
	params := mux.Vars(route)

	undoAnime(params["id"])
	json.NewEncoder(writer).Encode(params["id"])
}

func DeleteAnime(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Acess-Control-Allow-Methods", "DELETE")
	writer.Header().Set("Acess-Control-Allow-Header", "Content-Type")
	writer.Header().Set("Acess-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version, Authorization")

	params := mux.Vars(route)

	deleteOneAnime(params["id"])
}

func DeleteAllAnimes(writer http.ResponseWriter, route *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Acess-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Acess-Control-Allow-Origin", "*")
	writer.Header().Set("Acess-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version, Authorization")

	// aux := deleteAllAnimes()
	// json.NewEncoder(writer).Encode(aux)
}

// Replicando funcs

func getAllAnimes() []primitive.M {
	// Procurando na coleção
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	for cur.Next(context.Background()) {
		var result bson.M
		// decodificando
		error2 := cur.Decode(&result)
		if error2 != nil {
			log.Fatal(error2)
		}
		// Adicionando os animes
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results

}

func animeFinished(anime string) {
	id, _ := primitive.ObjectIDFromHex(anime)
	filter := bson.M{"_id": id}
	// trocando o status do anime de não visto pra visto
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contador: ", result.ModifiedCount)
}

func insertOneAnime(anime models.WatchList) {
	result, err := collection.InsertOne(context.Background(), anime)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Adicionado: ", result.InsertedID)
}

func undoAnime(anime string) {
	id, _ := primitive.ObjectIDFromHex(anime)
	filter := bson.M{"_id": id}
	// Atualizando o status
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contador: ", result.ModifiedCount)
}

func deleteOneAnime(anime string) {
	id, _ := primitive.ObjectIDFromHex(anime)
	filter := bson.M{"_id": id}
	// Deletando com mongo db function
	del, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deletado: ", del.DeletedCount)

}

// func deleteAllAnimes() int64 {
// 	// Delete many
// 	dels, err := collection.DeleteMany(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Elementos deletados : ", dels.DeletedCount)
// }
