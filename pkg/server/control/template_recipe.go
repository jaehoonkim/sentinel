package control

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/NexClipper/sudory/pkg/server/database/vanilla"
	"github.com/NexClipper/sudory/pkg/server/database/vanilla/prepare"
	"github.com/NexClipper/sudory/pkg/server/macro/echoutil"
	recipev2 "github.com/NexClipper/sudory/pkg/server/model/template_recipe/v2"
	"github.com/NexClipper/sudory/pkg/server/status/state"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// @Description Find TemplateRecipe
// @Produce     json
// @Tags        server/template_recipe
// @Router      /server/template_recipe [get]
// @Param       x_auth_token header string false "client session token"
// @Param       method       query  string false "Template Command Method"
// @Success     200 {array} v2.HttpRsp_TemplateRecipe
func (ctl ControlVanilla) FindTemplateRecipe(ctx echo.Context) (err error) {
	method := echoutil.QueryParam(ctx)["method"]
	buff := bytes.Buffer{}
	for i, s := range strings.Split(method, ".") {
		if 0 < i {
			buff.WriteString(".")
		}
		buff.WriteString(s)
	}
	//뒤에 like 조회 와일드 카드를 붙여준다
	buff.WriteString("%")

	var p *prepare.Pagination = vanilla.Limit(50, 1).Parse()
	if 0 < len(echoutil.QueryParam(ctx)["p"]) {
		p, err = prepare.NewPagination(echoutil.QueryParam(ctx)["p"])
		err = errors.Wrapf(err, "failed to parse pagination")
		if err != nil {
			return HttpError(err, http.StatusBadRequest)
		}
	}

	rsp := make([]recipev2.HttpRsp_TemplateRecipe, 0, state.ENV__INIT_SLICE_CAPACITY__())

	recipe := recipev2.TemplateRecipe{}
	recipe.Method = buff.String()
	like_method := vanilla.Like("method", recipe.Method)
	order := vanilla.Asc("name", "args")

	err = vanilla.Stmt.Select(recipe.TableName(), recipe.ColumnNames(), like_method.Parse(), order.Parse(), p).
		QueryRows(ctl)(func(scan vanilla.Scanner, _ int) (err error) {
		err = recipe.Scan(scan)
		if err != nil {
			return errors.Wrapf(err, "failed to scan")
		}
		rsp = append(rsp, recipev2.HttpRsp_TemplateRecipe(recipe))
		return
	})
	if err != nil {
		return
	}

	return ctx.JSON(http.StatusOK, rsp)
}
