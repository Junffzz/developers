/*
@Time : 2020/5/18 10:38 
@Author : ZhaoJunfeng
@File : main
*/
package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()
var logger *logrus.Entry
func init()  {
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = os.Stdout
	log.Level = logrus.TraceLevel
	log.SetReportCaller(true)

	logger= log.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})

}

func main()  {

	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")

	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logger.WithFields(logrus.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")


}