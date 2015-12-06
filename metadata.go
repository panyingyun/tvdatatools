package main

/**
  单条视频记录的定义
**/
type Video struct {
	ID      int    //视频ID
	Name    string //视频名称
	Vendor  string //厂商
	Path1   string //栏目1
	Path2   string //栏目2
	Path3   string //栏目3
	VodCnt  int    //点播数
	PerCnt  int    //点播用户数
	TVID    string // TVID
	Product string // 产品线 Android1.3 /2.0/3.0
	Chip    string // 芯片
}

//总点播
type VodT struct {
	Name     string  //总点播
	VodCnt   int     //总点播数
	PerCnt   int     //总点播用户数
	VodPer   float32 //人均点播次数
	VodRadio float32 //点播量占比
	PerRadio float32 //用户数占比
}

//一级栏目点播
type VodP1 struct {
	Name     string  //一级栏目名称
	VodCnt   int     //点播数
	PerCnt   int     //点播用户数
	VodPer   float32 //人均点播次数
	VodRadio float32 //点播量占比
	PerRadio float32 //用户数占比
	Videos   []Video
}

type VodP1S []VodP1

func (slice VodP1S) Len() int {
	return len(slice)
}

func (slice VodP1S) Less(i, j int) bool {
	return slice[i].VodCnt > slice[j].VodCnt
}

func (slice VodP1S) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//二级栏目点播
type VodP2 struct {
	PName    string  //频道名称
	Name     string  //二级栏目名称
	VodCnt   int     //点播数
	PerCnt   int     //点播用户数
	VodPer   float32 //人均点播次数
	VodRadio float32 //点播量占比
	PerRadio float32 //用户数占比
	Videos   []Video
}

type VodP2S []VodP2

func (slice VodP2S) Len() int {
	return len(slice)
}

func (slice VodP2S) Less(i, j int) bool {
	return slice[i].VodCnt > slice[j].VodCnt
}

func (slice VodP2S) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//三级栏目点播
type VodP3 struct {
	Name   string //三级栏目名称
	VodCnt int    //点播数
	PerCnt int    //点播用户数
}

type VodP3S []VodP3

func (slice VodP3S) Len() int {
	return len(slice)
}

func (slice VodP3S) Less(i, j int) bool {
	return slice[i].VodCnt > slice[j].VodCnt
}

func (slice VodP3S) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//总点播的单条资产Top20和各频道的Top20
type Top20 struct {
	ID     int    //视频ID
	Name   string //视频名称
	VodCnt int    //点播数
	PerCnt int    //点播用户数
}
type Top20S []Top20

func (slice Top20S) Len() int {
	return len(slice)
}

func (slice Top20S) Less(i, j int) bool {
	return slice[i].VodCnt > slice[j].VodCnt
}

func (slice Top20S) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//==============典型样本===============
//Vender和Product 维度（VPro）
type VPro struct {
	Name     string  //样本名称
	Vendor   string  //厂商
	Product  string  //产品
	VodCnt   int     //点播数
	PerCnt   int     //点播用户数
	VodPer   float32 //人均点播次数
	VodRadio float32 //点播量占比
	PerRadio float32 //用户数占比
	Videos   []Video
}

type VProS []VPro

func (slice VProS) Len() int {
	return len(slice)
}

func (slice VProS) Less(i, j int) bool {
	return slice[i].VodCnt > slice[j].VodCnt
}

func (slice VProS) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
