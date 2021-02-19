//有关github的所有函数
package download

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const GET_VERSION = "https://api.github.com/repos/%s/releases/latest"

func GetLatestVersionUrl(repo string, index int) (string, error) {
	/* 获取仓库最新Release的下载url
	 *
	 * 参数：
	 * 	repo: 仓库地址，例："EricTianC/GraiaOK"
	 *	index: 下载文件的下标，下载第一个就是0
	 * 返回:
	 *	最新仓库地址  string
	 *	异常  error
	 */
	log.Printf("正在获取仓库%s的最新Release\n", repo)
	resp, err := http.Get(fmt.Sprintf(GET_VERSION, repo))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败：%s", resp.Status)
	}

	data := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&data)
	log.Printf("target version: %s    %s\n%s\n", data["tag_name"], data["published_at"], data["body"])
	download_url := data["assets"].([]interface{})[index].(map[string]interface{})["browser_download_url"].(string)
	return download_url, nil
}
