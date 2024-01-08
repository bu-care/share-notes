package chain

import "testing"

func TestChain(t *testing.T) {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}
	// 将责任链串起来
	adHandler.next_handler = yellowHandler
	yellowHandler.next_handler = sensitiveHandler

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}
