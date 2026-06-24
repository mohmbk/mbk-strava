package main 
import ("fmt" ; "net/http" ; "encoding/json" ; "time" ; "context" ; "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" ; "go.mongodb.org/mongo-driver/bson/primitive" ; "github.com/golang-jwt/jwt" ; "strings" )

var client *mongo.Client
var stravausercollection *mongo.Collection
var sessioncollection *mongo.Collection

type StravaUser struct {
	 ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type loginuser struct {
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type session struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"userId" json:"userId"`
	TITLE string `json:"title" bson:"title"`
	Distance int `json:"distance" bson:"distance"`
	Time int `json:"time" bson:"time"`
}

type stat struct {
	TotalDistance int `json:"totalDistance" bson:"totalDistance"`
	TotalTime int `json:"totalTime" bson:"totalTime"`
	averagetotalpace float64 `json:"averagePace" bson:"averagePace"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type sessionsent struct {
	Title string `json:"title"`
	Distance int `json:"distance"`
	Time int `json:"time"`
}

var jwtSecret = []byte("ma_cle_secrete")


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


func login(w http.ResponseWriter , r *http.Request) {
	if r.Method != "POST" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed);
		return ;
	}
	var loginuser loginuser;
	err := json.NewDecoder(r.Body).Decode(&loginuser);
	if err != nil {
		http.Error(w , "Invalid request body" , http.StatusBadRequest);
		return ;
	}
	var stravauser StravaUser;
	err = stravausercollection.FindOne(context.Background() , bson.M{"email" : loginuser.Email}).Decode(&stravauser);
	if err != nil {
		http.Error(w , "StravaUser not found" , http.StatusNotFound);
		return ;
	}

	if stravauser.Password != loginuser.Password {
		http.Error(w , "Invalid password" , http.StatusUnauthorized);
		return ;
	}
	

	claims := jwt.MapClaims{
		"userId": stravauser.ID,
		"email": stravauser.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := loginResponse{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(response);
	

}


func getsession(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed);
		return ;
	}

	authHeader := r.Header.Get("Authorization");
	tokenString := strings.TrimPrefix(authHeader, "Bearer ");

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    	return jwtSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims);
	userId := claims["userId"] ;

	cursor, err := sessioncollection.Find(context.Background() , bson.M{"userId" : userId});
	if err != nil {
		http.Error(w , "Error fetching sessions" , http.StatusInternalServerError);
		return ;
	}

	var sessions []session;
	for cursor.Next(context.Background()) {
		var session session
		err := cursor.Decode(&session);
		if err != nil {
			http.Error(w , "Error decoding session" , http.StatusInternalServerError);
			return ;
		}
		sessions = append(sessions , session);
	}

	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(sessions);

	
}


func getstat(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed);
		return ;
	}

	authHeader := r.Header.Get("Authorization");
	tokenstring := strings.TrimPrefix(authHeader, "Bearer ");

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims);
	userId := claims["userId"];

	cursor , err := sessioncollection.Find(context.Background() , bson.M{"userId" : userId});
	if err != nil {
		http.Error(w , "Error fetching sessions" , http.StatusInternalServerError);
		return ;
	}

	var sessions []session;
	for cursor.Next(context.Background()) {
		var session session ;
		err := cursor.Decode(&session);
		if err != nil {
			http.Error(w , "Error decoding session" , http.StatusInternalServerError);
			return ;
		}
		sessions = append(sessions , session);
	}

	var totalDistance int = 0 ;
	var totalTime int = 0 ;
	var totalPace float64 = 0.0 ;
	for _, session := range sessions {
		totalDistance += session.Distance ;
		totalTime += session.Time ;
	}

	totalPace = float64(totalTime) / float64(totalDistance) ;

	stat := stat{
		TotalDistance: totalDistance,
		TotalTime: totalTime,
		averagetotalpace: totalPace,
	}

	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(stat);
}


func createSession(w http.ResponseWriter , r *http.Request) {
	if r.Method != "POST" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed);
		return ;
	}
	var newSession sessionsent;
	err := json.NewDecoder(r.Body).Decode(&newSession);
	if err != nil {
		http.Error(w , "Invalid request body" , http.StatusBadRequest);
		return ;
	}

	authHeader := r.Header.Get("Authorization");
	tokenString := strings.TrimPrefix(authHeader, "Bearer ");

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims);
	userId := claims["userId"];

	session := session{
		UserID: userId.(primitive.ObjectID),
		TITLE: newSession.Title,
		Distance: newSession.Distance,
		Time: newSession.Time,
	}

	result, err := sessioncollection.InsertOne(context.Background(), session);
	if err != nil {
		http.Error(w , "Error inserting session" , http.StatusInternalServerError);
		return ;
	}

	fmt.Println("Inserted session with ID: " , result.InsertedID);


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
	sessioncollection = client.Database("stravaDB").Collection("sessioncollection");
	
	http.HandleFunc("/stat" , enableCORS(getstat));
	http.HandleFunc("/signup" , enableCORS(createStravaUser));
	http.HandleFunc("/login" , enableCORS(login));
	http.HandleFunc("/session" , enableCORS(getsession));
	fmt.Println("Server running on http://localhost:8080");
	 http.ListenAndServe(":8080", nil);

}