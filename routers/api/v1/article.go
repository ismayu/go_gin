package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors(){
		if models.ExistTagByID(id){
			data = models.GetArticle(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _, err := range valid.Errors{
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg"  : e.GetMsg(code),
		"data" : data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg"  : e.GetMsg(code),
			"data" : data,
		})
	}
}
//新增文章
func AddArticle(c *gin.Context) {
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}

	valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不可为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或者1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors(){
		if models.ExistTagByID(tagID) {
			data := make(map[string]interface{})
			data["tag_id"] = tagID
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}else {
		for _, err := range valid.Errors{
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg"  : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context)  {
	
}

//删除文章
func DeleteArticle(c *gin.Context)  {
	
}