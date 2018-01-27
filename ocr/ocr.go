package ocr

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"encoding/json"
	"fmt"
)

var (
	BaiduAiApiKey = "is6NaY4E8KPdhYcKdeHCDA0d"
	BaiduAiSecretKey = "fW9vVdB5ESMVF0P3kPRjBpgR8k7vW15G"
)

type Result struct {
	LogId int64 `json:"log_id"`
	WordsResultNum int64 `json:"words_result_num"`
	WordsResult []map[string]string `json:"words_result"`

}

func GetStrByBaiduAi(picPath string) (string, error) {
	ocrClient := ocr.NewOCRClient(BaiduAiApiKey, BaiduAiSecretKey)
	pic, err := vision.FromFile(picPath)
	if err != nil {
		return "", err
	}
	ret, err := ocrClient.GeneralRecognizeBasic(pic)
	if err != nil {
		return "", err
	}
	return ret.String(), nil
}

func DealOcrRet(ocrRet string) {
	var (
		result Result
		resultStr string
		tmpCount int64
	)
	json.Unmarshal([]byte(ocrRet), &result)
	for _, val := range result.WordsResult {
		tmpCount += 1
		resultStr += val["words"] + " "
		if tmpCount % 5 == 0 {
			resultStr += "\n"
		}
	}
	fmt.Println(resultStr)
}
