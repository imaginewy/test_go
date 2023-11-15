package logging

import (
	"fmt"
	"gin_blog/pkg/file"
	"gin_blog/pkg/setting"
	"os"
	"time"
)

// var (
// 	LogSavePath = setting.AppSetting.LogSavePath
// 	LogSaveName = setting.AppSetting.LogSaveName
// 	LogFileExt  = setting.AppSetting.LogFileExt
// 	TimeFormat  = setting.AppSetting.TimeFormat
// )

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//	func getLogFileFullPath() string {
//		prefixPath := getLogFilePath()
//		suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
//		return fmt.Sprintf("%s%s", prefixPath, suffixPath)
//	}
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
func openLogFile(filePath, fileName string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := dir + "/" + filePath
	fmt.Println(src)
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	handle, err := os.OpenFile(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return handle, nil
}
