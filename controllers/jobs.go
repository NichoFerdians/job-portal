package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NichoFerdians/job-portal/token"

	"github.com/gin-gonic/gin"
)

func JobList(c *gin.Context) {
	_ = c.MustGet(authorizationPayloadKey).(*token.Payload)

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://dev3.dansmultipro.co.id/api/recruitment/positions.json", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// c.JSON(http.StatusOK, resp.Body)
	c.Data(200, "application/json", bodyBytes)
}

func JobDetail(c *gin.Context) {
	_ = c.MustGet(authorizationPayloadKey).(*token.Payload)

	id := c.Query("id")
	fmt.Println("id", "http://dev3.dansmultipro.co.id/api/recruitment/positions/"+id)

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://dev3.dansmultipro.co.id/api/recruitment/positions/"+id, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// c.JSON(http.StatusOK, resp.Body)
	c.Data(200, "application/json", bodyBytes)
}
