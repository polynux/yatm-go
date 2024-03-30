package handlers

import (
	"context"
	"net/http"
	"polynux/yatm/db"
	"polynux/yatm/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Params struct {
  Page int
  Profession string
}

const limit = 20

func GetPraticiens(ctx context.Context, params Params) ([]db.Praticien, error) {
  if params.Profession != "" {
    dbParams := db.GetPraticiensByProfessionParams{
      Profession: params.Profession,
      Limit: limit,
      Offset: int64(params.Page * limit),
    }
    return utils.Q.GetPraticiensByProfession(ctx, dbParams)
  }

  return utils.Q.GetPraticiens(ctx, db.GetPraticiensParams{
    Limit: limit,
    Offset: int64(params.Page * limit),
  })
}

func HomeHandler(c echo.Context) error {
  ctx := context.Background()

  page, _ := strconv.Atoi(c.QueryParam("page"))
  profession := c.QueryParam("profession")

  params := Params{
    Page: page,
    Profession: profession,
  }

  doctors, err := GetPraticiens(ctx, params)
  if err != nil {
    return err
  }

  professions, err := utils.Q.GetAllProfessions(ctx)
  if err != nil {
    return err
  }

  data := map[string]interface{}{
    "Title": "Trouvez un docteur près de chez vous",
    "Doctors": doctors,
    "Page": page + 1,
    "Professions": professions,
    "Limit": limit,
  }

  return c.Render(http.StatusOK, "home.html", data)
}
