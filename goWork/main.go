package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type product struct {
	productURL         string
	Img                string
	Img1               string
	Img2               string
	Img3               string
	Brand              string
	skuStyleIsNow      string
	season             string
	year               string
	skuSupplier        string
	variant            string
	colorDetail        string
	colorSupplier      string
	madeIn             string
	material           string
	Name               string
	description        string
	categories         string
	Qty                string
	RetailPrice        string
	discount           string
	sizeInfo           string
	size               string
	qtyDetail          string
	bagLength          string
	bagHeight          string
	bagWeight          string
	handleHeight       string
	shoulderBagLength  string
	beltLength         string
	beltHeight         string
	accessoryLength    string
	accessoryHeight    string
	accessoryWeight    string
	heelHeight         string
	plateauHeight      string
	insoleLength       string
	colorStyleisnowITA string
	FTA                string
	EAN                string
	nomeITA            string
	descrizioneITA     string
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

	print("test")
	//fi, err := os.Open("C:\\Users\\user\\Desktop\\Workspace\\goWork\\1.txt")
	//fo, err := os.Create("C:\\Users\\user\\Desktop\\Workspace\\goWork\\2.txt")

	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		rst := parsingCSV()
		//d := (*rst)[0]
		cntList := len(*rst)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":     "ParseCSV",
			"FirstName": "hodong",
			"length":    cntList,
			"URL":       c.Request.URL.Path,
			"rst":       rst,
		})
	})

	return r
}

func parsingCSV() *[]product {
	read, _ := os.Open("./allProducts.csv")
	defer read.Close()

	const maxCapacity = 1024 * 1024 * 16
	//buf := make([]byte, maxCapacity)
	reader := bufio.NewReaderSize(read, maxCapacity)
	r := csv.NewReader(reader)
	r.Comma = ';'
	csvLines, err := r.ReadAll()

	var resultsSet []product

	if err != nil {
		fmt.Println(err)
	}
	coloumLength := len(csvLines[0])
	println(coloumLength)

	for _, line := range csvLines {
		vo := product{
			line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8], line[9], line[10],
			line[11], line[12], line[13], line[14], line[15], line[16], line[17], line[18], line[19], line[20],
			line[21], line[22], line[23], line[24], line[25], line[26], line[27], line[28], line[29], line[30],
			line[31], line[32], line[33], line[34], line[35], line[36], line[37], line[38], line[39], line[40], line[41]}

		//fmt.Println(line[0] + " " + line[1] + " " + line[2])
		resultsSet = append(resultsSet, vo)
	}

	return &resultsSet
}
