package utils

import (
	"context"
	"encoding/base64"
	uuid2 "github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
)

func UploadAvatar(avatar_base64 *string) (string, error) {

	if *avatar_base64 == "" {
		return "https://blog-1301138214.cos.ap-guangzhou.myqcloud.com/avatar.png", nil
	}
	decodeString, err := base64.StdEncoding.DecodeString(*avatar_base64)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	uuid := uuid2.New().String()
	fillPath := "image/" + uuid + ".png"
	f, err := os.OpenFile(fillPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer f.Close()
	f.Write(decodeString)
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	u, _ := url.Parse("https://blog-1301138214.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "",
			SecretKey: "",
		},
	})

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "avatar/" + uuid + ".png"
	// 2.通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), name, fillPath, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "https://blog-1301138214.cos.ap-guangzhou.myqcloud.com/"+name, nil
}
