/*
@Time : 2020/5/18 10:38 
@Author : ZhaoJunfeng
@File : main
*/
package main

//import (
//	"fmt"
//	"time"
//	"github.com/rs/zerolog"
//	"github.com/rs/zerolog/log"
//)
//
//func main()  {
//	t := time.Now()
//	url:="this is url."
//
//	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
//
//	zerolog.TimestampFieldName = "x_timestamp"
//	zerolog.LevelFieldName = "x_level"
//	zerolog.MessageFieldName = "x_msg"
//	zerolog.CallerFieldName = "x_file"
//	elapsed := time.Since(t)
//
//	log.Logger = log.With().Stack().Caller().Logger()
//
//	log.Debug().
//		Str("url", url).
//		Float64("height", 100.0).
//		Dur("elapsed",elapsed).
//		Msg("create object")
//	var err error
//	log.Error().Stack().Err(err).Msg("This is error")
//	go func() {
//		fmt.Println("zjftest")
//		//sugar.Infof("Failed to fetch URL: %s", url)
//	}()
//}
