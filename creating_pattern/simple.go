package creatingpattern

import "fmt"

//创建接口API
type API interface {
	Say(name string) string
}

//对外暴露NewAPI方法，返回一个实现API接口的结构体实例
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//不对外暴露的结构体 hiAPI实现接口API
type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

//不对外暴露的结构体 helloAPI实现接口API
type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
