// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"context"
	"errors"
	"github.com/goplus/builder/spx-backend/internal/controller"
	"github.com/goplus/builder/spx-backend/internal/log"
	"github.com/goplus/builder/spx-backend/internal/model"
	"github.com/goplus/yap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const _ = true
const (
	firstPageIndex  = 1
	defaultPageSize = 10
)

type delete_asset_id struct {
	yap.Handler
	*AppV2
}
type delete_project_owner_name struct {
	yap.Handler
	*AppV2
}
type get_asset_id struct {
	yap.Handler
	*AppV2
}
type get_assets_list struct {
	yap.Handler
	*AppV2
}
type get_project_owner_name struct {
	yap.Handler
	*AppV2
}
type get_projects_list struct {
	yap.Handler
	*AppV2
}
type get_util_upinfo struct {
	yap.Handler
	*AppV2
}
type AppV2 struct {
	yap.AppV2
	ctrl *controller.Controller
	err  error
}
type post_asset struct {
	yap.Handler
	*AppV2
}
type post_asset_id_click struct {
	yap.Handler
	*AppV2
}
type post_project struct {
	yap.Handler
	*AppV2
}
type post_util_fileurls struct {
	yap.Handler
	*AppV2
}
type post_util_fmtcode struct {
	yap.Handler
	*AppV2
}
type put_asset_id struct {
	yap.Handler
	*AppV2
}
type put_project_owner_name struct {
	yap.Handler
	*AppV2
}
//line cmd/spx-backend/main.yap:26
func (this *AppV2) MainEntry() {
//line cmd/spx-backend/main.yap:26:1
	logger := log.GetLogger()
//line cmd/spx-backend/main.yap:28:1
	this.ctrl, this.err = controller.New(context.Background())
//line cmd/spx-backend/main.yap:29:1
	if this.err != nil {
//line cmd/spx-backend/main.yap:30:1
		logger.Fatalln("Failed to create a new controller:", this.err)
	}
//line cmd/spx-backend/main.yap:33:1
	port := os.Getenv("PORT")
//line cmd/spx-backend/main.yap:34:1
	if port == "" {
//line cmd/spx-backend/main.yap:35:1
		port = ":8080"
	}
//line cmd/spx-backend/main.yap:37:1
	logger.Printf("Listening to %s", port)
//line cmd/spx-backend/main.yap:39:1
	h := this.Handler(NewUserMiddleware(this.ctrl), NewReqIDMiddleware(), NewCORSMiddleware())
//line cmd/spx-backend/main.yap:40:1
	server := &http.Server{Addr: port, Handler: h}
//line cmd/spx-backend/main.yap:42:1
	stopCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
//line cmd/spx-backend/main.yap:43:1
	defer stop()
//line cmd/spx-backend/main.yap:44:1
	var serverErr error
//line cmd/spx-backend/main.yap:45:1
	go func() {
//line cmd/spx-backend/main.yap:46:1
		serverErr = server.ListenAndServe()
//line cmd/spx-backend/main.yap:47:1
		stop()
	}()
//line cmd/spx-backend/main.yap:49:1
	<-stopCtx.Done()
//line cmd/spx-backend/main.yap:50:1
	if serverErr != nil && !errors.Is(serverErr, http.ErrServerClosed) {
//line cmd/spx-backend/main.yap:51:1
		logger.Fatalln("Server error:", this.err)
	}
//line cmd/spx-backend/main.yap:54:1
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
//line cmd/spx-backend/main.yap:55:1
	defer cancel()
//line cmd/spx-backend/main.yap:56:1
	if
//line cmd/spx-backend/main.yap:56:1
	err := server.Shutdown(shutdownCtx); err != nil {
//line cmd/spx-backend/main.yap:57:1
		logger.Fatalln("Failed to gracefully shut down:", err)
	}
}
func (this *AppV2) Main() {
	yap.Gopt_AppV2_Main(this, new(delete_asset_id), new(delete_project_owner_name), new(get_asset_id), new(get_assets_list), new(get_project_owner_name), new(get_projects_list), new(get_util_upinfo), new(post_asset), new(post_asset_id_click), new(post_project), new(post_util_fileurls), new(post_util_fmtcode), new(put_asset_id), new(put_project_owner_name))
}
//line cmd/spx-backend/delete_asset_#id.yap:6
func (this *delete_asset_id) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/delete_asset_#id.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/delete_asset_#id.yap:8:1
	if
//line cmd/spx-backend/delete_asset_#id.yap:8:1
	err := this.ctrl.DeleteAsset(ctx.Context(), this.Gop_Env("id")); err != nil {
//line cmd/spx-backend/delete_asset_#id.yap:9:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/delete_asset_#id.yap:10:1
		return
	}
//line cmd/spx-backend/delete_asset_#id.yap:12:1
	this.Json__1(nil)
}
func (this *delete_asset_id) Classfname() string {
	return "delete_asset_#id"
}
//line cmd/spx-backend/delete_project_#owner_#name.yap:6
func (this *delete_project_owner_name) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/delete_project_#owner_#name.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/delete_project_#owner_#name.yap:8:1
	if
//line cmd/spx-backend/delete_project_#owner_#name.yap:8:1
	err := this.ctrl.DeleteProject(ctx.Context(), this.Gop_Env("owner"), this.Gop_Env("name")); err != nil {
//line cmd/spx-backend/delete_project_#owner_#name.yap:9:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/delete_project_#owner_#name.yap:10:1
		return
	}
//line cmd/spx-backend/delete_project_#owner_#name.yap:12:1
	this.Json__1(nil)
}
func (this *delete_project_owner_name) Classfname() string {
	return "delete_project_#owner_#name"
}
//line cmd/spx-backend/get_asset_#id.yap:6
func (this *get_asset_id) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/get_asset_#id.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/get_asset_#id.yap:8:1
	asset, err := this.ctrl.GetAsset(ctx.Context(), this.Gop_Env("id"))
//line cmd/spx-backend/get_asset_#id.yap:9:1
	if err != nil {
//line cmd/spx-backend/get_asset_#id.yap:10:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/get_asset_#id.yap:11:1
		return
	}
//line cmd/spx-backend/get_asset_#id.yap:13:1
	this.Json__1(asset)
}
func (this *get_asset_id) Classfname() string {
	return "get_asset_#id"
}
//line cmd/spx-backend/get_assets_list.yap:13
func (this *get_assets_list) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/get_assets_list.yap:13:1
	ctx := &this.Context
//line cmd/spx-backend/get_assets_list.yap:15:1
	user, _ := controller.UserFromContext(ctx.Context())
//line cmd/spx-backend/get_assets_list.yap:16:1
	params := &controller.ListAssetsParams{}
//line cmd/spx-backend/get_assets_list.yap:18:1
	params.Keyword = this.Gop_Env("keyword")
//line cmd/spx-backend/get_assets_list.yap:20:1
	switch
//line cmd/spx-backend/get_assets_list.yap:20:1
	owner := this.Gop_Env("owner"); owner {
//line cmd/spx-backend/get_assets_list.yap:21:1
	case "":
//line cmd/spx-backend/get_assets_list.yap:22:1
		if user == nil {
//line cmd/spx-backend/get_assets_list.yap:23:1
			replyWithCode(ctx, errorUnauthorized)
//line cmd/spx-backend/get_assets_list.yap:24:1
			return
		}
//line cmd/spx-backend/get_assets_list.yap:26:1
		params.Owner = &user.Name
//line cmd/spx-backend/get_assets_list.yap:27:1
	case "*":
//line cmd/spx-backend/get_assets_list.yap:28:1
		params.Owner = nil
//line cmd/spx-backend/get_assets_list.yap:29:1
	default:
//line cmd/spx-backend/get_assets_list.yap:30:1
		params.Owner = &owner
	}
//line cmd/spx-backend/get_assets_list.yap:33:1
	if
//line cmd/spx-backend/get_assets_list.yap:33:1
	category := this.Gop_Env("category"); category != "" {
//line cmd/spx-backend/get_assets_list.yap:34:1
		params.Category = &category
	}
//line cmd/spx-backend/get_assets_list.yap:37:1
	if
//line cmd/spx-backend/get_assets_list.yap:37:1
	assetTypeParam := this.Gop_Env("assetType"); assetTypeParam != "" {
//line cmd/spx-backend/get_assets_list.yap:38:1
		assetTypeInt, err := strconv.Atoi(assetTypeParam)
//line cmd/spx-backend/get_assets_list.yap:39:1
		if err != nil {
//line cmd/spx-backend/get_assets_list.yap:40:1
			replyWithCode(ctx, errorInvalidArgs)
//line cmd/spx-backend/get_assets_list.yap:41:1
			return
		}
//line cmd/spx-backend/get_assets_list.yap:43:1
		assetType := model.AssetType(assetTypeInt)
//line cmd/spx-backend/get_assets_list.yap:44:1
		params.AssetType = &assetType
	}
//line cmd/spx-backend/get_assets_list.yap:47:1
	if
//line cmd/spx-backend/get_assets_list.yap:47:1
	isPublicParam := this.Gop_Env("isPublic"); isPublicParam != "" {
//line cmd/spx-backend/get_assets_list.yap:48:1
		isPublicInt, err := strconv.Atoi(isPublicParam)
//line cmd/spx-backend/get_assets_list.yap:49:1
		if err != nil {
//line cmd/spx-backend/get_assets_list.yap:50:1
			replyWithCode(ctx, errorInvalidArgs)
//line cmd/spx-backend/get_assets_list.yap:51:1
			return
		}
//line cmd/spx-backend/get_assets_list.yap:53:1
		isPublic := model.IsPublic(isPublicInt)
//line cmd/spx-backend/get_assets_list.yap:54:1
		params.IsPublic = &isPublic
	}
//line cmd/spx-backend/get_assets_list.yap:57:1
	if
//line cmd/spx-backend/get_assets_list.yap:57:1
	orderBy := this.Gop_Env("orderBy"); orderBy != "" {
//line cmd/spx-backend/get_assets_list.yap:58:1
		params.OrderBy = controller.ListAssetsOrderBy(orderBy)
	}
//line cmd/spx-backend/get_assets_list.yap:61:1
	params.Pagination.Index = ctx.ParamInt("pageIndex", firstPageIndex)
//line cmd/spx-backend/get_assets_list.yap:62:1
	params.Pagination.Size = ctx.ParamInt("pageSize", defaultPageSize)
//line cmd/spx-backend/get_assets_list.yap:63:1
	if
//line cmd/spx-backend/get_assets_list.yap:63:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/get_assets_list.yap:64:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/get_assets_list.yap:65:1
		return
	}
//line cmd/spx-backend/get_assets_list.yap:68:1
	assets, err := this.ctrl.ListAssets(ctx.Context(), params)
//line cmd/spx-backend/get_assets_list.yap:69:1
	if err != nil {
//line cmd/spx-backend/get_assets_list.yap:70:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/get_assets_list.yap:71:1
		return
	}
//line cmd/spx-backend/get_assets_list.yap:73:1
	this.Json__1(assets)
}
func (this *get_assets_list) Classfname() string {
	return "get_assets_list"
}
//line cmd/spx-backend/get_project_#owner_#name.yap:6
func (this *get_project_owner_name) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/get_project_#owner_#name.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/get_project_#owner_#name.yap:8:1
	project, err := this.ctrl.GetProject(ctx.Context(), this.Gop_Env("owner"), this.Gop_Env("name"))
//line cmd/spx-backend/get_project_#owner_#name.yap:9:1
	if err != nil {
//line cmd/spx-backend/get_project_#owner_#name.yap:10:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/get_project_#owner_#name.yap:11:1
		return
	}
//line cmd/spx-backend/get_project_#owner_#name.yap:13:1
	this.Json__1(project)
}
func (this *get_project_owner_name) Classfname() string {
	return "get_project_#owner_#name"
}
//line cmd/spx-backend/get_projects_list.yap:13
func (this *get_projects_list) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/get_projects_list.yap:13:1
	ctx := &this.Context
//line cmd/spx-backend/get_projects_list.yap:15:1
	user, _ := controller.UserFromContext(ctx.Context())
//line cmd/spx-backend/get_projects_list.yap:16:1
	params := &controller.ListProjectsParams{}
//line cmd/spx-backend/get_projects_list.yap:18:1
	if
//line cmd/spx-backend/get_projects_list.yap:18:1
	isPublicParam := this.Gop_Env("isPublic"); isPublicParam != "" {
//line cmd/spx-backend/get_projects_list.yap:19:1
		isPublicInt, err := strconv.Atoi(isPublicParam)
//line cmd/spx-backend/get_projects_list.yap:20:1
		if err != nil {
//line cmd/spx-backend/get_projects_list.yap:21:1
			replyWithCode(ctx, errorInvalidArgs)
//line cmd/spx-backend/get_projects_list.yap:22:1
			return
		}
//line cmd/spx-backend/get_projects_list.yap:24:1
		isPublic := model.IsPublic(isPublicInt)
//line cmd/spx-backend/get_projects_list.yap:25:1
		params.IsPublic = &isPublic
	}
//line cmd/spx-backend/get_projects_list.yap:28:1
	switch
//line cmd/spx-backend/get_projects_list.yap:28:1
	owner := this.Gop_Env("owner"); owner {
//line cmd/spx-backend/get_projects_list.yap:29:1
	case "":
//line cmd/spx-backend/get_projects_list.yap:30:1
		if user == nil {
//line cmd/spx-backend/get_projects_list.yap:31:1
			replyWithCode(ctx, errorUnauthorized)
//line cmd/spx-backend/get_projects_list.yap:32:1
			return
		}
//line cmd/spx-backend/get_projects_list.yap:34:1
		params.Owner = &user.Name
//line cmd/spx-backend/get_projects_list.yap:35:1
	case "*":
//line cmd/spx-backend/get_projects_list.yap:36:1
		params.Owner = nil
//line cmd/spx-backend/get_projects_list.yap:37:1
	default:
//line cmd/spx-backend/get_projects_list.yap:38:1
		params.Owner = &owner
	}
//line cmd/spx-backend/get_projects_list.yap:41:1
	params.Pagination.Index = this.ParamInt("pageIndex", firstPageIndex)
//line cmd/spx-backend/get_projects_list.yap:42:1
	params.Pagination.Size = this.ParamInt("pageSize", defaultPageSize)
//line cmd/spx-backend/get_projects_list.yap:43:1
	if
//line cmd/spx-backend/get_projects_list.yap:43:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/get_projects_list.yap:44:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/get_projects_list.yap:45:1
		return
	}
//line cmd/spx-backend/get_projects_list.yap:48:1
	projects, err := this.ctrl.ListProjects(ctx.Context(), params)
//line cmd/spx-backend/get_projects_list.yap:49:1
	if err != nil {
//line cmd/spx-backend/get_projects_list.yap:50:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/get_projects_list.yap:51:1
		return
	}
//line cmd/spx-backend/get_projects_list.yap:53:1
	this.Json__1(projects)
}
func (this *get_projects_list) Classfname() string {
	return "get_projects_list"
}
//line cmd/spx-backend/get_util_upinfo.yap:6
func (this *get_util_upinfo) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/get_util_upinfo.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/get_util_upinfo.yap:8:1
	if
//line cmd/spx-backend/get_util_upinfo.yap:8:1
	_, ok := ensureUser(ctx); !ok {
//line cmd/spx-backend/get_util_upinfo.yap:9:1
		return
	}
//line cmd/spx-backend/get_util_upinfo.yap:11:1
	upInfo, err := this.ctrl.GetUpInfo(ctx.Context())
//line cmd/spx-backend/get_util_upinfo.yap:12:1
	if err != nil {
//line cmd/spx-backend/get_util_upinfo.yap:13:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/get_util_upinfo.yap:14:1
		return
	}
//line cmd/spx-backend/get_util_upinfo.yap:16:1
	this.Json__1(upInfo)
}
func (this *get_util_upinfo) Classfname() string {
	return "get_util_upinfo"
}
//line cmd/spx-backend/post_asset.yap:10
func (this *post_asset) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/post_asset.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/post_asset.yap:12:1
	user, ok := ensureUser(ctx)
//line cmd/spx-backend/post_asset.yap:13:1
	if !ok {
//line cmd/spx-backend/post_asset.yap:14:1
		return
	}
//line cmd/spx-backend/post_asset.yap:17:1
	params := &controller.AddAssetParams{}
//line cmd/spx-backend/post_asset.yap:18:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/post_asset.yap:19:1
		return
	}
//line cmd/spx-backend/post_asset.yap:21:1
	params.Owner = user.Name
//line cmd/spx-backend/post_asset.yap:22:1
	if
//line cmd/spx-backend/post_asset.yap:22:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/post_asset.yap:23:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/post_asset.yap:24:1
		return
	}
//line cmd/spx-backend/post_asset.yap:27:1
	asset, err := this.ctrl.AddAsset(ctx.Context(), params)
//line cmd/spx-backend/post_asset.yap:28:1
	if err != nil {
//line cmd/spx-backend/post_asset.yap:29:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/post_asset.yap:30:1
		return
	}
//line cmd/spx-backend/post_asset.yap:32:1
	this.Json__1(asset)
}
func (this *post_asset) Classfname() string {
	return "post_asset"
}
//line cmd/spx-backend/post_asset_#id_click.yap:6
func (this *post_asset_id_click) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/post_asset_#id_click.yap:6:1
	ctx := &this.Context
//line cmd/spx-backend/post_asset_#id_click.yap:8:1
	if
//line cmd/spx-backend/post_asset_#id_click.yap:8:1
	err := this.ctrl.IncreaseAssetClickCount(ctx.Context(), this.Gop_Env("id")); err != nil {
//line cmd/spx-backend/post_asset_#id_click.yap:9:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/post_asset_#id_click.yap:10:1
		return
	}
//line cmd/spx-backend/post_asset_#id_click.yap:12:1
	this.Json__1(nil)
}
func (this *post_asset_id_click) Classfname() string {
	return "post_asset_#id_click"
}
//line cmd/spx-backend/post_project.yap:10
func (this *post_project) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/post_project.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/post_project.yap:12:1
	user, ok := ensureUser(ctx)
//line cmd/spx-backend/post_project.yap:13:1
	if !ok {
//line cmd/spx-backend/post_project.yap:14:1
		return
	}
//line cmd/spx-backend/post_project.yap:17:1
	params := &controller.AddProjectParams{}
//line cmd/spx-backend/post_project.yap:18:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/post_project.yap:19:1
		return
	}
//line cmd/spx-backend/post_project.yap:21:1
	params.Owner = user.Name
//line cmd/spx-backend/post_project.yap:22:1
	if
//line cmd/spx-backend/post_project.yap:22:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/post_project.yap:23:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/post_project.yap:24:1
		return
	}
//line cmd/spx-backend/post_project.yap:27:1
	project, err := this.ctrl.AddProject(ctx.Context(), params)
//line cmd/spx-backend/post_project.yap:28:1
	if err != nil {
//line cmd/spx-backend/post_project.yap:29:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/post_project.yap:30:1
		return
	}
//line cmd/spx-backend/post_project.yap:32:1
	this.Json__1(project)
}
func (this *post_project) Classfname() string {
	return "post_project"
}
//line cmd/spx-backend/post_util_fileurls.yap:10
func (this *post_util_fileurls) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/post_util_fileurls.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/post_util_fileurls.yap:12:1
	if
//line cmd/spx-backend/post_util_fileurls.yap:12:1
	_, ok := ensureUser(ctx); !ok {
//line cmd/spx-backend/post_util_fileurls.yap:13:1
		return
	}
//line cmd/spx-backend/post_util_fileurls.yap:16:1
	params := &controller.MakeFileURLsParams{}
//line cmd/spx-backend/post_util_fileurls.yap:17:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/post_util_fileurls.yap:18:1
		return
	}
//line cmd/spx-backend/post_util_fileurls.yap:20:1
	if
//line cmd/spx-backend/post_util_fileurls.yap:20:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/post_util_fileurls.yap:21:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/post_util_fileurls.yap:22:1
		return
	}
//line cmd/spx-backend/post_util_fileurls.yap:25:1
	fileURLs, err := this.ctrl.MakeFileURLs(ctx.Context(), params)
//line cmd/spx-backend/post_util_fileurls.yap:26:1
	if err != nil {
//line cmd/spx-backend/post_util_fileurls.yap:27:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/post_util_fileurls.yap:28:1
		return
	}
//line cmd/spx-backend/post_util_fileurls.yap:30:1
	this.Json__1(fileURLs)
}
func (this *post_util_fileurls) Classfname() string {
	return "post_util_fileurls"
}
//line cmd/spx-backend/post_util_fmtcode.yap:10
func (this *post_util_fmtcode) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/post_util_fmtcode.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/post_util_fmtcode.yap:12:1
	params := &controller.FmtCodeParams{}
//line cmd/spx-backend/post_util_fmtcode.yap:13:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/post_util_fmtcode.yap:14:1
		return
	}
//line cmd/spx-backend/post_util_fmtcode.yap:16:1
	if
//line cmd/spx-backend/post_util_fmtcode.yap:16:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/post_util_fmtcode.yap:17:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/post_util_fmtcode.yap:18:1
		return
	}
//line cmd/spx-backend/post_util_fmtcode.yap:21:1
	formattedCode, err := this.ctrl.FmtCode(ctx.Context(), params)
//line cmd/spx-backend/post_util_fmtcode.yap:22:1
	if err != nil {
//line cmd/spx-backend/post_util_fmtcode.yap:23:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/post_util_fmtcode.yap:24:1
		return
	}
//line cmd/spx-backend/post_util_fmtcode.yap:26:1
	this.Json__1(formattedCode)
}
func (this *post_util_fmtcode) Classfname() string {
	return "post_util_fmtcode"
}
//line cmd/spx-backend/put_asset_#id.yap:10
func (this *put_asset_id) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/put_asset_#id.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/put_asset_#id.yap:12:1
	if
//line cmd/spx-backend/put_asset_#id.yap:12:1
	_, ok := ensureUser(ctx); !ok {
//line cmd/spx-backend/put_asset_#id.yap:13:1
		return
	}
//line cmd/spx-backend/put_asset_#id.yap:16:1
	params := &controller.UpdateAssetParams{}
//line cmd/spx-backend/put_asset_#id.yap:17:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/put_asset_#id.yap:18:1
		return
	}
//line cmd/spx-backend/put_asset_#id.yap:20:1
	if
//line cmd/spx-backend/put_asset_#id.yap:20:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/put_asset_#id.yap:21:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/put_asset_#id.yap:22:1
		return
	}
//line cmd/spx-backend/put_asset_#id.yap:25:1
	asset, err := this.ctrl.UpdateAsset(ctx.Context(), this.Gop_Env("id"), params)
//line cmd/spx-backend/put_asset_#id.yap:26:1
	if err != nil {
//line cmd/spx-backend/put_asset_#id.yap:27:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/put_asset_#id.yap:28:1
		return
	}
//line cmd/spx-backend/put_asset_#id.yap:30:1
	this.Json__1(asset)
}
func (this *put_asset_id) Classfname() string {
	return "put_asset_#id"
}
//line cmd/spx-backend/put_project_#owner_#name.yap:10
func (this *put_project_owner_name) Main(_gop_arg0 *yap.Context) {
	this.Handler.Main(_gop_arg0)
//line cmd/spx-backend/put_project_#owner_#name.yap:10:1
	ctx := &this.Context
//line cmd/spx-backend/put_project_#owner_#name.yap:12:1
	if
//line cmd/spx-backend/put_project_#owner_#name.yap:12:1
	_, ok := ensureUser(ctx); !ok {
//line cmd/spx-backend/put_project_#owner_#name.yap:13:1
		return
	}
//line cmd/spx-backend/put_project_#owner_#name.yap:16:1
	params := &controller.UpdateProjectParams{}
//line cmd/spx-backend/put_project_#owner_#name.yap:17:1
	if !parseJSON(ctx, params) {
//line cmd/spx-backend/put_project_#owner_#name.yap:18:1
		return
	}
//line cmd/spx-backend/put_project_#owner_#name.yap:20:1
	if
//line cmd/spx-backend/put_project_#owner_#name.yap:20:1
	ok, msg := params.Validate(); !ok {
//line cmd/spx-backend/put_project_#owner_#name.yap:21:1
		replyWithCodeMsg(ctx, errorInvalidArgs, msg)
//line cmd/spx-backend/put_project_#owner_#name.yap:22:1
		return
	}
//line cmd/spx-backend/put_project_#owner_#name.yap:25:1
	project, err := this.ctrl.UpdateProject(ctx.Context(), this.Gop_Env("owner"), this.Gop_Env("name"), params)
//line cmd/spx-backend/put_project_#owner_#name.yap:26:1
	if err != nil {
//line cmd/spx-backend/put_project_#owner_#name.yap:27:1
		replyWithInnerError(ctx, err)
//line cmd/spx-backend/put_project_#owner_#name.yap:28:1
		return
	}
//line cmd/spx-backend/put_project_#owner_#name.yap:30:1
	this.Json__1(project)
}
func (this *put_project_owner_name) Classfname() string {
	return "put_project_#owner_#name"
}
func main() {
	new(AppV2).Main()
}
