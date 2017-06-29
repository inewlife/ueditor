package controllers

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
)

const ueBasePath = `assets/upload/ueditor/`

type UeditorController struct {
	beego.Controller
}

func (c *UeditorController) Get() {
	c.TplName = "ueditor.html"
}
func (c *UeditorController) Action() {
	action := c.GetString("action")
	datePath := time.Now().Format("20060102") + `/`
	switch action {
	case "config":
		jsonByte, _ := ioutil.ReadFile("conf/ueditor.json")
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
		jsonByte = re.ReplaceAll(jsonByte, []byte(""))
		c.Ctx.ResponseWriter.Write(jsonByte)
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	case "uploadimage":
		{
			err := os.MkdirAll(ueBasePath+datePath, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			_, h, err := c.GetFile("upfile")
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			name := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(h.Filename)
			path := ueBasePath + `image/` + datePath
			err = os.MkdirAll(path, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			err = c.SaveToFile("upfile", path+name) //.Join("uploadfile", uploadfile)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			c.Data["json"] = map[string]interface{}{
				"state":    "SUCCESS",
				"url":      `/` + path + name,
				"title":    h.Filename,
				"original": h.Filename,
			}
			c.ServeJSON()
			return
		}
	case "uploadvideo":
		{
			err := os.MkdirAll(ueBasePath+datePath, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			_, h, err := c.GetFile("upfile")
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			name := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(h.Filename)
			path := ueBasePath + `video/` + datePath
			err = os.MkdirAll(path, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			err = c.SaveToFile("upfile", path+name) //.Join("uploadfile", uploadfile)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			c.Data["json"] = map[string]interface{}{
				"state": "SUCCESS",
				"url":   `/` + path + name,
			}
			c.ServeJSON()
			return
		}
	case "uploadfile":
		{
			err := os.MkdirAll(ueBasePath+datePath, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			_, h, err := c.GetFile("upfile")
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			name := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(h.Filename)
			path := ueBasePath + `file/` + datePath
			err = os.MkdirAll(path, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			err = c.SaveToFile("upfile", path+name) //.Join("uploadfile", uploadfile)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			c.Data["json"] = map[string]interface{}{
				"state":    "SUCCESS",
				"url":      `/` + path + name,
				"title":    h.Filename,
				"original": h.Filename,
			}
			c.ServeJSON()
			return
		}
	case "uploadscrawl":
		{
			name := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + `.jpg`
			path := ueBasePath + `scrawl/` + datePath
			err := os.MkdirAll(path, 0777)
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			//upfile为base64格式文件，转成图片保存
			upfile := c.Input().Get("upfile")
			upBytes, err := base64.StdEncoding.DecodeString(upfile) // + "_" + filename
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			err = ioutil.WriteFile(path+name, upBytes, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			if err != nil {
				c.Data["json"] = map[string]interface{}{
					"state": err.Error(),
				}
				c.ServeJSON()
				return
			}
			c.Data["json"] = map[string]interface{}{
				"state":    "SUCCESS",
				"url":      `/` + path + name,
				"title":    `涂鸦`,
				"original": `涂鸦不见了`,
			}
			c.ServeJSON()
			return
		}
	}
}
