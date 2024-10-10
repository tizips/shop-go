package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"path/filepath"
	req "project.io/shop/web/http/request/basic"
	"project.io/shop/web/http/response/basic"
)

func DoUploadOfFile(c context.Context, ctx *app.RequestContext) {

	var request req.DoUploadOfFile

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		http.Fail(ctx, "上传失败：%v", err)
		return
	}

	fp, err := file.Open()

	if err != nil {
		http.Fail(ctx, "文件读取失败：%v", err)
		return
	}

	filename := facades.Snowflake.Generate().String() + filepath.Ext(file.Filename)
	uri := request.Dir + "/" + filename

	if err = facades.Storage.Put(uri, fp, file.Size); err != nil {
		http.Fail(ctx, "上传失败：%v", err)
		return
	}

	responses := basic.DoUploadOfFile{
		Name: filename,
		Uri:  uri,
		Url:  facades.Storage.Url(uri),
	}

	http.Success(ctx, responses)
}
