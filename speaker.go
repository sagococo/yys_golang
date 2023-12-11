package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func downloadIfNotExist(filename string, text string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		downloadTextTrace(filename, text)
	}
}

func downloadTextTrace(filename, text string) {
	url := fmt.Sprintf("https://fanyi.baidu.com/gettts?lan=zh&text=%v&spd=5&source=web", text)
	// 发送请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	deleteFileIfEmpty(filename, text)
}

func deleteFileIfEmpty(filename string, text string) {
	// 获取文件信息
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		err = os.Remove(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		downloadIfNotExist(filename, text)
	}
}

func play(text string) error {
	// return nil
	filename := fmt.Sprintf("audios/%v.mp3", text)
	downloadIfNotExist(filename, text)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()
	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
