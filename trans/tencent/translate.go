package trans

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

var keyPath = "config.toml"

type TencentClient struct {
	client *tmt.Client
}

func init() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	exeDir := filepath.Dir(exePath)
	keyPath = path.Join(exeDir, keyPath)
}

type TencentKey struct {
	SecretID  string `toml:"secretID"`
	SecretKey string `toml:"secretKey"`
}

type Config struct {
	TencentKey TencentKey `toml:"tencentKey"`
}

func getKey() *common.Credential {
	data, err := os.ReadFile(keyPath)
	if err != nil {
		fmt.Println(err)
	}
	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		fmt.Println(err)
	}
	return common.NewCredential(config.TencentKey.SecretID, config.TencentKey.SecretKey)
}

func TranslateText(text string) string {
	// 硬编码密钥到代码中有可能随代码泄露而暴露，有安全隐患，并不推荐。
	// 为了保护密钥安全，建议将密钥设置在环境变量中或者配置文件中，请参考本文凭证管理章节。
	credential := getKey()
	client, _ := tmt.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	languageRequest := tmt.NewLanguageDetectRequest()
	id := int64(0)
	languageRequest.Text = &text
	languageRequest.ProjectId = &id
	languageResponse, _ := client.LanguageDetect(languageRequest)
	lang := *languageResponse.Response.Lang

	var tar string
	request := tmt.NewTextTranslateRequest()
	request.Source = &lang
	if lang == "zh" {
		tar = "en"
	} else {
		tar = "zh"
	}
	request.SourceText = &text
	request.Target = &tar
	request.ProjectId = &id
	response, _ := client.TextTranslate(request)

	return *response.Response.TargetText
}

// 翻译img 返回原string 和翻译结果
func TranslateImg(imgPath string) (string, string) {
	credential := getKey()
	client, _ := tmt.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	languageRequest := tmt.NewLanguageDetectRequest()
	id := int64(0)
	languageRequest.ProjectId = &id
	fimg, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatal("failed to open img file")
	}
	base64Img := base64.StdEncoding.EncodeToString(fimg)

	request := tmt.NewImageTranslateRequest()
	auto, target, scene := "auto", "zh", "doc"
	uuid := uuid.New().String()
	request.Source = &auto
	request.Target = &target
	request.ProjectId = &id
	request.Data = &base64Img
	request.Scene = &scene
	request.SessionUuid = &uuid
	response, err := client.ImageTranslate(request)
	if err != nil {
		return "", ""
	}
	sor, resp := "", ""
	for _, v := range response.Response.ImageRecord.Value {
		sor += *v.SourceText + "\n"
		resp += *v.TargetText + "\n"
	}
	return sor, resp
}
