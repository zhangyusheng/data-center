package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/zhangyusheng/data-center/pkg/app"
	"github.com/zhangyusheng/data-center/pkg/e"
	"github.com/zhangyusheng/data-center/pkg/logging"
)

type AddTagForm struct {
	Name      string `form:"name" valid:"Required"`
	CreatedBy string `form:"created_by" valid:"Required"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/test [post]
func Test(c *gin.Context) {
	logging.Info(map[string]interface{}{"name":"zhangyusheng"}, "good")
	var (
		appG = app.Gin{C: c}
	)
	appG.Response(http.StatusOK, e.SUCCESS, nil)


}

