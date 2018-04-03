package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	. "github.com/laurentlouk/geo-trouble/config"
	. "github.com/laurentlouk/geo-trouble/dao"
	"github.com/laurentlouk/geo-trouble/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = AccidentsDAO{}
var config = Config{}
var color = Range{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
	color.Read()

	// remove old collection
	dao.RemoveAll()

	// insert database with CSV test data
	file, err := os.Open("accidents.csv")
	if err != nil {
		log.Println("file does not exist")
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// format to Time
		t, _ := time.Parse("2006-01-02", record[1])
		n, _ := strconv.Atoi(record[3])
		dao.Insert(models.Accident{ID: bson.NewObjectId(), Date: t, City: record[2], Number: n})
	}
	log.Println("DataBase filed")
}

func main() {
	r := mux.NewRouter()
	log.Println("Server started")
	// text response
	r.HandleFunc("/city/{city}/dates/{from}/{to}", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		c := make(chan int)
		vars := mux.Vars(r)
		city := vars["city"]
		startDate := vars["from"]
		endDate := vars["to"]

		go accidentFromDateCity(startDate, endDate, city, c)

		fmt.Fprintf(w, "<html><body><h4>There is %d accidents at %s from %s to %s\n</h4></body></html>", <-c, city, startDate, endDate)
	})
	// color response
	r.HandleFunc("/city/{city}/dates/{from}/{to}/color", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		c := make(chan int)
		vars := mux.Vars(r)
		city := vars["city"]
		startDate := vars["from"]
		endDate := vars["to"]

		go accidentFromDateCity(startDate, endDate, city, c)

		fmt.Fprintf(w, "<html><body style=\"background:%s\"><h1>%s</h1></body></body></html>", setColors(<-c), city)
	})
	http.ListenAndServe(config.Port, r)
}

func accidentFromDateCity(from, to, city string, c chan int) {
	fromDate, errFrom := time.Parse("2006-01-02", from)
	toDate, errTo := time.Parse("2006-01-02", to)
	sum := 0
	if errTo != nil || errFrom != nil {
		c <- sum
		return
	}
	accidents, err := dao.FindByDateCity(fromDate, toDate, city)
	if err != nil {
		log.Println("error parsing request FindByDateCity")
	}
	// sum of accidents
	for k := range accidents {
		sum += accidents[k].Number
	}
	c <- sum
}

func setColors(n int) string {
	res := "rgb(255, 255, 255)"
	switch {
	case n >= color.Blue[0] && n <= color.Blue[1]:
		res = "rgb(70, 70, 255)"
	case n >= color.Green[0] && n < color.Green[1]:
		res = "rgb(70, 186, 70)"
	case n >= color.Orange[0] && n < color.Orange[1]:
		res = "rgb(255, 165, 0)"
	case n >= color.Red[0]:
		res = "rgb(230, 70, 70)"
	}
	return res
}
