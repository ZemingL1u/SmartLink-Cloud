package apis

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
)

func (e File) SceneUploadFile(c *gin.Context) {
	e.MakeContext(c)
	path, _ := c.GetPostForm("path")
	var fileResponse FileResponse

	var done bool
	fileResponse, done = e.sceneFile(c, fileResponse, path)
	if done {
		return
	}
	e.OK(fileResponse, "上传成功")
	return
}

func (e File) sceneFile(c *gin.Context, fileResponse FileResponse, path string) (FileResponse, bool) {
	// 尝试从请求中获取名为 "file" 的文件
	file, err := c.FormFile("file")

	// 如果获取文件时出现错误（例如请求中没有文件），则返回错误信息并结束函数
	if err != nil {
		e.Error(200, errors.New(""), "文件不能为空") // 200 表示错误的HTTP状态码，但通常应为 400 或其他表示错误的代码
		return FileResponse{}, true            // 返回空的 FileResponse 和 true 表示处理完成
	}

	// 确保文件存储路径存在，如果不存在则创建
	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败") // 500 表示服务器内部错误
	}

	// 保存上传的文件到服务器
	_ = c.SaveUploadedFile(file, path)

	// 获取文件的MIME类型
	fileType, _ := utils.GetType(path)

	// 构建返回的 FileResponse 结构体，包含文件的大小、路径、完整路径、原始文件名和类型
	fileResponse = FileResponse{
		Size: pkg.GetFileSize(path), // 获取文件大小
		Path: path,                  // 文件存储路径
		Name: path,                  // 原始文件名
		Type: fileType,              // 文件类型
	}
	// 返回 fileResponse 和 false 表示处理未完成，可以继续后续操作
	return fileResponse, false
}
