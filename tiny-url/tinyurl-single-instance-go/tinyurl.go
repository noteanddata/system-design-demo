package main
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"
const db_url = "tinyr_url_user:Adeg*#%23f@tcp(localhost:3306)/tinyurl_single"

func main() {
  run()
}

func run() {
  log.Println("hello, tinyurl")
  short_url := insert_url("http://www.yahoo.com")
  log.Println(short_url)
  long_url := get_long_url(short_url)
  log.Println(long_url)
}

func encode_base_62(id int64) string {
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


func insert_url(long_url string) string {
  db, err  := sql.Open("mysql", db_url)
  if err != nil {
    log.Println(err)
    panic("Error opening db with url " + db_url + err.Error())
  }
  defer db.Close()
  
  ret, err := db.Exec("insert into sequences values ();")
  if err != nil {
    log.Println(err)
    panic("Error inserting to sequences" + err.Error())
  } 
  
  id, err := ret.LastInsertId()
  if err != nil {
    log.Println(err)
    panic("Error getting last insert id" + err.Error())
  }
  
  statement, err := db.Prepare("insert into urls (id, full_url, short_url_key) values (?, ?, ?)")  
  if err != nil {
    log.Println(err)
    panic("Error preparing insert to urls sql" + err.Error())
  }
  defer statement.Close()
  
  // encode to a short_url_key
  short_url_key := encode_base_62(id)
  
  ret, err = statement.Exec(id, long_url, short_url_key)
  if err != nil {
    log.Println(err)
    panic("Error inserting urls" + err.Error())
  }
  
  return short_url_key
}

func get_long_url(short_url_key string) string {
  db, err  := sql.Open("mysql", db_url)
  if err != nil {
    log.Println(err)
    panic("Error opening db with url " + db_url + err.Error())
  }
  defer db.Close()

  statement, err := db.Prepare("select full_url from urls where short_url_key = ?")
  if err != nil {
    log.Println(err)
    panic("Error preparing select full_url with short_url_key=" + short_url_key + err.Error())
  }
  defer statement.Close()
  
  rows, err := statement.Query(short_url_key)
  if err != nil {
    log.Println(err)
    panic("Error running query" + err.Error())
  }
  var long_url string
  if rows.Next() {
    rows.Scan(&long_url)
  }
  
  return long_url
}