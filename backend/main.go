package main 
import ("fmt" ; "net/http" ; "encoding/json" ; "time" ; "context" ; "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" ; "go.mongodb.org/mongo-driver/bson/primitive")

var client *mongo.Client
var stravausercollection *mongo.Collection

type StravaUser struct {
	 ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}


func createStravaUser(w http.ResponseWriter , r *http.Request) {
	if r.Method != "POST" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed);
		return ;
	}
	var stravauser StravaUser;
	err := json.NewDecoder(r.Body).Decode(&stravauser);
	if err != nil {
		http.Error(w , "Invalid request body" , http.StatusBadRequest);
		return ;
	}

	count , err := stravausercollection.CountDocuments(context.Background() , bson.M{"email" : stravauser.Email});
	if err != nil {
		http.Error(w , "Error checking for existing stravauser" , http.StatusInternalServerError);
		return ;
	}
	if count > 0 {
		http.Error(w , "StravaUser with this email already exists" , http.StatusConflict);
		return ;
	}

	result , err := stravausercollection.InsertOne(context.Background() , stravauser);
	if err != nil {
		http.Error(w , "Error inserting stravauser" , http.StatusInternalServerError);
		return ;
	}
	fmt.Println("Inserted stravauser with ID: " , result.InsertedID);
}


func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next(w, r)
    }
}



func main() {
	ctx , cancel := context.WithTimeout(context.Background() , 10*time.Second);
	defer cancel();
	var err error;
	client , err = mongo.Connect(ctx , options.Client().ApplyURI("mongodb://localhost:27017"));
	if err != nil {
		fmt.Println("Error connecting to MongoDB: " , err);
		return ;
	}

	stravausercollection = client.Database("stravaDB").Collection("stravausercollection");
	
	
	http.HandleFunc("/signup" , enableCORS(createStravaUser));
	fmt.Println("Server running on http://localhost:8080");
	 http.ListenAndServe(":8080", nil);

}