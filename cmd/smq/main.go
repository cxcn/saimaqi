package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	smq "github.com/cxcn/gosmq"
	"github.com/jessevdk/go-flags"
)

func main() {

	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	if len(os.Args) == 1 {
		fmt.Println("请在命令行中运行此程序\n按Enter键退出...")
		fmt.Scanln()
		return
	}

	type option struct {
		Fpd  []string `short:"i" long:"input" description:"[]string\t码表路径，可设置多个"`
		Ding int      `short:"d" long:"ding" description:"int\t普通码表起顶码长，码长大于等于此数，首选不会追加空格"`
		IsS  bool     `short:"s" long:"single" description:"bool\t是否只跑单字"`

		Fpt string `short:"t" long:"text" description:"string\t文本"`
		Csk string `short:"c" default:";'" description:"string\t自定义选重键(2重开始)"`
		AS  bool   `short:"k" description:"bool\t空格是否互击"`

		IsO bool `short:"o" long:"output" description:"bool\t是否输出结果"`
		Ver bool `short:"v" long:"version" description:"bool\t查看版本信息"`
	}

	var opt option
	flags.Parse(&opt)
	if opt.Ver {
		fmt.Printf("smq-cli version 0.10 %s/%s\n\n", runtime.GOOS, runtime.GOARCH)
		fmt.Println("repo address: https://github.com/cxcn/gosmq/")
		return
	}

	if len(opt.Fpd) == 0 {
		return
	}

	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Println("main cost time = ", cost)
	}()

	h := NewHTML(smq.GetFileName(opt.Fpt))
	for _, v := range opt.Fpd {
		si := smq.SmqIn{
			Fpd:  v,
			Ding: opt.Ding,
			IsS:  opt.IsS,
			Fpt:  opt.Fpt,
			Csk:  opt.Csk,
			As:   opt.AS,
		}
		so := si.Smq()
		if so.CodeLen == 0 {
			continue
		}
		h.AddResult(so, smq.GetFileName(v))
		output(so)

		if opt.IsO {
			var wb bytes.Buffer
			for i, v := range so.WordSlice {
				wb.WriteString(string(v))
				wb.WriteByte('\t')
				wb.WriteString(so.CodeSlice[i])
				wb.WriteByte('\n')
			}
			_ = os.Mkdir("result", 0666)
			err := ioutil.WriteFile(".\\result\\"+so.TextName+"_"+so.DictName+".txt", wb.Bytes(), 0666)
			if err != nil {
				fmt.Println("输出结果错误:", err)
			} else {
				fmt.Println("输出结果成功:", ".\\result\\"+so.TextName+"_"+so.DictName+".txt")
			}
		}
	}
	h.OutputHTMLFile("result.html")

	// time.Sleep(5 * time.Second)
}
