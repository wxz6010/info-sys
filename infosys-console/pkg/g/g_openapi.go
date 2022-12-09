package g

import (
	"bufio"
	"com.alex.infosys/pkg/errs"
	"io/ioutil"
	"net/http"
	"os"
)

//相对于当前 main 函数所在目录的相对路径
func SwagJson2OpenApi(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errs.EXCEPTION.Wrap(err, "打开文件错误")
	}
	req, err := http.NewRequest("POST", "https://converter.swagger.io/api/convert", bufio.NewReader(file))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errs.BadGateway.Wrap(err, "Swag json 文件转换失败")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
