package test

import (
	"ginStudy/service"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"

	//"net/http"
	"net/http/httptest"
	"testing"
)

/**
测试接口
*/
func TestPingRoute(t *testing.T) {
	router := service.SetEngine()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

/**
自定义hook
*/

func TestHook(t *testing.T) {
	CreateHookFunc()
	AddMethod(func() {
		log.Println("插入到第二个方法上")
	}, 2)

	log.Println(funcChain)

	for _, f := range funcChain {
		f()
	}
}

var funcChain []func()

var i = 0

func CreateHookFunc() {
	funcChain = append(funcChain, method)
	funcChain = append(funcChain, method)
	funcChain = append(funcChain, method)
	funcChain = append(funcChain, method)
	funcChain = append(funcChain, method)
}

func AddMethod(method func(), index int) {
	log.Println(funcChain)

	funcChain1 := funcChain[:index]
	//funcChain2 := funcChain[index:]// 插入的是地址
	//funcChain2 := make([]func(), 5)
	//copy(funcChain2, funcChain[index:]) // 必须对应相同的len
	var funcChain2 []func()
	funcChain2 = append(funcChain2, funcChain[index:]...) // 这样插入值

	log.Println(funcChain1)
	log.Println(funcChain2)

	funcChain1 = append(funcChain1, method)
	funcChain = append(funcChain1, funcChain2...)
}

func method() {
	log.Println(i)
	i++
}
