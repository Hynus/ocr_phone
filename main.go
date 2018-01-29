package main

import (
	dp "ocr_phone/deal_picture"
	oc "ocr_phone/ocr"
	t2v "ocr_phone/text2voice"
)

func main() {
	//dp.GetScreenShotSaveInPC())
	picPath, json, _ := oc.GetStrByBaiduAi(dp.GetPicAfterCut())
	voiceName, retStr := oc.DealOcrRet(picPath, json)
	t2v.GetVoice(voiceName, retStr)
}
