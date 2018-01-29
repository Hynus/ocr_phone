package text2voice

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"os"
	"strconv"
)

var (
	BaiduAiApiKey = "grF7SveGj2LIwIYobfNBPO9n"
	BaiduAiSecretKey = "iUDZHYH8qhWWWuUc9m9A7oPcmzdzkRGH"
)

func GetVoice(voiceName, oriText string) {
	oriTextRune := []rune(oriText)
	splitOriTextNum := len(oriTextRune) / 100
	var i int
	i = 1
	for i <= splitOriTextNum {
		beginIdx := (i - 1) * 100
		endIdx := i * 100
		branchName := "_part" + strconv.Itoa(i)
		voicePartName := voiceName + branchName
		partString := string(oriTextRune[beginIdx : endIdx])
		//fmt.Println(voicePartName)
		//fmt.Println(partString)
		Text2Voice(voicePartName, partString)
		i++
	}
	restBranchName := "_part" + strconv.Itoa(i)
	restVoicePartName := voiceName + restBranchName
	restBeginIdx := (i - 1) * 100
	restPartString := string(oriTextRune[restBeginIdx:])
	Text2Voice(restVoicePartName, restPartString)
}

func Text2Voice(voicePartName, text string) error {
	//timestamp := time.Now().Unix()
	filename := "/Users/sunyh/MyWork/MyGolang/src/ocr_phone/voice_from_baiduAI/"+ voicePartName + ".mp3"
	//for windows
	//filename := "E:/ACA/Learning/MyGoLang/MyGo/src/baiduAi_voice/genVoice/" + strconv.FormatInt(timestamp,10) + ".mp3"
	client := voice.NewVoiceClient(BaiduAiApiKey, BaiduAiSecretKey)
	voiceFile, err := client.TextToSpeech(text)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(voiceFile); err != nil {
		return err
	}
	return nil
}
