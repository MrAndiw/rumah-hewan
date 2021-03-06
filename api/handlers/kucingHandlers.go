package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Kucing struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetKucingFunc(c echo.Context) error {

	kucing := c.QueryParam("kucing")
	status := c.QueryParam("status")

	dataType := c.Param("type")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Ini kucingnya %s \nDan ini statusnya %s", kucing, status))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"kucing": kucing,
			"status": status,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "tipe harus string atau json",
	})
}

func AddKucingFunc(c echo.Context) error {
	kucing := Kucing{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&kucing)
	if err != nil {
		log.Printf("Gagal melakukan decode %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "Gagal melakukan decode",
		})
	}

	// == save to database here ==
	log.Printf("Berhasil Menyimpan kucing dari request %v", kucing)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
	})
}
