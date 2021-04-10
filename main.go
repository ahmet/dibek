package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Any("/*", func(c echo.Context) error {
		response := make(map[string]interface{})

		headers := make(map[string][]string)
		for name, values := range c.Request().Header {
			_, keyExists := headers[name]

			if !keyExists {
				headers[name] = []string{}
			}

			headers[name] = append(headers[name], values...)
		}
		response["headers"] = headers

		response["method"] = c.Request().Method
		response["path"] = c.Request().URL.Path

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request().Body)

		if c.Request().Header.Get("Content-Type") == "application/json" {
			var result map[string]interface{}
			json.Unmarshal(buf.Bytes(), &result)
			response["body"] = result
		} else {
			response["body"] = buf.String()
		}

		queryParams := make(map[string][]string)
		for name, values := range c.Request().URL.Query() {
			_, keyExists := queryParams[name]

			if !keyExists {
				queryParams[name] = []string{}
			}

			queryParams[name] = append(queryParams[name], values...)
		}
		response["query"] = queryParams

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
