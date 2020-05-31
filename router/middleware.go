package router

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"log"
	"net/http"
    "time"
)

func parseAndExecute(w http.ResponseWriter, r *http.Request, filename string, execute interface{}) {
	t := template.Must(template.ParseFiles("./static/index.html", "./static/navigation.html"))
	template.Must(t.ParseFiles(filename))
	t.Execute(w, execute)
}

func tweets(w http.ResponseWriter, r *http.Request) {
    tweets, err := GetTweets("")
    if err != nil {
        log.Printf("failed fetching tweets: %v", err)
    }

	parseAndExecute(w, r, "./static/posts.html", tweets)
}

func tweet(w http.ResponseWriter, r *http.Request) {
	parseAndExecute(w, r, "./static/tweet.html", "")

	if r.Method == http.MethodPost {
		r.ParseForm()

        tweet := Tweet{
            ID: primitive.NewObjectID(),
            Body: r.Form["tweet"][0],
            Date: time.Now(),
            UserIP: r.RemoteAddr,
        }

        user := User{
            ID: primitive.NewObjectID(),
            UserIP: r.RemoteAddr,
            Following: []primitive.ObjectID{},
        }

        performPost(tweet, "tweets")
        performPost(user, "users")
    }
}

func followed(w http.ResponseWriter, r *http.Request) {
	parseAndExecute(w, r, "./static/followed.html", "")
}
