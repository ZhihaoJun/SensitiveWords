package main

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	e         *echo.Echo
	wordsFile string
	words     []string
}

func (a *App) checkSensitives(c echo.Context) error {
	str := c.QueryParam("s")
	sensitive := ""
	for _, word := range a.words {
		if strings.Contains(str, word) {
			sensitive = word
			break
		}
	}
	if len(sensitive) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": "ok",
			"msg":   "check sensitive success",
			"sensitives": []string{
				sensitive,
			},
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":      "ok",
		"msg":        "check sensitive success",
		"sensitives": []string{},
	})
}

func (a *App) ping(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (a *App) loadEnv() {
	a.wordsFile = "/sensitives.txt"
}

func (a *App) readFile(fp string) {
	absFp, err := filepath.Abs(fp)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(absFp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a.words = append(a.words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (a *App) logSkipper(c echo.Context) bool {
	return c.Request().RequestURI == "/ping"
}

func (a *App) Setup() {
	a.loadEnv()
	a.readFile(a.wordsFile)

	a.e = echo.New()
	a.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: a.logSkipper,
	}))
	a.e.Use(middleware.Recover())
	a.e.GET("/api/check", a.checkSensitives)
	a.e.HEAD("/ping", a.ping)
	a.e.GET("/ping", a.ping)
}

func (a *App) Run() {
	a.e.Start(":1323")
}

func (a *App) Close() {

}

func main() {
	a := &App{}
	a.Setup()
	defer a.Close()
	a.Run()
}
