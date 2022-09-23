package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		SetDebug(true).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
	}

func TestCreate(t *testing.T) {
	r := gofight.New()

	r.POST("/create").
		SetDebug(true).
		SetForm(gofight.H{
			"name":  "Apple Watch",
			"price": "1000",
		}).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			name, _ := jsonparser.GetString(data, "name")
			price, _ := jsonparser.GetString(data, "price")

			assert.Equal(t, "Apple Watch", name)
			assert.Equal(t, "1000", price)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestRead(t *testing.T) {
	r := gofight.New()

	r.GET("/read").
		SetDebug(true).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestUpdate(t *testing.T) {
	r := gofight.New()

	r.PUT("/update").
		SetDebug(true).
		SetForm(gofight.H{
			"name":  "Apple Watch",
			"price": "2000",
		}).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			name, _ := jsonparser.GetString(data, "name")
			price, _ := jsonparser.GetString(data, "price")

			assert.Equal(t, "Apple Watch", name)
			assert.Equal(t, "2000", price)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestDelete(t *testing.T) {
	r := gofight.New()

	r.DELETE("/delete/2").
		SetDebug(true).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}