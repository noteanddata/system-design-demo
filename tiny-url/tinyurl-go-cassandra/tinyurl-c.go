package main

import "log"
import "net/http"
import "github.com/gocql/gocql"

// CassandraConfig contains configuration of Cassandra cluster to connect
type CassandraConfig struct {
	hostname    string
	keyspace    string
	consistency gocql.Consistency
}

var cassandraConfig = CassandraConfig{hostname: "127.0.0.1", keyspace: "tinyurl", consistency: gocql.One}
var sequenceCache *SequenceCache

func main() {
	log.Println("now to start tinyurl with cassandra")
	run()
}

func run() {
	http.HandleFunc("/generate_short_url", generateShortURL)
	http.HandleFunc("/s/", redirectToLongURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	var err error
	sequenceCache, err = createSequenceCache(cassandraConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func encodeBase62(id uint64) string {
	// standard "encoding/base64" seems operating at string level
	// table is a shuffled string of "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	table := "ZegsI24TAUn1jcQ7hwP9mlKFOf0SLHozDuE6xXGrti3dWaVqYbBCk85ypMvRJN"

	bytes := make([] byte, 0)
	for ; id != 0; {
		r := id%62
		ch := table[r]
		bytes = append([]byte{ch}, bytes...)
		id = id / 62
	}
	for size := 8 - len(bytes); size > 0; size-- {
		bytes = append(bytes, table[0])
	}
	return string(bytes)
}


func generateShortURL(w http.ResponseWriter, r *http.Request)  {
	longURL := r.FormValue("long_url")
	log.Println("longurl=", longURL)
	seq := sequenceCache.getSeq()
	shortURL := encodeBase62(seq)
	insertURL(shortURL, longURL)
}

func insertURL(shortURL string, longURL string) error {
	err := session.Query("insert into urls (short_url, long_url) values (?, ?)", shortURL, longURL).Exec()
	if  err != nil {
		log.Fatal(err)
	}
	return err
}

func getLongURL(shortURLKey string) string {
	var longURL string
	iter := session.Query("select long_url from urls where short_url = ? ", shortURLKey).Iter()
	for iter.Scan(&longURL) {
	}
	return longURL
}

func redirectToLongURL(w http.ResponseWriter, r *http.Request) {
	shortUrlKey := r.URL.Path[len("/s/"):]
	longURL := getLongURL(shortUrlKey)
	http.Redirect(w, r, longURL, 301)
}
