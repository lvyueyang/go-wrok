package api

import (
	"errors"
	"fmt"
	"io"
	"path"
	"server/config"
	"server/consts"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils"
	"server/utils/resp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AdminUserController struct {
	service *service.AdminUserService
}

func NewAdminUserController(e *gin.Engine) {
	c := &AdminUserController{
		service: service.NewAdminUserService(),
	}
	admin := e.Group("/api/admin/user")

	admin.GET("/current", middleware.AdminAuth(), c.CurrentInfo)
	admin.GET("", middleware.AdminRole(utils.CreatePermission("admin:user:find:list", "查询管理员列表")), c.FindList)
	admin.POST("", middleware.AdminRole(utils.CreatePermission("admin:user:create", "创建管理员")), c.Create)
	admin.PUT("/:id", middleware.AdminRole(utils.CreatePermission("admin:user:update:info", "修改管理员基本信息")), c.Update)
	admin.DELETE("/:id", middleware.AdminRole(utils.CreatePermission("admin:user:delete", "删除管理员")), c.Delete)
	admin.PUT("/reset-password/:id", middleware.AdminRole(utils.CreatePermission("admin:user:update:password", "修改管理员密码")), c.ResetPassword)
	admin.PUT("/status/:id", middleware.AdminRole(utils.CreatePermission("admin:user:update:status", "修改管理员状态")), c.UpdateStatus)
	admin.PUT("/role", middleware.AdminRole(utils.CreatePermission("admin:user:update:role", "修改管理员角色")), c.UpdateRole)
	admin.POST("/upload", middleware.AdminRole(utils.CreatePermission("admin:user:upload:file", "上传文件到本地")), c.Upload)
}

// FindList
//
//	@Summary	管理员列表
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		current		query		number													false	"当前页"	default(1)
//	@Param		page_size	query		number													false	"每页条数"	default(20)
//	@Param		order_key	query		string													false	"需要排序的列"
//	@Param		order_type	query		string													false	"排序方式"	Enums(ase,desc)
//	@Param		keyword		query		string													false	"按用户名搜索"
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.AdminUser}}	"resp"
//	@Router		/api/admin/user [get]
func (c *AdminUserController) FindList(ctx *gin.Context) {
	query := service.FindAdminUserListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// Create
//
//	@Summary	新增管理员
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user [post]
func (c *AdminUserController) Create(ctx *gin.Context) {
	var body CreateAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Create(model.AdminUser{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
		Avatar:   body.Avatar,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Update
//
//	@Summary	修改管理员信息
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user/{id} [put]
func (c *AdminUserController) Update(ctx *gin.Context) {
	var body UpdateAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Update(uint(id), model.AdminUser{
		Name:   body.Name,
		Avatar: body.Avatar,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdateStatus
//
//	@Summary	修改管理员状态(封禁/解封)
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number							true	"管理员ID"
//	@Param		req	body		UpdateAdminUserStatusBodyDto	true	"用户状态"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/status/{id} [put]
func (c *AdminUserController) UpdateStatus(ctx *gin.Context) {
	var body UpdateAdminUserStatusBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.UpdateStatus(uint(id), body.Status); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除管理员
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number		true	"管理员ID"
//	@Success	200	{object}	resp.Result	"resp"
//	@Router		/api/admin/user/{id} [delete]
func (c *AdminUserController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// ResetPassword
//
//	@Summary	重置管理员密码
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		ResetPasswordAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/reset-password/{id} [put]
func (c *AdminUserController) ResetPassword(ctx *gin.Context) {
	var body ResetPasswordAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	currentUser := utils.GetCurrentAdminUser(ctx)

	if err := c.service.OnlyRootAdminUser(uint(id), currentUser.ID); err != nil {
		ctx.JSON(resp.ParamErr(err.Error()))
		return
	}

	if err := c.service.UpdatePassword(uint(id), body.Password); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdateRole
//
//	@Summary	为管理用户更新角色
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		AdminUserUpdateRolesBodyDto	true	"req"
//	@Success	200	{object}	resp.Result					"resp"
//	@Router		/api/admin/user/role [put]
func (c *AdminUserController) UpdateRole(ctx *gin.Context) {
	var body AdminUserUpdateRolesBodyDto

	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	if err := c.service.UpdateRole(body.UserId, body.RoleIds); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}

	ctx.JSON(resp.Succ(nil))
}

// CurrentInfo
//
//	@Summary	当前登陆者信息
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=model.AdminUser}	"用户详情"
//	@Router		/api/admin/user/current [get]
func (c *AdminUserController) CurrentInfo(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	ctx.JSON(resp.Succ(user))
}

// Upload
//
//	@Summary	文件上传
//	@Tags		管理后台-通用接口
//	@Accept		json
//	@Produce	json
//	@Param		file	formData	file						true	"文件"
//	@Success	200		{object}	resp.Result{data=string}	"文件地址"
//	@Router		/api/admin/user/upload [post]
func (c *AdminUserController) Upload(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	file, errF := ctx.FormFile("file")
	if errF != nil {
		ctx.JSON(resp.ParseErr(errF))
		return
	}
	filePath := strconv.Itoa(int(user.ID)) + "/" + file.Filename
	fmt.Println(config.Config.FileUploadDir)
	dst := path.Join(config.Config.FileUploadDir + "/" + filePath)
	fmt.Println(dst)
	// 上传文件至指定的完整文件路径
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ("/" + consts.UploadFilePathName + "/" + filePath))
}

// UploadToAliOss
//
//	@Summary	文件上传至 阿里云 oss
//	@Tags		管理后台-通用接口
//	@Accept		json
//	@Produce	json
//	@Param		file	formData	file						true	"文件"
//	@Param		prefix	formData	string						true	"文件路径"
//	@Success	200		{object}	resp.Result{data=string}	"文件地址"
//	@Router		/api/admin/user/upload-oss [post]
func (c *AdminUserController) UploadToAliOss(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	file, errF := ctx.FormFile("file")
	prefix := strings.TrimSpace(ctx.PostForm("prefix"))
	if len(prefix) < 3 {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(errors.New("prefix 长度不能小于 3"))))
		return
	}
	if prefix[0:1] == "/" {
		prefix = prefix[1:]
	}
	if prefix[len(prefix)-1:] == "/" {
		prefix = prefix[0 : len(prefix)-1]
	}
	if errF != nil {
		ctx.JSON(resp.ParseErr(errF))
		return
	}
	envPath := "test"
	if config.Config.IsProd {
		envPath = "prod"
	}
	date := time.Now().Format("2006/01/02")
	filePath := envPath + "/file/" + strconv.Itoa(int(user.ID)) + "/" + prefix + "/" + date + "/" + file.Filename
	ossConfig := config.Config.AliOss
	fmt.Printf("ossconfig %+v \n", ossConfig)
	f, errO := file.Open()
	if errF != nil {
		ctx.JSON(resp.ParseErr(errO))
		return
	}
	errU := utils.UploadFileToAliOss(utils.UploadOssOptions{
		AccessKeyID:     ossConfig.AccessKeyID,
		AccessKeySecret: ossConfig.AccessKeySecret,
		Endpoint:        ossConfig.Endpoint,
		Bucket:          ossConfig.Bucket,
		File:            io.Reader(f),
		FilePath:        path.Join(filePath),
		Options:         nil,
	})
	if errU != nil {
		ctx.JSON(resp.ParseErr(errU))
		return
	}
	ctx.JSON(resp.Succ("https://" + ossConfig.Bucket + "." + ossConfig.Endpoint + "/" + filePath))
}

type CreateAdminUserBodyDto struct {
	Name     string `json:"name" binding:"required" label:"姓名"`      // 姓名
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
	Avatar   string `json:"avatar"`                                  // 头像
}

type UpdateAdminUserBodyDto struct {
	Name   string `json:"name" binding:"required" label:"姓名"` // 姓名
	Avatar string `json:"avatar"`                             // 头像
}

type UpdateAdminUserStatusBodyDto struct {
	Status consts.AdminUserStatus `json:"status" binding:"required" label:"用户状态" enums:"-1,1"` // 状态 1-解封 2-封禁
}

type DeleteAdminUserStatusBodyDto struct {
	ID uint `json:"id" binding:"required"`
}
type ResetPasswordAdminUserBodyDto struct {
	Password string `json:"password" binding:"required" label:"密码"` // 密码
}

type AdminUserUpdateRolesBodyDto struct {
	UserId  uint   `json:"user_id" binding:"required" label:"用户 ID"`  // 用户 ID
	RoleIds []uint `json:"role_ids" binding:"required" label:"角色 ID"` // 角色 ID
}
