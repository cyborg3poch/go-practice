package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const LogicToken string = "iamsecuredtoken" //Public

// waitgroup pointer
var wg sync.WaitGroup

func main() {

	waitgroupExample()

}

func waitgroupExample() {
	var array = []string{
		"https://www.youtube.com",
		"https://www.javaguides.net",
		"https://www.brianstorti.com",
		"https://www.google.com",
		"https://www.facebook.com",
	}

	for _, website := range array {
		// create a go routine | parallism not concurency

		go getStatusCode(website)

		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(url string) {

	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%d status code for %s: \n", res.StatusCode, url)

}

func greeter(msg string) {

	for i := 0; i < 6; i++ {

		time.Sleep(3 * time.Millisecond)
		fmt.Println(msg)
	}
}

func Defaulthandler() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>I am alive</h1>"))

}

type item struct {
	Name     string   `json:"item_name"`
	Price    int      `json:"item_price"`
	Platform string   `json:"ecom_website"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJSON() {
	items := []item{
		{"MRF bat", 500, "amazon", []string{"sport", "toys"}},
		{"Apple watch", 25000, "flipkart", []string{"electronics", "apple"}},
		{"Logitech mouse", 670, "meesho", []string{"electronics", "logitech"}},
	}

	finaljson, err := json.MarshalIndent(items, "", "\t")
	CheckError(err)

	fmt.Printf("%s\n", finaljson)
}

func fileHandling() {
	content := "This need to be written on a txt file-Shivank Verma"

	file, err := os.Create("./mylogfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of written content is:", length)

	// deferring this to execute last
	defer file.Close()
}

func StructUsage() {
	shivank := User{"Shivank", 31, "shivank@email.co", true}

	name := shivank.getUserName()
	fmt.Println(name)
}

// Go doesnot support inheritance | Structs mimic classes functionality
type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

//Getter  Method to get Name of user

func (u User) getUserName() string {
	return u.Name
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func forLoopwMap() {

	//Initialising a MAP
	names := make(map[string]string)
	names["sv"] = "Shivank Verma"
	names["mv"] = "Mayank Verma"
	names["hv"] = "Harshit Verma"

	//Foreach  Loop
	for key, value := range names {
		fmt.Printf("for Key %v , Value is %v\n", key, value)
	}

	// initialising a slice
	days := []string{"Sunday", "Monday", "tuesday", "Wednesday", "Thurusday"}

	//For Loop2
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

}

func basicio() {
	//fmt.Println("Hi from cyborg")
	var username string = "Shivank"
	fmt.Println(username)
	fmt.Printf("Variable of type : %T \n", username)

	// implicit declaration
	var websitename = "kwikbiz.in"
	fmt.Println(websitename)

	// no var style | walrus operator | not allowed outside the method
	someBigDFloat := 30000.00
	fmt.Println(someBigDFloat)

	//Reading user input via bufiio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some value :")

	//comma -ok || comma-err syntax for error handling
	input, _ := reader.ReadString('\n')

	fmt.Println("You have entered \t", input)

	// Type concersion using "strconv"
	changeduserInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 5 to your input :", changeduserInput+5)
	}
}
