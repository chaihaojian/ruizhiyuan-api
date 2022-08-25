package utils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
)

func FileUpLoad(file multipart.File, header *multipart.FileHeader) (path string, err error) {
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename

	// 创建一个文件，文件名为filename，这里的返回值newFile也是一个File指针
	addr := viper.GetString("file.path") + filename
	newFile, err := os.Create(addr)
	if err != nil {
		zap.L().Error("os.Create failed", zap.Error(err))
		return "", err
	}

	defer newFile.Close()

	// 将file的内容拷贝到newFile
	_, err = io.Copy(newFile, file)
	if err != nil {
		zap.L().Error("io.Copy failed", zap.Error(err))
		return "", err
	}

	return addr, nil
}
