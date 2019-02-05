package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Job : The job struct holds what we think is the most important data.
type Job struct {
	JobType     string `json:"type"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

// getJob : The getJob function returns the first job posting as HTML in a string based on the location passed in.
func getJob(location string) string {
	// Calling GitHub jobs API
	response, _ := http.Get("https://jobs.github.com/positions.json?description=python&full_time=true&location=" + location)
	// Get body of the response
	data, _ := ioutil.ReadAll(response.Body)
	// Instantiate a new job struct
	var job []Job
	// Unmarshal data into a pointer to the job struct
	json.Unmarshal(data, &job)
	// Create HTML string
	jobPosting := "Company: <br>" + job[0].Company + "<br><br> Description: " + job[0].Description

	return jobPosting
}

func main() {
	// Echo instance
	e := echo.New()

	// Logger middleware
	e.Use(middleware.Logger())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		// location get the query parameter, so you must add /?location=sf to the end
		location := c.QueryParam("location")
		jobPost := getJob(location)
		return c.HTML(http.StatusOK, jobPost)
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
