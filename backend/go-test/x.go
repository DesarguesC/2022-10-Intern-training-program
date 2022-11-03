package main
import(
	"fmt";
	"context";
	"log";
	"go.mongodb.org/mongo-driver/bson";
	"go.mongodb.org/mongo-driver/mongo";
	"go.mongodb.org/mongo-driver/mongo/options";
)

type  one struct {
	Id int;
	Name string;

}

func main() {
	url := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("Have connected successfully")
	collection := client.Database("test").Collection("stu")
	// link the database
	// 
	
	insert(collection)

	delete(collection)

	read(collection)

	update(collection)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All has been done")
}

func update(collection *mongo.Collection) {
	fmt.Println("chang the name of a person (input the name)")
	var x string
	fmt.Scanf("%s\n", &x)
	filter := bson.M{"name":x}
	update_ := bson.M {
		"$inc": bson.M {
			"name": x,
		},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update_)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Having matched %v document(s) and updated %v document(s).", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func insert(collection *mongo.Collection) {
	fmt.Println("Input user(s).\nFormat: id, name.\nwhere id is an integer and name is a srting.")
	var (
		x int;
		y string
	)
	fmt.Scanf("%d %s\n", &x, &y)
	insertResult, err := collection.InsertOne(context.TODO(), one{x, y})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Having inserted: ", insertResult.InsertedID)
}
func delete(collection *mongo.Collection) {
	fmt.Println("Input a name that you wanna delete.")
	var x string
	fmt.Scanf("%s\n", &x)
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.M{"name": x})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Having deleted: ", deleteResult1.DeletedCount)
}

func read(collection *mongo.Collection) {
	fmt.Println("Input a name that you wanna search.")
	var x string
	fmt.Scanf("%s\n", &x)
	filter := bson.M{"name": x}
	var result one
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a single documetn: %v", result)

	findOptions := options.Find()
	findOptions.SetLimit(5)

	var ones []*one

	curr,  err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for curr.Next(context.TODO()) {
		var ele one
		err := curr.Decode(&ele)
		if err != nil {
			log.Fatal(err)
		}
		ones = append(ones, &ele)
	}
	if err := curr.Err();err != nil {
		log.Fatal(err)
	}
	if err := curr.Err();err != nil {
		log.Fatal(err)
	}
}

