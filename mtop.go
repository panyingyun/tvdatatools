package main

import (
	"fmt"
	"sort"
	"strings"
)

//总榜单的单条资产TOP20
func analysisAllTop20(videos []Video, outfile string) {
	var top20s Top20S
	total := len(videos)
	for i, v1 := range videos {
		if total != 0 {
			fmt.Printf("\r总榜单的单条资产TOP20计算中... %6.3f%%", float64((i+1)*100)/float64(total))
		}

		bfind := false
		for index, v2 := range top20s {
			if v1.ID == v2.ID {
				bfind = true
				top20s[index].PerCnt += v1.PerCnt
				top20s[index].VodCnt += v1.VodCnt
				break
			}
		}
		if !bfind {
			var top20 Top20
			top20.ID = v1.ID
			top20.Name = v1.Name
			top20.PerCnt = v1.PerCnt
			top20.VodCnt = v1.VodCnt
			top20s = append(top20s, top20)
		}
	}
	//形成Top榜单
	sort.Sort(top20s)
	writeAllTop20(top20s, outfile)
}

//各频道的Top20三级栏目
func analysisP1Top20P3(vodp1s VodP1S, outfile string) {
	createfile(outfile)
	//各频道
	for i, vp1 := range vodp1s {
		if i > 13 {
			break
		}
		var vodP3s VodP3S
		total := len(vp1.Videos)
		for m, v1 := range vp1.Videos {
			if total != 0 {
				fmt.Printf("\r频道(%v)Top20三级栏目计算中 %6.3f%%...............", vp1.Name, float64((m+1)*100)/float64(total))
			}
			bfind := false
			for index, v2 := range vodP3s {
				if strings.EqualFold(v1.Path3, v2.Name) {
					bfind = true
					vodP3s[index].PerCnt += v1.PerCnt
					vodP3s[index].VodCnt += v1.VodCnt
					break
				}
			}
			if !bfind {
				var vodp3 VodP3
				vodp3.Name = v1.Path3
				vodp3.PerCnt = v1.PerCnt
				vodp3.VodCnt = v1.VodCnt
				vodP3s = append(vodP3s, vodp3)
			}
		}
		//形成Top榜单
		sort.Sort(vodP3s)
		writeP1Top20P3(vp1.Name, vodP3s, outfile)
	}
}

//各频道的Top20单条资产
func analysisP1Top20(vodp1s VodP1S, outfile string) {
	createfile(outfile)
	//各频道
	for i, vp1 := range vodp1s {
		if i > 13 {
			break
		}
		var top20s Top20S
		total := len(vp1.Videos)
		for m, v1 := range vp1.Videos {
			if total != 0 {
				fmt.Printf("\r频道(%v)Top20单个资产计算中 %6.3f%%...............", vp1.Name, float64((m+1)*100)/float64(total))
			}
			bfind := false
			for index, v2 := range top20s {
				if v1.ID == v2.ID {
					bfind = true
					top20s[index].PerCnt += v1.PerCnt
					top20s[index].VodCnt += v1.VodCnt
					break
				}
			}
			if !bfind {
				var top20 Top20
				top20.ID = v1.ID
				top20.Name = v1.Name
				top20.PerCnt = v1.PerCnt
				top20.VodCnt = v1.VodCnt
				top20s = append(top20s, top20)
			}
		}
		//形成Top榜单
		sort.Sort(top20s)
		writeP1Top20(vp1.Name, top20s, outfile)
	}
}

//
//各频道二级栏目下Top20单片资产排行
func analysisP1P2TOP20(vodp2ss []VodP2S, outfile string) {
	createfile(outfile)
	for _, vodp2s := range vodp2ss {
		for m, vp1 := range vodp2s {
			total := len(vodp2s)
			if total != 0 {
				fmt.Printf("\r频道(%v) 二级栏目（%v） Top20单个资产计算中 %6.3f%%...............", vp1.PName, vp1.Name, float64((m+1)*100)/float64(total))
			}
			var top20s Top20S
			for _, v1 := range vp1.Videos {
				bfind := false
				for index, v2 := range top20s {
					if v1.ID == v2.ID {
						bfind = true
						top20s[index].PerCnt += v1.PerCnt
						top20s[index].VodCnt += v1.VodCnt
						break
					}
				}
				if !bfind {
					var top20 Top20
					top20.ID = v1.ID
					top20.Name = v1.Name
					top20.PerCnt = v1.PerCnt
					top20.VodCnt = v1.VodCnt
					top20s = append(top20s, top20)
				}
			}
			//形成Top榜单
			sort.Sort(top20s)
			writeP1P2Top20(vp1.PName, vp1.Name, top20s, outfile)
		}
	}
}
