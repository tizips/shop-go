package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/davecgh/go-spew/spew"
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
		http.Fail(ctx, "File upload failed. Please try again later.")
		return
	}

	fp, err := file.Open()

	if err != nil {
		http.Fail(ctx, "File reading failed. Please try again later.")
		return
	}

	filename := facades.Snowflake.Generate().String() + filepath.Ext(file.Filename)
	uri := request.Dir + "/" + filename

	if err = facades.Storage.Put(uri, fp, file.Size); err != nil {
		spew.Dump(err)
		http.Fail(ctx, "File upload failed. Please try again later.")
		return
	}

	responses := basic.DoUploadOfFile{
		Name: filename,
		Uri:  uri,
		Url:  facades.Storage.Url(uri),
	}

	http.Success(ctx, responses)
}
