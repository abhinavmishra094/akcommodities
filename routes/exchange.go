package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gitee.com/golang-module/carbon"
	"github.com/gin-gonic/gin"
)

func Exchange() func(c *gin.Context) {
	return func(c *gin.Context) {
		current_time := carbon.Now().ToDateString()
		time_ten := carbon.Now().SubDays(10).ToDateString()
		log.Println(current_time)
		log.Println(time_ten)
		url := fmt.Sprintf("http://api.currencylayer.com/timeframe?access_key=%s&currencies=USD,EUR,GBP&start_date=%v&end_date=%v", os.Getenv("ACCESS_KEY"), time_ten, current_time)
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(string(body)), &jsonMap)
		c.JSON(http.StatusOK, jsonMap)
	}
}
