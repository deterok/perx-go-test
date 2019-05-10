package views

import (
	"net/http"

	"github.com/labstack/echo"

	"perx-go-test/resources"
)

type CodesCheckView struct {
	codesRes *resources.CodesResource
}

func NewCodesCheckView(codesRes *resources.CodesResource) *CodesCheckView {
	return &CodesCheckView{
		codesRes: codesRes,
	}
}

func (v *CodesCheckView) Post(c echo.Context) error {
	request := new(struct {
		Code string `json:"code" validate:"required"`
	})

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := v.codesRes.Check(request.Code)
	if err != nil {
		if err == resources.ErrCodeNotFound || err == resources.ErrCodeAlreadyUsed {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := struct {
		Checked bool `json:"checked"`
	}{
		Checked: result,
	}

	return c.JSON(http.StatusOK, resp)
}
