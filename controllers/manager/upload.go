package manager

import (
	"code/util"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//上传文件的控制器
type UploadController struct {
	beego.Controller
}

func (this *UploadController) Post() {
	//获取上传的文件
	_, fileheader, err := this.GetFile("editormd-image-file")
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJson()
		return
	}
	//创建保存文件的路径
	datePath := time.Now().Format("2006/01/02")
	dirPath := "./upload/" + datePath
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		beego.Error(err)
	}
	//生成文件名
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()
	fileName := strconv.FormatInt(randNum, 10) + util.RandStringRunes(16) + filepath.Ext(fileheader.Filename)
	//保存文件
	err = this.SaveToFile("editormd-image-file", dirPath+"/"+fileName)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{
			"success": 1,
			"message": "文件上传成功!",
			"url":     "/upload/" + datePath + "/" + fileName,
		}
	}
	this.ServeJson()
	return
}

/* End of file 	: upload.go */
/* Location 	:  ./controllers/upload.go */
