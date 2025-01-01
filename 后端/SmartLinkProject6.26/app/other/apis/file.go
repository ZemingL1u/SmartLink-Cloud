package apis

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
	"github.com/google/uuid"

	"SmartLinkProject/common/file_store"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const path = "static/uploadfile/"
const devicepath = "static/deviceuploadfile/"

type File struct {
	api.Api
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e File) UploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("%s://%s/", "http", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1": // 单图
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2": // 多图
		multipartFile := e.multipleFile(c, urlPrefix)
		e.OK(multipartFile, "上传成功")
		return
	case "3": // base64
		fileResponse = e.baseImg(c, fileResponse, urlPrefix)
		e.OK(fileResponse, "上传成功")
	default:
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	}

}

func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"
	err := utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	base64File := path + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	source, _ := c.GetPostForm("source")
	err = thirdUpload(source, fileName, base64File)
	if err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}
	if source != "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	}
	return fileResponse
}

// 定义一个函数multipleFile，它接收一个*gin.Context和url前缀字符串作为参数，
// 返回一个FileResponse类型的切片。
func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	// 从gin.Context中获取multipart表单中的文件列表，该列表是通过表单字段"file"上传的。
	files := c.Request.MultipartForm.File["file"]

	// 从POST表单中获取"source"字段的值。
	source, _ := c.GetPostForm("source")

	// 初始化一个用于存放FileResponse对象的切片。
	var multipartFile []FileResponse

	// 遍历所有上传的文件。
	for _, f := range files {
		// 生成一个全局唯一标识符(guid)。
		guid := uuid.New().String()
		// 通过原文件名和guid生成新的文件名，并保留原始文件的扩展名。
		fileName := guid + utils.GetExt(f.Filename)

		// 检查文件存储路径是否存在，如果不存在则创建它。
		// utils.IsNotExistMkDir(path)的具体实现未给出，假设它创建目录并在失败时返回错误。
		err := utils.IsNotExistMkDir(path)
		if err != nil {
			// 如果初始化文件路径失败，记录一个500错误。
			e.Error(500, errors.New(""), "初始化文件路径失败")
		}

		// 构造完整的文件存储路径。
		multipartFileName := path + fileName

		// 将上传的文件保存到服务器上的multipartFileName路径。
		// c.SaveUploadedFile会保存文件，并在失败时返回错误。
		err1 := c.SaveUploadedFile(f, multipartFileName)

		// 获取文件的MIME类型。
		fileType, _ := utils.GetType(multipartFileName)

		// 如果文件保存成功，则继续处理。
		if err1 == nil {
			// 调用thirdUpload函数尝试将文件上传到第三方存储服务。
			err := thirdUpload(source, fileName, multipartFileName)
			if err != nil {
				// 如果上传到第三方失败，记录一个500错误。
				e.Error(500, errors.New(""), "上传第三方失败")
			} else {
				// 如果上传成功，创建FileResponse对象以响应客户端。
				fileResponse := FileResponse{
					Size:     pkg.GetFileSize(multipartFileName), // 获取文件大小
					Path:     multipartFileName,                  // 文件在服务器上的存储路径
					FullPath: urlPerfix + multipartFileName,      // 文件的完整URL路径
					Name:     f.Filename,                         // 原始文件名
					Type:     fileType,                           // 文件MIME类型
				}
				// 如果source参数不等于"1"，则更改Path和FullPath为静态资源路径。
				if source != "1" {
					fileResponse.Path = "/static/uploadfile/" + fileName
					fileResponse.FullPath = "/static/uploadfile/" + fileName
				}
				// 将FileResponse对象添加到切片中。
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	// 返回包含所有FileResponse对象的切片。
	return multipartFile
}

// singleFile 处理单个文件上传的逻辑
// c 是 Gin 上下文对象，用于获取请求信息和返回响应
// fileResponse 是用于存放上传文件信息的结构体实例
// urlPrefix 是请求的URL前缀，用于构建文件的完整访问路径
// 该函数返回 fileResponse 和一个布尔值，布尔值指示是否已经处理完成（例如遇到错误）
func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPrefix string) (FileResponse, bool) {
	// 尝试从请求中获取名为 "file" 的文件
	files, err := c.FormFile("file")

	// 如果获取文件时出现错误（例如请求中没有文件），则返回错误信息并结束函数
	if err != nil {
		e.Error(200, errors.New(""), "图片不能为空") // 200 表示错误的HTTP状态码，但通常应为 400 或其他表示错误的代码
		return FileResponse{}, true            // 返回空的 FileResponse 和 true 表示处理完成
	}

	// 使用 uuid 生成文件的唯一标识符，并获取原始文件的扩展名
	guid := uuid.New().String()
	fileName := guid + utils.GetExt(files.Filename)

	// 确保文件存储路径存在，如果不存在则创建
	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败") // 500 表示服务器内部错误
	}

	// 构建文件的完整存储路径
	singleFile := path + fileName

	// 保存上传的文件到服务器
	_ = c.SaveUploadedFile(files, singleFile)

	// 获取文件的MIME类型
	fileType, _ := utils.GetType(singleFile)

	// 构建返回的 FileResponse 结构体，包含文件的大小、路径、完整路径、原始文件名和类型
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(singleFile), // 获取文件大小
		Path:     singleFile,                  // 文件存储路径
		FullPath: urlPrefix + singleFile,      // 文件的完整访问路径
		Name:     files.Filename,              // 原始文件名
		Type:     fileType,                    // 文件类型
	}

	// 注释掉的代码是用于将文件上传到第三方存储服务的逻辑，当前未启用
	// source, _ := c.GetPostForm("source")
	// err = thirdUpload(source, fileName, singleFile)
	// if err != nil {
	//     e.Error(200, errors.New(""), "上传第三方失败")
	//     return FileResponse{}, true
	// }

	// 将文件路径和完整路径设置为静态资源路径，以便可以通过Web服务器直接访问
	fileResponse.Path = "/static/uploadfile/" + fileName
	fileResponse.FullPath = "/static/uploadfile/" + fileName

	// 返回 fileResponse 和 false 表示处理未完成，可以继续后续操作
	return fileResponse, false
}

func (e File) DeviceUploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("%s://%s/", "http", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1": // 单图
		var done bool
		fileResponse, done = e.devicesingleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2": // 多图
		multipartFile := e.devicemultipleFile(c, urlPrefix)
		e.OK(multipartFile, "上传成功")
		return
	case "3": // base64
		fileResponse = e.devicebaseImg(c, fileResponse, urlPrefix)
		e.OK(fileResponse, "上传成功")
	default:
		var done bool
		fileResponse, done = e.devicesingleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	}

}

func (e File) devicebaseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	deviceId, _ := c.GetPostForm("deviceId")
	fileName := deviceId + ".png"
	err := utils.IsNotExistMkDir(devicepath)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	base64File := devicepath + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	source, _ := c.GetPostForm("source")
	err = thirdUpload(source, fileName, base64File)
	if err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}
	if source != "1" {
		fileResponse.Path = "/static/deviceuploadfile/" + fileName
		fileResponse.FullPath = "/static/deviceuploadfile/" + fileName
	}
	return fileResponse
}

func (e File) devicemultipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	var multipartFile []FileResponse
	for i, f := range files { // 使用索引i
		deviceId, _ := c.GetPostForm("deviceId")
		// 将fileName设置为deviceId + 当前循环次数i + ".png"
		fileName := fmt.Sprintf("%s_%d.png", deviceId, i)

		// 其他代码保持不变...
		err := utils.IsNotExistMkDir(devicepath)
		if err != nil {
			e.Error(500, errors.New(""), "初始化文件路径失败")
		}
		multipartFileName := devicepath + deviceId + "/" + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := utils.GetType(multipartFileName)
		if err1 == nil {
			fileResponse := FileResponse{
				Size:     pkg.GetFileSize(multipartFileName),
				Path:     multipartFileName,
				FullPath: urlPerfix + multipartFileName,
				Name:     f.Filename,
				Type:     fileType,
			}
			multipartFile = append(multipartFile, fileResponse)
		}
	}
	return multipartFile
}

// singleFile 处理单个文件上传的逻辑
// c 是 Gin 上下文对象，用于获取请求信息和返回响应
// fileResponse 是用于存放上传文件信息的结构体实例
// urlPrefix 是请求的URL前缀，用于构建文件的完整访问路径
// 该函数返回 fileResponse 和一个布尔值，布尔值指示是否已经处理完成（例如遇到错误）
func (e File) devicesingleFile(c *gin.Context, fileResponse FileResponse, urlPrefix string) (FileResponse, bool) {
	// 尝试从请求中获取名为 "file" 的文件
	files, err := c.FormFile("file")

	// 如果获取文件时出现错误（例如请求中没有文件），则返回错误信息并结束函数
	if err != nil {
		e.Error(200, errors.New(""), "图片不能为空") // 200 表示错误的HTTP状态码，但通常应为 400 或其他表示错误的代码
		return FileResponse{}, true            // 返回空的 FileResponse 和 true 表示处理完成
	}

	deviceId, _ := c.GetPostForm("deviceId")
	fileName := deviceId + ".png"
	// 确保文件存储路径存在，如果不存在则创建
	err = utils.IsNotExistMkDir(devicepath)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败") // 500 表示服务器内部错误
	}

	// 构建文件的完整存储路径
	singleFile := devicepath + deviceId + "/" + fileName

	// 保存上传的文件到服务器
	_ = c.SaveUploadedFile(files, singleFile)

	// 获取文件的MIME类型
	fileType, _ := utils.GetType(singleFile)

	// 构建返回的 FileResponse 结构体，包含文件的大小、路径、完整路径、原始文件名和类型
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(singleFile), // 获取文件大小
		Path:     singleFile,                  // 文件存储路径
		FullPath: urlPrefix + singleFile,      // 文件的完整访问路径
		Name:     fileName,                    // 原始文件名
		Type:     fileType,                    // 文件类型
	}

	// 注释掉的代码是用于将文件上传到第三方存储服务的逻辑，当前未启用
	// source, _ := c.GetPostForm("source")
	// err = thirdUpload(source, fileName, singleFile)
	// if err != nil {
	//     e.Error(200, errors.New(""), "上传第三方失败")
	//     return FileResponse{}, true
	// }

	// 将文件路径和完整路径设置为静态资源路径，以便可以通过Web服务器直接访问
	fileResponse.Path = "/static/deviceuploadfile/" + fileName
	fileResponse.FullPath = "/static/deviceuploadfile/" + fileName

	// 返回 fileResponse 和 false 表示处理未完成，可以继续后续操作
	return fileResponse, false
}

// thirdUpload 根据提供的源类型，将文件上传到不同的第三方存储服务
// source 是一个字符串，表示上传到哪个第三方服务（如 "2" 表示OSS，"3" 表示七牛云）
// name 是要上传的文件名
// path 是文件在本地的路径
// 该函数返回一个错误，如果上传过程中发生错误
func thirdUpload(source string, name string, path string) error {
	switch source {
	case "2":
		// 如果 source 是 "2"，则调用 ossUpload 函数上传文件到阿里云OSS
		return ossUpload("img/"+name, path)
	case "3":
		// 如果 source 是 "3"，则调用 qiniuUpload 函数上传文件到七牛云
		return qiniuUpload("img/"+name, path)
	}
	// 如果 source 不匹配任何已知的源类型，则不执行上传，返回nil表示没有错误
	return nil
}

// ossUpload 负责将文件上传到阿里云OSS
// name 是要上传的文件名，包括OSS上的路径前缀，如 "img/somefile.jpg"
// path 是文件在本地的路径
// 该函数返回一个错误，如果上传过程中发生错误
func ossUpload(name string, path string) error {
	// 实例化一个 ALiYunOSS 对象，它是 file_store 包中定义的结构体，用于与OSS交互
	oss := file_store.ALiYunOSS{}
	// 调用 UpLoad 方法执行上传操作，传入OSS上的文件名和本地文件路径
	return oss.UpLoad(name, path)
}

// qiniuUpload 负责将文件上传到七牛云
// name 是要上传的文件名，包括七牛云上的路径前缀，如 "img/somefile.jpg"
// path 是文件在本地的路径
// 该函数返回一个错误，如果上传过程中发生错误
func qiniuUpload(name string, path string) error {
	// 实例化一个 ALiYunOSS 对象，尽管函数名是 qiniuUpload，但这里使用的是 ALiYunOSS 对象
	// 这可能是一个错误，除非 ALiYunOSS 类同时支持七牛云的上传
	oss := file_store.ALiYunOSS{}
	// 调用 UpLoad 方法执行上传操作
	return oss.UpLoad(name, path)
}
