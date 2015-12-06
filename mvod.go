package main

import (
	//"fmt"
	"sort"
	"strings"
)

//
//分析频道（一级栏目）点播数据
//
func analysisVodP1(videos []Video, vodt VodT, outfile string) VodP1S {
	var vodp1s VodP1S
	// 一级栏目数据提取
	for _, v1 := range videos {
		bfind := false
		for index, v2 := range vodp1s {
			if strings.EqualFold(v1.Path1, v2.Name) {
				bfind = true
				vodp1s[index].PerCnt += v1.PerCnt
				vodp1s[index].VodCnt += v1.VodCnt
				vodp1s[index].Videos = append(vodp1s[index].Videos, v1)
				break
			}
		}
		if !bfind {
			var vodp1 VodP1
			vodp1.Name = v1.Path1
			vodp1.PerCnt = v1.PerCnt
			vodp1.VodCnt = v1.VodCnt
			vodp1.Videos = append(vodp1.Videos, v1)
			vodp1s = append(vodp1s, vodp1)
		}
	}

	//形成Top榜单
	sort.Sort(vodp1s)

	//一级栏目的结果计算处理
	for index, v3 := range vodp1s {
		if vodt.PerCnt != 0 {
			vodp1s[index].PerRadio = float32(v3.PerCnt) / float32(vodt.PerCnt)
		}

		if vodt.VodCnt != 0 {
			vodp1s[index].VodRadio = float32(v3.VodCnt) / float32(vodt.VodCnt)
		}

		if v3.PerCnt != 0 {
			vodp1s[index].VodPer = float32(v3.VodCnt) / float32(v3.PerCnt)
		}
	}
	writeVodP1(vodp1s, outfile)
	return vodp1s
}

//分析各频道二级栏目点播数据
//输入：videos 点播记录数据
//输入：VodP1S 频道数据
//输入：outfile 文件前缀
func analysisVodP2(vodp1s VodP1S, outfile string) []VodP2S {
	vodp2ss := make([]VodP2S, 0)
	for i, vp1 := range vodp1s {
		if i > 29 {
			break
		}
		//二级栏目的统计
		var vodp2s VodP2S
		for _, v1 := range vp1.Videos {
			bfind := false
			for index, v2 := range vodp2s {
				if strings.EqualFold(v1.Path2, v2.Name) {
					bfind = true
					vodp2s[index].PerCnt += v1.PerCnt
					vodp2s[index].VodCnt += v1.VodCnt
					vodp2s[index].Videos = append(vodp2s[index].Videos, v1)
					break
				}
			}
			if !bfind {
				var vodp2 VodP2
				vodp2.PName = v1.Path1
				vodp2.Name = v1.Path2
				vodp2.PerCnt = v1.PerCnt
				vodp2.VodCnt = v1.VodCnt
				vodp2.Videos = append(vodp2.Videos, v1)
				vodp2s = append(vodp2s, vodp2)
			}
		}

		//形成Top榜单
		sort.Sort(vodp2s)

		//二级栏目的结果计算处理
		for index, v3 := range vodp2s {
			if vp1.PerCnt != 0 {
				vodp2s[index].PerRadio = float32(v3.PerCnt) / float32(vp1.PerCnt)
			}

			if vp1.VodCnt != 0 {
				vodp2s[index].VodRadio = float32(v3.VodCnt) / float32(vp1.VodCnt)
			}

			if v3.PerCnt != 0 {
				vodp2s[index].VodPer = float32(v3.VodCnt) / float32(v3.PerCnt)
			}
		}
		vodp2ss = append(vodp2ss, vodp2s)
		writeVodP2(vp1.Name, vodp2s, outfile)
	}
	return vodp2ss
}

//分析总点播数据
//输入：videos 点播记录数据
//输入：outfile 文件前缀
func analysisVodT(videos []Video, outfile string) VodT {
	var vodt VodT
	vodt.Name = "总点播"

	for _, video := range videos {
		vodt.VodCnt += video.VodCnt
		vodt.PerCnt += video.PerCnt
	}
	if vodt.PerCnt != 0 {
		vodt.VodPer = float32(vodt.VodCnt) / float32(vodt.PerCnt)
	}
	vodt.VodRadio = 1.00
	vodt.PerRadio = 1.00
	//fmt.Printf("vodcnt = %v\n", vodt.VodCnt)
	//fmt.Printf("percnt = %v\n", vodt.PerCnt)
	//fmt.Printf("vodper = %v\n", vodt.VodPer)
	writeVodT(vodt, outfile)
	return vodt
}
