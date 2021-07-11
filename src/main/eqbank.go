package main

import (
  "encoding/json"
  "fmt"
  "html/template"
  "log"
  "net/http"
  "os"
)

//DAO MODELING
type Admin interface{}
type User struct { //USER Data Model
  Username string `json:"User"`
  Password string `json:"pwd"`
  Name string `json:"full_name"`
  ID int `json:"id"`
  Role string `json:"role"`
  //Age int `json:"age"`
  Accounts []int `json:"accounts"`

  //Active bool `json:"active"`
  //lastLoginAt string
}

//TESTING MODULE
func jocAdmin() []byte { //Can Marshal/UnMarshal JSON to Struct, vice versa
  model:= User{}
  n:= User{
    Username: "User1",
    Name: "JOC",
    ID:1,
    Role: "Admin",
    Password: "007"}

  u, err := json.Marshal(n)
  if err != nil {
    panic(err)}
  _ = json.Unmarshal(u, &model)
  fmt.Println(string(u),"\n",n, "\n", model) // {"Name":"Bob","Age":10,"Active":true}
  return u
}

//ROUTING TABLE
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	html, _:= template.ParseFiles("./front-end/templates/index.html")
	err:= html.Execute(w,nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}}
func loginHandler(w http.ResponseWriter, r *http.Request) {
  //model:= User{}
  if r.URL.Path != "/login" || r.Method!=http.MethodPost {
    http.NotFound(w, r)
    return}
  err := r.ParseForm()
  if err != nil {log.Fatal(err)}
  login:= User{
    Username: r.FormValue("uname"),
    Password: r.FormValue("pwd")}
  welcome,_:= json.Marshal(login)

  fmt.Println(string(welcome))

  html, _:= template.ParseFiles("./front-end/templates/birthday.html") //RIGHT NOW WE DO BIRTHDAY
  err= html.Execute(w,login)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
  }}

//START
func main() {
  jocAdmin() //Add to database
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/login", loginHandler)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Defaulting to port %s", port)
  } // http.ListenAndServe(":8080", nil)

  log.Printf("Listening on port %s", port)
  log.Printf("Open http://localhost:%s in the browser", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
