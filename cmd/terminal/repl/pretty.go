package repl

import (
	"strconv"
	"strings"

	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/cmd/terminal/color"
)

func byteToString(b byte) string {
	return strconv.Itoa(int(b))
}

func dimUnderline(b byte) string {
	return color.Combine(byteToString(b), color.CodeDim, color.CodeUnderline)
}

func prettyMessage(message []byte, status bool) string {
	symbols := make([]string, 0)
	for i, el := range message {
		switch i {
		case 0, 1, len(message) - 1:
			symbols = append(symbols, color.Dim(byteToString(el)))
		case 2:
			symbols = append(symbols, dimUnderline(el))
		case 3, 4:
			symbols = append(symbols, color.Blue(byteToString(el)))
		case 5:
			if status && el == myrtio.SuccessCode {
				symbols = append(symbols, color.Green(byteToString(el)))
			} else if status && el == myrtio.ErrorCode {
				symbols = append(symbols, color.Red(byteToString(el)))
			} else {
				symbols = append(symbols, byteToString(el))
			}
		default:
			symbols = append(symbols, byteToString(el))
		}
	}
	return strings.Join(symbols, " ")
}
