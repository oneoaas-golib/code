package manager

import (
	"code/util"
	"github.com/astaxie/beego"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Post() {
	//获取上传的文件
	_, fileheader, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}
	//生成文件名
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()
	fileName := util.RandStringRunes(8) + strconv.FormatInt(randNum, 10) + filepath.Ext(fileheader.Filename)
	//保存文件
	err = this.SaveToFile("attachment", "./tmp/"+fileName)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": "success",
			"info": "文件上传成功!",
			"data": map[string]string{"filename": fileName},
		}
	}
	this.ServeJson()
	return
}
