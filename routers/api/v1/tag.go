package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context){
	name := c.Query("name")
	//c.Query()可以用于获取?name=test&state=1这类URL参数
	//而c.DefaultQuery()则支持设置一个默认值
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != ""{
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	//until.GetPage()保证各接口的page处理是一致的

	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不可为空")
	valid.MaxSize(name, 100, "name").Message("名称最长100个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人名称最长100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors(){
		if ! models.ExistTagByName(name){
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}
func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1,"state").Message("状态只能为0或者1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100个字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id){
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != ""{
				data["name"] = name
			}
			if state != -1{
				data["state"] = state
			}
			models.EditTag(id, data)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"date" : make(map[string]string),
	})
}

func DeleteTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors(){
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})


}