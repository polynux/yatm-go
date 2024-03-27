package handlers

import (
	"context"
	"net/http"
	"polynux/yatm/db"
	"polynux/yatm/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
  ctx := context.Background()

  page, _ := strconv.Atoi(c.QueryParam("page"))

  limit := 20
  offset := page * limit

  params := db.GetPraticiensParams{
    Limit: int64(limit),
    Offset: int64(offset),
  }

  doctors, err := utils.Q.GetPraticiens(ctx, params)
  if err != nil {
    return err
  }

  data := map[string]interface{}{
    "Title": "Trouvez un docteur près de chez vous",
    "Doctors": doctors,
    "Page": page + 1,
  }

  return c.Render(http.StatusOK, "home.html", data)
}
