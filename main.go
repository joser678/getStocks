package main

import (
	//"github.com/joho/godotenv"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

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
	data := map[string]interface{}{}
	a := []string{}

	limit := 4    //should be an env var
	count := 0
	// should be change to env var values
	url := "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT"

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
							fmt.Println("-------************************-----------")
							fmt.Println("Id: ", count+1, "Key: ", ix , "Value: ", items)
							a = append(a, items.(string))
							count ++
						}
						if (limit == count) {
							fmt.Println("//////////////////////////////////////////")
							fmt.Println("Results: ====>>", a)
							fmt.Println("Average: ====>>")
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

