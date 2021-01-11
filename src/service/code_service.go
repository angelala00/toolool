package service

import (
	"encoding/hex"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

func Encode(c *gin.Context) {
	str := c.Query("str")

	gbkstr := mahonia.NewEncoder("gbk").ConvertString(str)
	gb18030str, _ := simplifiedchinese.GB18030.NewEncoder().String(str)
	gb2312str, _ := simplifiedchinese.HZGB2312.NewEncoder().String(str)
	utf8str, _ := unicode.UTF8.NewDecoder().String(str)

	gbkHexstr := ConvertToHexString(gbkstr)
	gb18030Hexstr := ConvertToHexString(gb18030str)
	gb2312Hexstr := ConvertToHexString(gb2312str)
	utf8Hexstr := ConvertToHexString(utf8str)

	objs := make(map[string]string)
	objs["gbkHexstr"] = gbkHexstr
	objs["gb18030Hexstr"] = gb18030Hexstr
	objs["gb2312Hexstr"] = gb2312Hexstr
	objs["utf8Hexstr"] = utf8Hexstr

	renderData(c, objs, nil)
}

func ConvertToHexString(sourcestr string) string {
	var data = []byte(sourcestr)
	encodedStr := hex.EncodeToString(data)
	return encodedStr
}
