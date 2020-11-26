package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var BgLogger *zap.SugaredLogger
const logDebugSwitch = false

func InitLogger(path string) error{
	if logDebugSwitch {
		if err := initLoggerR();err != nil {
			return err
		}
	} else {
		if err := initLoggerW(path);err != nil {
			return err
		}
	}
	defer BgLogger.Sync()

	return nil
}

func initLoggerW(filename string) (err error) {
	//写入文件的位置
	_ = os.MkdirAll("./logs",0755)
	path := fmt.Sprintf("./logs/%v",filename)
	file,err := os.OpenFile(path,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0755)
	if err != nil {
		return
	}
	writeSyncer := zapcore.AddSync(file)

	//时间格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel) //什么等级以上的日志会被写入
	logger := zap.New(core,zap.AddCaller()) //Add是打印行号和文件位置
	BgLogger = logger.Sugar()

	return
}

func initLoggerR() (err error){
	logger,err := zap.NewProduction()
	if err != nil {
		return
	}
	BgLogger = logger.Sugar()
	return
}

