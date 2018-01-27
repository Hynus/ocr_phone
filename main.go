package main

import (
	dp "ocr_phone/deal_picture"
	oc "ocr_phone/ocr"
)

func main() {
	//dp.GetScreenShotSaveInPC()
	xx, _ := oc.GetStrByBaiduAi(dp.GetPicAfterCut())
	oc.DealOcrRet(xx)
}
