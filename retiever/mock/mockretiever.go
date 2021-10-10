/*
@Time : 2020/10/22 08:51
@Author : ZhaoJunfeng
@File : mockretiever
*/
package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}