package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tvdatatools"
	app.Usage = "Auto analysis TV's data!"
	app.Author = "PanYingYun"
	app.Email = "panyingyun@gmail.com/panyy@wasu.com"
	app.Version = "1.7.0"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		//全网数据
		cli.StringFlag{Name: "input,i", Value: "", Usage: "tvdatatools -i XXX.csv"},
		//专网数据，按照TVID来进行分解
		cli.StringFlag{Name: "tvid,t", Value: "", Usage: "tvdatatools -i XXX.csv -t 7001A31B1"},
	}

	app.Action = func(c *cli.Context) {
		infile := c.String("input")
		tvid := c.String("tvid")
		fmt.Println("input = " + infile)
		fmt.Println("tvid = " + tvid)

		if len(infile) != 0 {
			if len(tvid) == 0 {
				//全网数据统计
				fmt.Println(<-ALLOTTVodData(infile))
			} else {
				//专网数据统计
				fmt.Println(<-SpecialOTTVodData(infile, tvid))
			}

		} else {
			fmt.Println("Input file is not exist!")
			fmt.Println("Usage: tvdatatools -i XXX.csv")
		}

	}

	app.Run(os.Args)
}

//全网OTT点播数据分析和处理
//输入:infile 点播数据文件名
func ALLOTTVodData(infile string) <-chan string {
	c := make(chan string)
	go func() {
		//获取文件名前缀
		prefixname := strings.Replace(infile, ".csv", "", -1) + "_"
		//从文件读取并预处理数据
		start := time.Now()

		videos, _ := readALLCSV(infile)
		//videos, _ := readSpecCSV(infile, "7001A31B1")

		end := time.Now()
		fmt.Printf("数据总条数 = %v\n", len(videos))
		fmt.Printf("数据读取耗时 %v(秒)\n", end.Sub(start).Seconds())

		//==================点播数据统计=====================
		outfile := prefixname + "点播.csv"
		//总点播统计
		start = time.Now()
		vodt := analysisVodT(videos, outfile)
		//各频道
		vodp1s := analysisVodP1(videos, vodt, outfile)
		//各频道二级栏目统计
		vodp2ss := analysisVodP2(vodp1s, outfile)
		end = time.Now()
		fmt.Printf("点播统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//=====================典型样本统计==================
		//典型样本 点播数据
		start = time.Now()
		outfile = prefixname + "产品线典型样本.csv"
		vproceb, vpro30, vprovidaa, vprocs10, vprohx13 := analysisVIPProducts(videos, vodt, outfile)
		end = time.Now()
		fmt.Printf("典型样本CEB/TV3.0/VIDAA/CS1.0耗时 %v(秒)\n", end.Sub(start).Seconds())
		//所有产品线点播数据
		start = time.Now()
		outfile = prefixname + "产品线全部样本.csv"
		analysisAllProducts(videos, vodt, outfile)
		end = time.Now()
		fmt.Printf("所有产品线样本耗时 %v(秒)\n", end.Sub(start).Seconds())

		//CEB各频道排行统计数据
		start = time.Now()
		outfile = prefixname + "CEB各频道排行统计数据.csv"
		createfile(outfile)
		analysisVIPP1(vproceb, outfile)
		end = time.Now()
		fmt.Printf("CEB各频道排行统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//TV3.0各频道排行统计数据
		start = time.Now()
		outfile = prefixname + "TV3.0各频道排行统计数据.csv"
		createfile(outfile)
		analysisVIPP1(vpro30, outfile)
		end = time.Now()
		fmt.Printf("TV3.0各频道排行统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//Vidaa各频道排行统计数据
		start = time.Now()
		outfile = prefixname + "Vidaa各频道排行统计数据.csv"
		createfile(outfile)
		analysisVIPP1(vprovidaa, outfile)
		end = time.Now()
		fmt.Printf("Vidaa各频道排行统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//CS1.0各频道排行统计数据
		start = time.Now()
		outfile = prefixname + "CS1.0各频道排行统计数据.csv"
		createfile(outfile)
		analysisVIPP1(vprocs10, outfile)
		end = time.Now()
		fmt.Printf("CS1.0各频道排行统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//海信1.3各频道排行统计数据
		start = time.Now()
		outfile = prefixname + "海信1.3各频道排行统计数据.csv"
		createfile(outfile)
		analysisVIPP1(vprohx13, outfile)
		end = time.Now()
		fmt.Printf("海信1.3频道排行统计耗时 %v(秒)\n", end.Sub(start).Seconds())

		//====================排行数据统计===================
		start = time.Now()
		outfile = prefixname + "各频道三级栏目TOP20.csv"
		analysisP1Top20P3(vodp1s, outfile)
		end = time.Now()
		fmt.Printf("\n各频道三级栏目TOP20耗时 %v(秒)\n", end.Sub(start).Seconds())

		//各频道各二级栏目的TOP20单片
		start = time.Now()
		outfile = prefixname + "各频道各二级栏目的TOP20.csv"
		analysisP1P2TOP20(vodp2ss, outfile)
		end = time.Now()
		fmt.Printf("\n各频道各二级栏目的TOP20耗时 %v(秒)\n", end.Sub(start).Seconds())

		//各频道单个资产TOP20
		start = time.Now()
		outfile = prefixname + "各频道单个资产TOP20.csv"
		analysisP1Top20(vodp1s, outfile)
		end = time.Now()
		fmt.Printf("\n各频道单个资产TOP20耗时 %v(秒)\n", end.Sub(start).Seconds())

		//总表的资产TOP20（重点单片）
		start = time.Now()
		outfile = prefixname + "总单个资产TOP20(重点单片).csv"
		analysisAllTop20(videos, outfile)
		end = time.Now()
		fmt.Printf("\n总单个资产TOP20耗时 %v(秒)\n", end.Sub(start).Seconds())

		c <- "OK, All TV Data Analysis Finish!!!"
	}()
	return c
}

//专网OTT点播数据分析和处理
//输入:infile 点播数据文件名
func SpecialOTTVodData(infile string, tvid string) <-chan string {
	c := make(chan string)
	go func() {
		//获取文件名前缀
		prefixname := strings.Replace(infile, ".csv", "", -1) + "_" + tvid + "_"
		//从文件读取并预处理数据
		start := time.Now()
		videos, _ := readSpecCSV(infile, tvid)
		end := time.Now()
		fmt.Printf("TVID %v 数据总条数 = %v\n", tvid, len(videos))
		fmt.Printf("TVID %v 数据读取耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		//==================点播数据统计=====================
		outfile := prefixname + "点播.csv"
		//总点播统计
		start = time.Now()
		vodt := analysisVodT(videos, outfile)
		//各频道
		vodp1s := analysisVodP1(videos, vodt, outfile)
		//各频道二级栏目统计
		vodp2ss := analysisVodP2(vodp1s, outfile)
		end = time.Now()
		fmt.Printf("TVID %v 点播统计耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		//====================排行数据统计===================
		start = time.Now()
		outfile = prefixname + "各频道三级栏目TOP20.csv"
		analysisP1Top20P3(vodp1s, outfile)
		end = time.Now()
		fmt.Printf("\n TVID %v 各频道三级栏目TOP20耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		//各频道各二级栏目的TOP20单片
		start = time.Now()
		outfile = prefixname + "各频道各二级栏目的TOP20.csv"
		analysisP1P2TOP20(vodp2ss, outfile)
		end = time.Now()
		fmt.Printf("\n TVID %v 各频道各二级栏目的TOP20耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		//各频道单个资产TOP20
		start = time.Now()
		outfile = prefixname + "各频道单个资产TOP20.csv"
		analysisP1Top20(vodp1s, outfile)
		end = time.Now()

		fmt.Printf("\n TVID %v 各频道单个资产TOP20耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		//总表的资产TOP20（重点单片）
		start = time.Now()
		outfile = prefixname + "总单个资产TOP20(重点单片).csv"
		analysisAllTop20(videos, outfile)
		end = time.Now()
		fmt.Printf("\n TVID %v 总单个资产TOP20耗时 %v(秒)\n", tvid, end.Sub(start).Seconds())

		c <- "OK, All TV Data Analysis Finish!!!"
	}()
	return c
}
