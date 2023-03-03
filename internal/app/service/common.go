package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/response"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

// UploadImage
// @function: UploadImage
// @description: 上传图片到本地
// @param: header *multipart.FileHeader
// @return: err error, list interface{}, total int64
func UploadImage(c *gin.Context, header *multipart.FileHeader) (file response.UploadImageResponse, err error) {
	// 文件大小限制
	if header.Size > 1024*1024*5 {
		return file, errors.New("文件大小超过限制！")
	}
	chain := GetChain(c) // 获取链
	// 读取文件后缀
	ext := strings.ToLower(path.Ext(header.Filename))
	// 拼接新文件名
	UUID := uuid.NewV4().String()
	filename := UUID + ext
	now := time.Now()
	director := global.CONFIG.Local.Path + "/" + now.Format("2006/01/")
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(director, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return file, errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := director + filename
	f, openError := header.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return file, errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return file, errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return file, errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	upload := model.Upload{Url: p, Name: filename, Key: UUID}
	if err = global.MAPDB[chain].Model(&model.Upload{}).Create(&upload).Error; err != nil {
		return file, err
	}
	file.Name = filename
	file.Url = p
	return file, err
}

func HashPending(c *gin.Context, hash string) (int, error) {
	chain := GetChain(c) // 获取链
	db := global.MAPDB[chain].Model(&model.TransHash{})
	var transHash model.TransHash
	if err := db.Unscoped().Where("hash = ?", hash).First(&transHash).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	if transHash.DeletedAt.Time.IsZero() {
		return 1, nil
	}
	return 2, nil
}

func GetChain(c *gin.Context) string {
	if c.GetString("chain") == "" {
		return global.CONFIG.Contract.DefaultNet
	}
	return c.GetString("chain")
}

func GetAddress(c *gin.Context) string {
	return c.GetString("address")
}

func GetLang(c *gin.Context) string {
	if c.GetString("lang") == "" {
		return "en"
	}
	return c.GetString("lang")
}
