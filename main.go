package main

import (
	//"github.com/joho/godotenv"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

// init is invoked before main()
/*func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
*/

func main() {
	// declare vars
	data := map[string]interface{}{}
	a := []string{}
	floats := []float64{}

	//should be an env var
	limit := 3
	count := 0
	var avg float64 = 0

	// should be change to env var values
	url := "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT"

	// get response from the url GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// pass to unmarshal
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}else{
		for k, value := range data {
		/*	fmt.Println("Key:", k, "Value:", value, "\n")*/
		if k == "Time Series (Daily)" {
				for _, item := range value.(map[string]interface{}) {
				//	fmt.Println("key: ", i , "value: ", item)
					for ix, items := range  item.(map[string]interface{}) {
						if ix == "4. close" {
						//	fmt.Println("-------************************-----------")
						//	fmt.Println("Id: ", count+1, "Key: ", ix , "Value: ", items)
							a = append(a, items.(string))
							count ++
						}
						if (limit == count) {
							for o, b := range a {
								var tmp, tmp2 = strconv.ParseFloat(b, o)
								tmp2 = tmp2
								floats = append(floats, tmp)
								// Calculate the average
								avg = (avg + tmp)
							}
							// cast limit to float and divide into the limit
							limit := float64(limit)
							avg = (avg/limit)

							fmt.Println("//////////////////////////////////////////")
							fmt.Println("Results: ====>>", floats)
							fmt.Println("Average: ====>>", avg)
							fmt.Println("//////////////////////////////////////////")
							os.Exit(0)
						}
					}
				}
			}
		}
	}
	os.Exit(0)
}

