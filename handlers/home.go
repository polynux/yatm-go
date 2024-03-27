package handlers

import (
	"context"
	"net/http"
	"polynux/yatm/utils"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
  ctx := context.Background()

  doctors, err := utils.Q.GetPraticiens(ctx)
  if err != nil {
    return err
  }

  data := map[string]interface{}{
    "Title": "Trouvez un docteur près de chez vous",
    "Doctors": doctors,
  }

  return c.Render(http.StatusOK, "home.html", data)
}
