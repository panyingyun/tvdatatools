package main

import (
	"sort"
	"strings"
)

//典型渠道分析
func analysisVIPProducts(videos []Video, vodt VodT, outfile string) (VPro, VPro, VPro, VPro, VPro) {
	vpros := make([]VPro, 0)
	vproceb := analysisPro("CEB", "索尼", "Android2.0", videos, vodt)
	vpro30 := analysisPro("BS3.0", "", "Android3.0", videos, vodt)
	vprovidaa := analysisPro("vidaa", "海信", "SDK", videos, vodt)
	vprocs10 := analysisPro("CS1.0", "索尼", "CS1.0", videos, vodt)
	vprohx13 := analysisPro("海信1.3", "海信", "Android1.3", videos, vodt)

	vpros = append(vpros, vproceb)
	vpros = append(vpros, vpro30)
	vpros = append(vpros, vprovidaa)
	vpros = append(vpros, vprocs10)
	vpros = append(vpros, vprohx13)

	writeVProS(vpros, outfile)
	return vproceb, vpro30, vprovidaa, vprocs10, vprohx13
}

func analysisAllProducts(videos []Video, vodt VodT, outfile string) {
	var vpros VProS
	//所以产品线
	products := make([]string, 13)
	products[0] = "1.0产品线"
	products[1] = "2.0产品线"
	products[2] = "Android1.2"
	products[3] = "Android1.3"
	products[4] = "Android2.0"
	products[5] = "Android2.0/SDK"
	products[6] = "Android3.0"
	products[7] = "CS1.0"
	products[8] = "Linux2.0"
	products[9] = "SDK"
	products[10] = "其它"
	products[11] = "全部"
	products[12] = "浙江联通产品线"
	for _, product := range products {
		vpro := analysisPro(product, "", product, videos, vodt)
		vpros = append(vpros, vpro)
	}
	sort.Sort(vpros)
	writeVProS(vpros, outfile)
	return
}

//按照厂商和产品线进行分析
func analysisPro(name string, vendor string, product string, videos []Video, vodt VodT) VPro {
	var vpro VPro
	vpro.Name = name
	vpro.Vendor = vendor
	vpro.Product = product
	//考虑厂商为空的情况
	if len(vendor) != 0 && len(product) != 0 {
		for _, video := range videos {
			if strings.EqualFold(video.Vendor, vpro.Vendor) && strings.EqualFold(video.Product, vpro.Product) {
				vpro.VodCnt += video.VodCnt
				vpro.PerCnt += video.PerCnt
				vpro.Videos = append(vpro.Videos, video)
			}
		}
	} else if len(product) != 0 {
		for _, video := range videos {
			if strings.EqualFold(video.Product, vpro.Product) {
				vpro.VodCnt += video.VodCnt
				vpro.PerCnt += video.PerCnt
				vpro.Videos = append(vpro.Videos, video)
			}
		}
	} else if len(vendor) != 0 {
		for _, video := range videos {
			if strings.EqualFold(video.Vendor, vpro.Vendor) {
				vpro.VodCnt += video.VodCnt
				vpro.PerCnt += video.PerCnt
				vpro.Videos = append(vpro.Videos, video)
			}
		}
	}

	if vpro.PerCnt != 0 {
		vpro.VodPer = float32(vpro.VodCnt) / float32(vpro.PerCnt)
	}

	if vodt.PerCnt != 0 {
		vpro.PerRadio = float32(vpro.PerCnt) / float32(vodt.PerCnt)
	}

	if vodt.VodCnt != 0 {
		vpro.VodRadio = float32(vpro.VodCnt) / float32(vodt.VodCnt)
	}
	return vpro
}

//
//分析典型样本的一级栏目
//
func analysisVIPP1(vpro VPro, outfile string) VodP1S {
	var vodp1s VodP1S
	// 一级栏目数据提取
	for _, v1 := range vpro.Videos {
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
		if vpro.PerCnt != 0 {
			vodp1s[index].PerRadio = float32(v3.PerCnt) / float32(vpro.PerCnt)
		}

		if vpro.VodCnt != 0 {
			vodp1s[index].VodRadio = float32(v3.VodCnt) / float32(vpro.VodCnt)
		}

		if v3.PerCnt != 0 {
			vodp1s[index].VodPer = float32(v3.VodCnt) / float32(v3.PerCnt)
		}
	}
	writeVodP1(vodp1s, outfile)
	return vodp1s
}
