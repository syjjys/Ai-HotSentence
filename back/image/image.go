package image

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

/**
 * @Author: thzhang4
 * @Description: 文生图TTI demo
 * @File: ttiDemo
 * @Date: 2023/9/26 17:33
 */

const (
	addr      = "https://spark-api.cn-huabei-1.xf-yun.com/v2.1/tti" // tti现网地址
	appId     = ""                                          // 你的appId  (需要产品授权 才可以调用)
	apiKey    = ""                  // 你的appid对应的key
	apiSecret = ""                  // 你的appid对应的secret
)

func GenerateImage(prompt string) string {
	authAddr := AssembleAuthUrl("POST", addr, apiKey, apiSecret)

	reqMsg := []message{
		{Content: prompt, Role: "user"}, // 只能有1个  不能传对话历史
	}

	req := map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": appId,
			"uid":    "tti_demo",
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain": "general",
				"width":  1024,
				"height": 1024,
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": reqMsg,
			},
		},
	}
	reqData, _ := json.Marshal(req)

	// 发起http请求
	respData, err := HttpTool("POST", authAddr, reqData, 20000)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 解析结果
	var result Response
	err = json.Unmarshal(respData, &result)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("code", result.Header.Code)
	fmt.Println("codeMsg", result.Header.Message)
	if len(result.Payload.Choices.Text) <= 0 {
		return ""
	}
	// 取出图片 解析并保存
	base64Image := result.Payload.Choices.Text[0].Content
	image, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 保存图片
	fname := uuid.New().String() + ".png"
	ioutil.WriteFile(fname, image, 0777)
	return fname
}

// 鉴权
// @hosturl :  like  wss://ws-api.xfyun.cn/v2/iat
// @apikey : apiKey
// @apiSecret : apiSecret
// 注意：如果是ws请求，method应该是"GET";  如果是POST请求,method应该是"POST"
func AssembleAuthUrl(method, addr, apiKey, apiSecret string) string {
	if apiKey == "" && apiSecret == "" { // 不鉴权
		return addr
	}
	ul, err := url.Parse(addr) // 地址不对  也不鉴权
	if err != nil {
		return addr
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, method + " " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	//签名结果
	sha := hmacWithShaToBase64("hmac-sha256", sgin, apiSecret)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	authAddr := addr + "?" + v.Encode()
	return authAddr
}

func hmacWithShaToBase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func HttpTool(method, authAddr string, data []byte, timeout int) ([]byte, error) {
	// 发起请求
	req, err := http.NewRequest(method, authAddr, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respData, err := io.ReadAll(resp.Body)
	return respData, err
}

// 响应给用户的数据格式
type Response struct {
	Header struct {
		Code    int    `json:"code"`    // 0
		Message string `json:"message"` // Success
		Sid     string `json:"sid"`
		Status  int    `json:"status"` // 1表示会话的中间结果 2表示整个会话最后结果
	} `json:"header"`
	Payload struct {
		Choices *ChoiceText `json:"choices,omitempty"`
	}
}

type ChoiceText struct {
	Status int       `json:"status"` // 0表示第一个结果 1表示中间结果 2表示最后一个结果
	Seq    int       `json:"seq"`    // 结果编号 0，1，2，3...
	Text   []message `json:"text"`   // 结果文本
}

type message struct {
	Content string `json:"content"` // 用户的对话内容 或者 是base64的图片
	Role    string `json:"role"`    // system表示存放基本的prompt,人设信息位于顶层 user表示用户(问题或者图片)  assistant表示大模型
	// 多模数据的描述
	ContentType string `json:"content_type,omitempty"` // image表示图片  text是文本
	Index       int    `json:"index"`                  // 大模型的结果序号，在多候选中使用。  这个传给大模型引擎是多余的，但是不影响。  但是这个字段必须要传给用户，因此不能加omitempty属性
}
