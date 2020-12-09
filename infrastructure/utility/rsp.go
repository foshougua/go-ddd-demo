package utility

import "github.com/gin-gonic/gin"

const (
	OK         = "OK"
	SERVER_ERR = "SERVER_ERR"
)

var errorText = map[string]string{
	"OK":         "成功",
	"SERVER_ERR": "服务异常",
}

func MakeLocRsp(c *gin.Context, httpStatus int, errCode string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"error":   errorText[errCode],
		"errcode": errCode,
		"data":    data,
	})
}

func MakeRsp(c *gin.Context, httpStatus int, errCode, errMsg string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"error":   errMsg,
		"errcode": errCode,
		"data":    data,
	})
}

type SvrRspInfo struct {
	HttpCode int
	ErrCode  string
	ErrMsg   string
}

func NewLocSvrRspInfo(status int, errcode string) *SvrRspInfo {
	return &SvrRspInfo{
		HttpCode: status,
		ErrCode:  errcode,
		ErrMsg:   errorText[errcode],
	}
}

func NewSvrRspInfo(status int, errcode, errmsg string) *SvrRspInfo {
	return &SvrRspInfo{
		HttpCode: status,
		ErrCode:  errcode,
		ErrMsg:   errmsg,
	}
}
func (s *SvrRspInfo) SetLoc(status int, errCode string, errMsg string) *SvrRspInfo {
	s.HttpCode = status
	s.ErrCode = errCode
	s.ErrMsg = errMsg
	return s
}

func (s *SvrRspInfo) MakeRsp(c *gin.Context, data interface{}) {
	c.JSON(s.HttpCode, gin.H{
		"error":   s.ErrMsg,
		"errcode": s.ErrCode,
		"data":    data,
	})
}
