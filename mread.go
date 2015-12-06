package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"strconv"
	"strings"
)

//公网读取CSV版本
func readALLCSV(file string) (videos []Video, err error) {
	return readCSV(file, "")
}

//专网读取CSV版本 仅仅处理 TVID为 7001A31B1 的站点数据
func readSpecCSV(file string, tvid string) (videos []Video, err error) {
	return readCSV(file, tvid)
}

//读取CSV文件，预处理数据，返回到一个slice,用于统计
//输入:file 文件名
//输出：video 数据记录
func readCSV(file string, tvid string) (videos []Video, err error) {
	videos = make([]Video, 0)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("(ErrorCode:0001) error is " + err.Error())
		return
	}
	decode := mahonia.NewDecoder("gbk")
	if decode == nil {
		fmt.Println("(ErrorCode:0002) error is " + err.Error())
		return videos, errors.New("tmahonia.NewDecoder")
	}
	rcsv := csv.NewReader(strings.NewReader(decode.ConvertString(string(buf))))
	//允许双引号字符，否则报错 (ErrorCode:0003) error is line 29752, column 15: extraneous " in field
	rcsv.LazyQuotes = true
	//检查字段个数
	//rcsv.FieldsPerRecord = 9

	ret, err := rcsv.ReadAll()
	if err != nil {
		fmt.Println("(ErrorCode:0003) error is " + err.Error())
		return
	}

	for i, v := range ret {
		var video Video
		if len(v[0]) < 3 {
			//fmt.Printf("ID有误的资产 v = %v\n", v)
			continue
		}
		id, _ := strconv.Atoi(v[0])
		vodcnt, _ := strconv.Atoi(v[4])
		percnt, _ := strconv.Atoi(v[5])
		paths := strings.Split(v[3], "#")
		if vodcnt != 0 && percnt != 0 {
			video.ID = id
			video.Name = v[1]
			video.Vendor = v[2]
			video.VodCnt = vodcnt
			video.PerCnt = percnt
			video.TVID = v[6]
			video.Product = v[7]
			video.Chip = v[8]
			if len(paths) == 0 {
				//一般不应该出现该情况
				video.Path1 = "栏目1空"
				video.Path2 = "栏目2空"
				video.Path3 = "栏目3空"
				//fmt.Printf("(ErrorCode:0006) 栏目全是空 is index i = %v, v= %v\n", i, v)
			} else if len(paths) == 1 {
				video.Path1 = paths[0]
				video.Path2 = "栏目2空"
				video.Path3 = "栏目3空"
				if len(video.Path1) == 0 {
					video.Path1 = "栏目1空"
					//fmt.Printf("栏目1个 %v\n", v)
				}
			} else if len(paths) == 2 {
				video.Path1 = paths[0]
				video.Path2 = paths[1]
				video.Path3 = "栏目3空"
				if len(video.Path1) == 0 {
					video.Path1 = "栏目1空"
					//fmt.Printf("栏目2个 %v\n", v)
				}
				if len(video.Path2) == 0 {
					video.Path2 = "栏目2空"
					//fmt.Printf("栏目2个 %v\n", v)
				}
			} else if len(paths) == 3 {
				video.Path1 = paths[0]
				video.Path2 = paths[1]
				video.Path3 = paths[2]
				if len(video.Path1) == 0 {
					video.Path1 = "栏目1空"
					//fmt.Printf("栏目3个 %v\n", v)
				}
				if len(video.Path2) == 0 {
					video.Path2 = "栏目2空"
					//fmt.Printf("栏目3个 %v\n", v)
				}
				if len(video.Path3) == 0 {
					video.Path3 = "栏目3空"
					//fmt.Printf("栏目3个 %v\n", v)
				}
			} else {
				fmt.Printf("(ErrorCode:0005) error is index i = %v, v= %v\n", i, v)
			}

			if strings.EqualFold(video.TVID, tvid) || strings.EqualFold(tvid, "") {
				videos = append(videos, video)
			}

		} else {
			if i != 0 {
				fmt.Printf("(ErrorCode:0004) error is index i = %v, v= %v\n", i, v)
			}
		}

	}
	return
}
