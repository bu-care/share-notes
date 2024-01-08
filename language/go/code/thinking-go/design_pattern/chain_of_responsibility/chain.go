package chain

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(content string)
	next(content string)
}

// 广告过滤
type AdHandler struct {
	next_handler Handler
}

func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
	ad.next(newContent)
}

func (ad *AdHandler) next(content string) {
	if ad.next_handler != nil {
		ad.next_handler.Handle(content)
	}
}

// 涉黄过滤
type YellowHandler struct {
	next_handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	fmt.Println(newContent)
	yellow.next(newContent)
}

func (yellow *YellowHandler) next(content string) {
	if yellow.next_handler != nil {
		yellow.next_handler.Handle(content)
	}
}

// 敏感词过滤
type SensitiveHandler struct {
	next_handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤。。。")
	newContent := strings.Replace(content, "敏感词", "***", 1)
	fmt.Println(newContent)
	sensitive.next(newContent)
}

func (sensitive *SensitiveHandler) next(content string) {
	if sensitive.next_handler != nil {
		sensitive.next_handler.Handle(content)
	}
}
