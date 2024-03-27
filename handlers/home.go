package handlers

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "polynux/yatm/utils"
)

func HomeHandler(c echo.Context) error {
  doctors := utils.GetDoctors()
  data := map[string]interface{}{
    "Title": "Trouvez un docteur près de chez vous",
    "Doctors": doctors,
  }

  return c.Render(http.StatusOK, "home.html", data)
}
