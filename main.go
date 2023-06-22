package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
*
Receipt Struct
*/
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	id           string
	points       int
}

var processedReceipts []Receipt

/*
*
Main
*/
func main() {
	http.HandleFunc("/receipts/process", processReceipts)
	http.HandleFunc("/receipts/", getPoints)
	http.ListenAndServe("localhost:8080", nil)
}

/*
*
Receipt Processing Functions
*/
func openFiles() []Receipt {
	files, err := os.ReadDir("Receipts")
	if err != nil {
		log.Fatal(err)
	}
	var receiptsJSON []Receipt
	for i := 0; i < len(files); i++ {
		filepath := "Receipts/" + files[i].Name()
		receiptsJSON = append(receiptsJSON, openJSON(filepath))
		receiptsJSON[i].id = fmt.Sprintf("R%d", i)
		receiptsJSON[i].points = calculatePoints(receiptsJSON[i])
	}
	return receiptsJSON
}

func openJSON(filepath string) Receipt {

	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var receipt Receipt
	err = json.Unmarshal(file, &receipt)
	if err != nil {
		log.Fatal(err)
	}
	return receipt
}

/*
*
Point Calculation Functions
*/
func calculatePoints(receipt Receipt) int {
	points := 0
	points += retailerPoint(receipt.Retailer)
	points += roundDollarPoint(receipt.Total)
	points += quarterMultiplePoint(receipt.Total)
	points += twoItemPoint(receipt.Items)
	points += trimItemPoint(receipt.Items)
	points += oddDayPoint(receipt.PurchaseDate)
	points += betweenTwoFourPoint(receipt.PurchaseTime)
	return points
}

func retailerPoint(retailer string) int {
	return len(alphanumericStr(retailer))
}

func roundDollarPoint(total string) int {
	if strings.Contains(total, ".00") {
		return 50
	}
	return 0
}

func quarterMultiplePoint(total string) int {
	if strings.Contains(total, ".00") || strings.Contains(total, ".25") || strings.Contains(total, ".50") || strings.Contains(total, ".75") {
		return 25
	}
	return 0
}

func twoItemPoint(items []Item) int {
	return (len(items) / 2) * 5
}

func trimItemPoint(items []Item) int {
	sumPoint := 0
	for i := 0; i < len(items); i++ {
		trimmedItem := strings.TrimLeft(items[i].ShortDescription, " ")
		trimmedItem = strings.TrimRight(trimmedItem, " ")
		if len(trimmedItem)%3 == 0 {
			totalFloat, err := strconv.ParseFloat(items[i].Price, 8)
			if err != nil {
				log.Fatal(err)
			}
			totalFloat = math.Ceil(totalFloat * .2)
			sumPoint += int(totalFloat)
		}

	}
	return sumPoint
}

func oddDayPoint(date string) int {
	trimmedDate := alphanumericStr(date)
	dateInt, err := strconv.Atoi(trimmedDate)
	if err != nil {
		log.Fatal(err)
	}
	if dateInt%2 == 1 {
		return 6
	}
	return 0
}

func betweenTwoFourPoint(time string) int {
	trimmedTime := alphanumericStr(time)
	dateInt, err := strconv.Atoi(trimmedTime)
	if err != nil {
		log.Fatal(err)
	}
	if dateInt > 1400 && dateInt < 1600 {
		return 10
	}
	return 0
}

func alphanumericStr(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	str = strings.Replace(str, " ", "", -1)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

/*
*
HTTPS Function
*/
func processReceipts(write http.ResponseWriter, res *http.Request) {
	if res.URL.Path != "/receipts/process" {
		http.Error(write, "404", http.StatusNotFound)
		return
	}

	switch res.Method {
	case "POST":
		err := res.ParseForm()
		if err != nil {
			log.Fatal(err)
			return
		}
		processedReceipts = openFiles()
		jsonBody := "["
		for i := 0; i < len(processedReceipts); i++ {
			idJSON := fmt.Sprintf("{\"id\": \"%s\"}", processedReceipts[i].id)
			jsonBody += idJSON
			if i != len(processedReceipts)-1 && len(processedReceipts) > 2 {
				jsonBody += ","
			}
		}
		jsonBody += "]"
		write.Write([]byte(jsonBody))
	}
}

func getPoints(write http.ResponseWriter, res *http.Request) {
	id := strings.TrimLeft(res.URL.Path, "/receipts/'")
	id = strings.TrimRight(id, "/points'")

	switch res.Method {
	case "GET":
		for i := 0; i < len(processedReceipts); i++ {
			if id == processedReceipts[i].id {
				pointsJSON := fmt.Sprintf("{\"points\": \"%d\"}", processedReceipts[i].points)
				write.Write([]byte(pointsJSON))
				break
			}
		}
	}
}
