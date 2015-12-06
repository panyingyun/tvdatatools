package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"os"
	"strconv"
)

//统计数据输出到文件 总点播 到 CSV文件（文件名_总点播.csv）
//输入vodt:总点播数据
//输入outfile：输出文件名
func writeVodT(vodt VodT, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to create file %v\n", outfile)
		return err
	}

	defer file.Close()

	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}

	writer := csv.NewWriter(file)
	recordVodTHead := make([]string, 6)
	recordVodTHead[0] = encode.ConvertString("名称")
	recordVodTHead[1] = encode.ConvertString("点播数")
	recordVodTHead[2] = encode.ConvertString("点播用户")
	recordVodTHead[3] = encode.ConvertString("人均点播")
	recordVodTHead[4] = encode.ConvertString("点播量占比")
	recordVodTHead[5] = encode.ConvertString("用户数占比")

	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 6)
	recordVodT[0] = encode.ConvertString(vodt.Name)
	recordVodT[1] = encode.ConvertString(strconv.Itoa(vodt.VodCnt))
	recordVodT[2] = encode.ConvertString(strconv.Itoa(vodt.PerCnt))
	recordVodT[3] = encode.ConvertString(fmt.Sprintf("%6.3f", vodt.VodPer))
	recordVodT[4] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vodt.VodRadio*100))
	recordVodT[5] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vodt.PerRadio*100))

	err = writer.Write(recordVodT)
	if err != nil {
		fmt.Printf("Fail to write file %v\n", outfile)
		return err
	}
	writer.Flush()

	return nil
}

//统计数据输出到文件 总点播 到 CSV文件（文件名_总点播.csv）
//vodp1s:一级栏目的输出数据
//输入outfile：输出文件名
func writeVodP1(vodp1s VodP1S, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 6)
	recordVodTHead[0] = encode.ConvertString("")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	recordVodTHead[4] = encode.ConvertString("")
	recordVodTHead[5] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("频道")
	recordVodTHead[1] = encode.ConvertString("点播数")
	recordVodTHead[2] = encode.ConvertString("点播用户")
	recordVodTHead[3] = encode.ConvertString("人均点播")
	recordVodTHead[4] = encode.ConvertString("点播量占比")
	recordVodTHead[5] = encode.ConvertString("用户数占比")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 6)
	for i, v := range vodp1s {
		if i > 49 {
			break
		}
		recordVodT[0] = encode.ConvertString(v.Name)
		recordVodT[1] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		recordVodT[3] = encode.ConvertString(fmt.Sprintf("%6.3f", v.VodPer))
		recordVodT[4] = encode.ConvertString(fmt.Sprintf("%6.3f%%", v.VodRadio*100))
		recordVodT[5] = encode.ConvertString(fmt.Sprintf("%6.3f%%", v.PerRadio*100))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

//统计数据输出到文件 总点播 到 CSV文件（文件名_总点播.csv）
//vodp2s:二级栏目的输出数据
//输入outfile：输出文件名
func writeVodP2(p1 string, vodp2s VodP2S, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 6)
	recordVodTHead[0] = encode.ConvertString("")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	recordVodTHead[4] = encode.ConvertString("")
	recordVodTHead[5] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString(p1)
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	recordVodTHead[4] = encode.ConvertString("")
	recordVodTHead[5] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("频道")
	recordVodTHead[1] = encode.ConvertString("点播数")
	recordVodTHead[2] = encode.ConvertString("点播用户")
	recordVodTHead[3] = encode.ConvertString("人均点播")
	recordVodTHead[4] = encode.ConvertString("点播量占比")
	recordVodTHead[5] = encode.ConvertString("用户数占比")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 6)
	for i, v := range vodp2s {
		if i > 49 {
			break
		}
		recordVodT[0] = encode.ConvertString(v.Name)
		recordVodT[1] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		recordVodT[3] = encode.ConvertString(fmt.Sprintf("%6.3f", v.VodPer))
		recordVodT[4] = encode.ConvertString(fmt.Sprintf("%6.3f%%", v.VodRadio*100))
		recordVodT[5] = encode.ConvertString(fmt.Sprintf("%6.3f%%", v.PerRadio*100))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

//总的点播数据Top20的资产进行数据输出，我们输出前500个，和重点单片合并输出
func writeAllTop20(top20s Top20S, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to create file %v\n", outfile)
		return err
	}

	defer file.Close()

	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}

	writer := csv.NewWriter(file)
	recordVodTHead := make([]string, 4)
	recordVodTHead[0] = encode.ConvertString("单条资产Top20")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("ID")
	recordVodTHead[1] = encode.ConvertString("视频名称")
	recordVodTHead[2] = encode.ConvertString("点播数")
	recordVodTHead[3] = encode.ConvertString("点播用户")
	err = writer.Write(recordVodTHead)

	//仅输出前99
	recordVodT := make([]string, 4)
	for i, v := range top20s {
		if i > 499 {
			break
		}
		recordVodT[0] = encode.ConvertString(strconv.Itoa(v.ID))
		recordVodT[1] = encode.ConvertString(v.Name)
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[3] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()

	return nil
}

//临时做法
func createfile(outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to create file %v\n", outfile)
		return err
	}

	defer file.Close()
	return nil
}

//各个频道 三级栏目Top20
func writeP1Top20P3(p1name string, vodP3s VodP3S, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 3)
	recordVodTHead[0] = encode.ConvertString("")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString(p1name)
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("三级栏目名称")
	recordVodTHead[1] = encode.ConvertString("点播数")
	recordVodTHead[2] = encode.ConvertString("点播用户")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 3)
	for i, v := range vodP3s {
		if i > 49 {
			break
		}
		recordVodT[0] = encode.ConvertString(v.Name)
		recordVodT[1] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

//各个频道 单个资产Top20
func writeP1Top20(p1name string, top20s Top20S, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 4)
	recordVodTHead[0] = encode.ConvertString("")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString(p1name)
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("ID")
	recordVodTHead[1] = encode.ConvertString("视频名称")
	recordVodTHead[2] = encode.ConvertString("点播数")
	recordVodTHead[3] = encode.ConvertString("点播用户")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 4)
	for i, v := range top20s {
		if i > 49 {
			break
		}
		recordVodT[0] = encode.ConvertString(strconv.Itoa(v.ID))
		recordVodT[1] = encode.ConvertString(v.Name)
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[3] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

//输出典型样本的统计输出
//输入数据：vpro 样本统计输出
//
func writeVPro(vpro VPro, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 8)
	recordVodTHead[0] = encode.ConvertString("样本名称")
	recordVodTHead[1] = encode.ConvertString("厂商")
	recordVodTHead[2] = encode.ConvertString("产品")
	recordVodTHead[3] = encode.ConvertString("点播数")
	recordVodTHead[4] = encode.ConvertString("点播用户")
	recordVodTHead[5] = encode.ConvertString("人均点播次数")
	recordVodTHead[6] = encode.ConvertString("点播量占比")
	recordVodTHead[7] = encode.ConvertString("用户数占比")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 8)
	recordVodT[0] = encode.ConvertString(vpro.Name)
	recordVodT[1] = encode.ConvertString(vpro.Vendor)
	recordVodT[2] = encode.ConvertString(vpro.Product)
	recordVodT[3] = encode.ConvertString(strconv.Itoa(vpro.VodCnt))
	recordVodT[4] = encode.ConvertString(strconv.Itoa(vpro.PerCnt))
	recordVodT[5] = encode.ConvertString(fmt.Sprintf("%6.3f", vpro.VodPer))
	recordVodT[6] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vpro.VodRadio*100))
	recordVodT[7] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vpro.PerRadio*100))
	err = writer.Write(recordVodT)
	if err != nil {
		fmt.Printf("Fail to write file %v\n", outfile)
		return err
	}
	writer.Flush()
	return nil
}

//厂商和版本样本的统计输出
//输入数据：vpros 样本统计输出
func writeVProS(vpros VProS, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 8)
	recordVodTHead[0] = encode.ConvertString("样本名称")
	recordVodTHead[1] = encode.ConvertString("厂商")
	recordVodTHead[2] = encode.ConvertString("产品")
	recordVodTHead[3] = encode.ConvertString("点播数")
	recordVodTHead[4] = encode.ConvertString("点播用户")
	recordVodTHead[5] = encode.ConvertString("人均点播次数")
	recordVodTHead[6] = encode.ConvertString("点播量占比")
	recordVodTHead[7] = encode.ConvertString("用户数占比")
	err = writer.Write(recordVodTHead)

	recordVodT := make([]string, 8)
	for _, vpro := range vpros {
		recordVodT[0] = encode.ConvertString(vpro.Name)
		recordVodT[1] = encode.ConvertString(vpro.Vendor)
		recordVodT[2] = encode.ConvertString(vpro.Product)
		recordVodT[3] = encode.ConvertString(strconv.Itoa(vpro.VodCnt))
		recordVodT[4] = encode.ConvertString(strconv.Itoa(vpro.PerCnt))
		recordVodT[5] = encode.ConvertString(fmt.Sprintf("%6.3f", vpro.VodPer))
		recordVodT[6] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vpro.VodRadio*100))
		recordVodT[7] = encode.ConvertString(fmt.Sprintf("%6.3f%%", vpro.PerRadio*100))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

//输出各频道各二级栏目的TOP20资产
func writeP1P2Top20(p1name string, p2name string, top20s Top20S, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 4)
	recordVodTHead[0] = encode.ConvertString("")
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString(p1name)
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString(p2name)
	recordVodTHead[1] = encode.ConvertString("")
	recordVodTHead[2] = encode.ConvertString("")
	recordVodTHead[3] = encode.ConvertString("")
	err = writer.Write(recordVodTHead)

	recordVodTHead[0] = encode.ConvertString("ID")
	recordVodTHead[1] = encode.ConvertString("视频名称")
	recordVodTHead[2] = encode.ConvertString("点播数")
	recordVodTHead[3] = encode.ConvertString("点播用户")
	err = writer.Write(recordVodTHead)

	// 写入记录
	recordVodT := make([]string, 4)
	for i, v := range top20s {
		if i > 49 {
			break
		}
		recordVodT[0] = encode.ConvertString(strconv.Itoa(v.ID))
		recordVodT[1] = encode.ConvertString(v.Name)
		recordVodT[2] = encode.ConvertString(strconv.Itoa(v.VodCnt))
		recordVodT[3] = encode.ConvertString(strconv.Itoa(v.PerCnt))
		err = writer.Write(recordVodT)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}
