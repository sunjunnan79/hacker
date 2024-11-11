package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	code := "muxi-backend"
	password := checkpoint1_1(code)
	link1, header1 := checkpoint1_2(password)

	info, link2, header2 := checkpoint2(link1, code, password)
	inBase64, err := encrypt.AESEncryptOutInBase64([]byte(info.ErrorCode), []byte(info.SecretKey))
	if err != nil {
		return
	}

	link3, header3 := checkpoint3(link2, code, password, string(inBase64))
	link4, link5 := checkpoint4_1(link3, code, password)

	checkpoint4_2(link4, code, password)

	header4 := checkpoint4_3(link5, code, password)

	finalLink := fmt.Sprintf("http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/" + header1 + "/" + header2 + "/" + header3 + "/" + header4)

	checkpoint5_1(finalLink, code, password)

	checkpoint5_2(finalLink, code, password)
}

func checkpoint1_1(code string) string {
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", "http://gtainmuxi.muxixyz.com/api/v1/organization/code?", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头
	req.Header.Add("code", code)
	// 你可以继续添加更多的请求头

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	fmt.Println("checkpoint1-1:", response.Data.Text)
	fmt.Println()
	// 打印响应
	return resp.Header.Get("Passport")
}

func checkpoint1_2(password string) (link string, header string) {
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", "http://gtainmuxi.muxixyz.com/api/v1/organization/code?", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头
	req.Header.Add("code", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}
	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	fmt.Println("checkpoint1-2:", response.Data.Text)
	fmt.Println()
	//正则匹配获取结果
	re := regexp.MustCompile(`(http[^,]*)，`)
	link = re.FindStringSubmatch(string(body))[1]
	return link, resp.Header.Get("map-fragments")
}

func checkpoint2(link, code, password string) (info DecodedInfo, link2 string, header string) {
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头

	req.Header.Add("code", code)

	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DecodedInfo{}, "", ""
	}

	//输出结果
	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	fmt.Println("checkpoint2:", response.Data.Text)
	fmt.Println()
	// 解析 extra_info 中的内容
	info, err = decodeExtraInfo(response.Data.ExtraInfo)
	if err != nil {
		log.Fatalf("extra_info 解码失败: %v", err)
	}

	//正则匹配获取结果
	re := regexp.MustCompile(`(http[^,]*) `)
	link = re.FindStringSubmatch(string(body))[1]
	return info, link, resp.Header.Get("map-fragments")
}

func checkpoint3(link, code, password string, content string) (resplink string, header string) {
	client := &http.Client{}
	query := struct {
		Content string `json:"content"`
	}{}
	query.Content = content

	// 将请求体转换为 JSON 字节
	bodyBytes, err := json.Marshal(query)
	if err != nil {
		log.Fatal("JSON 序列化失败:", err)
	}

	// 创建请求对象，传入 JSON 请求体
	req, err := http.NewRequest("PUT", link, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头

	req.Header.Add("code", code)

	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", ""
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	fmt.Println("checkpoint3:", response.Data.Text)
	fmt.Println()
	//正则匹配获取结果
	re := regexp.MustCompile(`(http[^,]*) `)
	resplink = re.FindStringSubmatch(string(body))[1]

	return resplink, resp.Header.Get("map-fragments")
}

func checkpoint4_1(link, code, password string) (string, string) {
	client := &http.Client{}

	// 创建请求对象，传入 JSON 请求体
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加请求头

	req.Header.Add("code", code)

	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", ""
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	fmt.Println("checkpoint4_1:", response.Data.Text)
	fmt.Println()
	//正则匹配获取结果
	re := regexp.MustCompile(`http[s]?://[^\s]+`)
	links := re.FindAllString(string(body), -1)

	return links[0], links[1]
}

func checkpoint4_2(link, code, password string) {
	client := &http.Client{}

	// 创建 GET 请求
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	// 添加请求头
	req.Header.Add("code", code)
	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}

	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	// 将 Base64 编码的 ExtraInfo 转换为二进制数据
	imageData, err := base64.StdEncoding.DecodeString(response.Data.ExtraInfo)
	if err != nil {
		log.Fatalf("Base64 解码失败: %v", err)
	}

	// 将解码后的数据写入文件 "瞳孔.png"
	if err := os.WriteFile("file/瞳孔.png", imageData, 0644); err != nil {
		log.Fatalf("写入文件失败: %v", err)
		return
	}

	fmt.Println("checkpoint4_2:瞳孔图片已成功写入到file目录下的 瞳孔.png 文件")
	fmt.Println()
	return
}

func checkpoint4_3(link, code, password string) string {
	client := &http.Client{}

	// 打开图片文件
	file, err := os.Open("file/瞳孔.png")
	if err != nil {
		log.Fatal("打开图片文件失败:", err)
	}
	defer file.Close()

	// 创建一个缓冲区用于保存 multipart 表单数据
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 创建文件字段并将文件内容写入
	fileWriter, err := writer.CreateFormFile("file", "瞳孔.png")
	if err != nil {
		log.Fatal("创建文件字段失败:", err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatal("写入文件内容失败:", err)
	}

	// 关闭 multipart writer 以设置结束边界
	writer.Close()

	// 创建 POST 请求
	req, err := http.NewRequest("POST", link, &body)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	// 添加请求头，包括动态生成的 Content-Type
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("code", code)
	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}

	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	fmt.Println("checkpoint4_3:", response.Data.Text)
	fmt.Println()
	return resp.Header.Get("map-fragments")
}

func checkpoint5_1(link, code, password string) {
	client := &http.Client{}

	// 创建 POST 请求
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	req.Header.Add("code", code)
	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}
	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	fmt.Println("checkpoint5_1:", response.Data.Text)
	fmt.Println()
	return
}

func checkpoint5_2(link, code, password string) {
	client := &http.Client{}

	// 打开图片文件
	file, err := os.Open("file/permute.go")
	if err != nil {
		log.Fatal("打开图片文件失败:", err)
	}
	defer file.Close()

	// 创建一个缓冲区用于保存 multipart 表单数据
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 创建文件字段并将文件内容写入
	fileWriter, err := writer.CreateFormFile("file", "permute.go")
	if err != nil {
		log.Fatal("创建文件字段失败:", err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatal("写入文件内容失败:", err)
	}

	// 关闭 multipart writer 以设置结束边界
	writer.Close()

	// 创建 POST 请求
	req, err := http.NewRequest("POST", link, &body)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("code", code)
	req.Header.Add("passport", password)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}

	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	fmt.Println("checkpoint5_2:", response.Data.Text)
	fmt.Println()
	return
}

// 从 extra_info 中解码并提取出 secret_key 和 error_code
func decodeExtraInfo(extraInfo string) (DecodedInfo, error) {
	// 先进行 Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(extraInfo)
	if err != nil {
		return DecodedInfo{}, fmt.Errorf("base64 解码失败: %v", err)
	}

	decodedStr := string(decodedBytes)

	// 分离出 secret_key 和 error_code
	parts := strings.Split(decodedStr, ", ")
	info := DecodedInfo{}
	for _, part := range parts {
		if strings.HasPrefix(part, "secret_key:") {
			info.SecretKey = strings.TrimPrefix(part, "secret_key:")
		} else if strings.HasPrefix(part, "error_code:") {
			info.ErrorCode = strings.TrimPrefix(part, "error_code:")
		}
	}

	return info, nil
}

// 定义用于解析响应的结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Text      string `json:"text"`
	ExtraInfo string `json:"extra_info"`
}

// 定义存放解码后的密钥和错误代码的结构体
type DecodedInfo struct {
	SecretKey string `json:"secret_key"`
	ErrorCode string `json:"error_code"`
}
