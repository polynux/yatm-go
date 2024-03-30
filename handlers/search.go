package handlers

import (
	"context"
	"fmt"
	"net/http"
	"polynux/yatm/db"
	"strconv"

	"github.com/lithammer/fuzzysearch/fuzzy"

	"github.com/labstack/echo/v4"
)

func Search(doctors *[]db.Praticien, query string) []db.Praticien {
	result := []db.Praticien{}

	for _, doctor := range *doctors {
		switch {
		case fuzzy.MatchFold(query, doctor.Name):
			result = append(result, doctor)
		case fuzzy.MatchFold(query, doctor.Profession):
			result = append(result, doctor)
		case fuzzy.MatchFold(query, doctor.Address):
			result = append(result, doctor)
		case fuzzy.MatchFold(query, doctor.City):
			result = append(result, doctor)
		case fuzzy.MatchFold(query, doctor.Zip):
			result = append(result, doctor)
		}
	}

	return result
}

func SearchHandler(c echo.Context) error {
	ctx := context.Background()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	profession := c.QueryParam("profession")

	params := Params{
		Page:       page,
		Profession: profession,
	}

	doctors, err := GetPraticiens(ctx, params)
	if err != nil {
		return err
	}

	query := c.QueryParam("search")
	result := []db.Praticien{}
	if query != "" {
	      result = Search(&doctors, query)
  }
  fmt.Println(result)

	data := map[string]interface{}{
		"Doctors": result,
		"Page":    page + 1,
	}

	return c.Render(http.StatusOK, "home.html", data)
}
