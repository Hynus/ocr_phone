package deal_picture

import (
	"time"
	"strconv"
	"os/exec"
	"fmt"
	"os"
	"bufio"
	"io"
	"image"
	"image/png"
)

var (
	savePathPhone = "/sdcard/from_ADB/screenshot_"
	savePathPC = "E:/ACA/Learning/MyGoLang/MyGo/src/ocr_phone/images/screenshot_"
	cutPicPath = "E:/ACA/Learning/MyGoLang/MyGo/src/ocr_phone/images/after_cut/cut_"
	topLeftCornerX = 45
	topLeftCornerY = 125
	LowerRightCornerX = 1040
	LowerRightCornerY = 1920
)

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
	return true
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getScreenShotSaveInPhone() (string, bool) {
	timestamp := time.Now().Unix()
	filePathPhone := savePathPhone + strconv.FormatInt(timestamp, 10) + ".png"
	params := []string{"/c", "adb", "shell", "/system/bin/screencap", "-p", filePathPhone}
	saveBool :=  execCommand("cmd", params)
	if saveBool {
		return filePathPhone, true
	}
	return "", false
}

func GetScreenShotSaveInPC() (string, bool) {
	timestamp := time.Now().Unix()
	filePathPC := savePathPC + strconv.FormatInt(timestamp, 10) + ".png"
	filePathPhone, _ := getScreenShotSaveInPhone()
	params := []string{"/c", "adb", "pull", filePathPhone, filePathPC}
	saveBool := execCommand("cmd", params)
	if saveBool {
		if fileExist(filePathPC) {
			return filePathPC, true
		}
	}
	return "", false
}

func cutPic(path string) (string, error) {
	timestamp := time.Now().Unix()
	reader, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return "", err
	}
	filename := cutPicPath + strconv.FormatInt(timestamp,10) + ".png"
	img, _, _ := image.Decode(reader)
	rgbImg := img.(*image.NRGBA)
	subImg := rgbImg.SubImage(image.Rect(topLeftCornerX, topLeftCornerY, LowerRightCornerX, LowerRightCornerY)).(*image.NRGBA)
	f, err := os.Create(filename)
	defer f.Close()
	err = png.Encode(f, subImg)
	return filename, nil
}

func GetPicAfterCut() string {
	p, _ := GetScreenShotSaveInPC()
	cutPicPath, _ := cutPic(p)
	return cutPicPath
}
