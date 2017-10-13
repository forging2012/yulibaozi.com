package toolkit


import (
	"bytes"
	"strconv"
	"math/rand"
)

var (
	defaultNumbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func RandomIntCaptcha(captchaLen int) string {
	captchaBuf := bytes.NewBufferString("");
	for i := 0; i < captchaLen; i ++ {
		captchaBuf.WriteString(strconv.Itoa(defaultNumbers[rand.Intn(len(defaultNumbers))]))
	}
	return captchaBuf.String()
}