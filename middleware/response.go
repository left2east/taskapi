package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写Write方法，将响应体内容写入缓冲区
func (w *bodyLogWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

// 返回格式msg,code,data
func RewriteResponse() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 创建一个 buffer 来捕获响应体
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 继续处理请求
		c.Next()

		// 捕获响应体的内容
		originalBody := blw.body.Bytes()
		// 解析原始响应体
		var originalResponse map[string]interface{}
		if err := json.Unmarshal(originalBody, &originalResponse); err != nil {
			log.Printf("Failed to parse original response: %v", err)
			// 如果解析失败，直接返回原始响应
			blw.ResponseWriter.Write(originalBody)
			return
		}

		// 构建新的响应体
		newResponse := gin.H{
			"code": http.StatusOK, // 默认状态码为 200
			"msg":  "success",     // 默认消息为 "success"
			"data": originalResponse,
		}

		// 如果原始响应体中有 "code" 字段，则使用它
		if code, exists := originalResponse["code"]; exists {
			newResponse["code"] = code
		}

		// 如果原始响应体中有 "msg" 字段，则使用它
		if msg, exists := originalResponse["msg"]; exists {
			newResponse["msg"] = msg
		}

		// 将新的响应体写回
		res, _ := json.Marshal(newResponse)
		blw.ResponseWriter.Write(res)
	}
}
