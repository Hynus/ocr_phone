package ocr

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	BaiduAiApiKey = "grF7SveGj2LIwIYobfNBPO9n"
	BaiduAiSecretKey = "iUDZHYH8qhWWWuUc9m9A7oPcmzdzkRGH"
)

type Result struct {
	LogId int64 `json:"log_id"`
	WordsResultNum int64 `json:"words_result_num"`
	WordsResult []map[string]string `json:"words_result"`

}

func GetStrByBaiduAi(picPath string) (string, string, error) {
	ocrClient := ocr.NewOCRClient(BaiduAiApiKey, BaiduAiSecretKey)
	pic, err := vision.FromFile(picPath)
	if err != nil {
		return "","", err
	}
	ret, err := ocrClient.GeneralRecognizeBasic(pic)
	if err != nil {
		return "","", err
	}
	return picPath, ret.String(), nil
}

func DealOcrRet(picPath, ocrRet string) (string, string) {
	var (
		result Result
		resultStr string
	)
	tPicSlice := strings.Split(picPath, "/")
	picName := tPicSlice[len(tPicSlice)-1]
	voiceName := strings.Split(picName,".")[0]
	json.Unmarshal([]byte(ocrRet), &result)
	for _, val := range result.WordsResult {
		resultStr += val["words"]
	}
	fmt.Printf(resultStr)
	return voiceName, resultStr
}
