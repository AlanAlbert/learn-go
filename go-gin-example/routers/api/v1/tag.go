package v1

import (
	"bytes"
	"github.com/anhoder/go-gin-example/models"
	"github.com/anhoder/go-gin-example/pkg/e"
	"github.com/anhoder/go-gin-example/pkg/setting"
	"github.com/anhoder/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(context *gin.Context) {
	where := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := context.Query("name"); name != "" {
		where["name"] = name
	}

	if state := context.Query("state"); state != "" {
		where["state"] = com.StrTo(state).MustInt()
	}

	data["list"] = models.GetTags(util.GetPage(context), setting.PageSize, where)
	data["total"] = models.GetTagTotal(where)

	code := e.SUCCESS

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

func AddTag(context *gin.Context) {
	name := context.PostForm("name")
	state := com.StrTo(context.DefaultPostForm("state", "0")).MustInt()
	createdBy := context.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var (
		code int
		msg string
	)
	if !valid.HasErrors() {
		if models.ExistTagByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		}
		msg = e.GetMsg(code)
	} else {
		var msgBuffer bytes.Buffer
		code = e.INVALID_PARAMS
		for index, err := range valid.Errors {
			if index != 0 {
				msgBuffer.WriteString(",")
			}
			msgBuffer.WriteString(err.Message)
		}
		msg = msgBuffer.String()
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
		"data": make(map[string]string),
	})
}

func UpdateTag(context *gin.Context) {

}

func DeleteTag(context *gin.Context) {

}
