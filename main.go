package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func numberFromServiceB(host string) (int, error) {
	res, err := http.Get(fmt.Sprintf("http://%s/n", host))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.Atoi(string(body))
	return ret, err
}

func main() {
	router := gin.Default()

	serviceB := os.Getenv("SERVICE_B_HOST")
	if serviceB == "" {
		panic("Missing SERVICE_B_HOST environment variable")
	}

	router.GET("/do-stuff", func(c *gin.Context) {
		n, err := numberFromServiceB(serviceB)
		if err != nil {
			c.String(http.StatusInternalServerError, "Sorry... something went wrong")
		} else {
			c.String(http.StatusOK, fmt.Sprint(n+1))
		}
	})

	router.Run(":8080")
}
