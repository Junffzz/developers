package main

import (
	"devtest/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error


type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

//函数包装一下，函数即作为输入，又作为输出。
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r:=recover();r!=nil{
				log.Printf("Panic:%v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}

		}()


		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request:%s\n", err.Error())

			if userErr,ok:=err.(userError);ok {
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/list/errhandle", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {

	}
}
