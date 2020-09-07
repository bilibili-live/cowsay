package cowsay

import (
	"fmt"
	"sort"
	"strings"
)

// 获取数组最长的
func findMaxLength(lines []string) int {
	if len(lines) == 1 {
		return len(lines[0])
	}
	var arr []int
	for _, item := range lines {
		var l = len(item)
		arr = append(arr, l)
	}
	sort.Ints(arr)
	var result = arr[len(arr)-1]
	return result
}

func createLengthPaddingString(padding string, length int) string {
	var result = ""
	for i := 0; i < length; i++ {
		result += padding
	}
	return result
}

func createContentString(padding string, rawString string, length int) string {
	var ctx = fmt.Sprintf("%v %v ", padding, rawString)
	var o = len(ctx)
	var r = length - o
	for i := 0; i < r; i++ {
		if i == r-1 {
			ctx += padding
		} else {
			ctx += " "
		}
	}
	return ctx
}

// 填充
func padCenter(padding int, paddingString string, targetString string) string {
	return ""
}

func buildBalloon(say string) string {
	var lines = strings.Split(say, "\n")
	var maxLen = findMaxLength(lines)
	maxLen = maxLen + 2*2
	var commentLine = createLengthPaddingString("#", maxLen)
	var result []string
	for _, item := range lines {
		var run = createContentString("#", item, maxLen)
		// fmt.Println(run)
		result = append(result, run)
	}
	var send []string = []string{
		commentLine,
	}
	send = append(send, result...)
	send = append(send, commentLine)
	var next = strings.Join(send, "\n")
	return next
}

// RenderCowsay render cowsay
func RenderCowsay(say string) string {
	var context = buildBalloon(say)
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`
	var result = context + "\n" + cow
	return result
}
