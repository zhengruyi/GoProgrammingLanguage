package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)
var sema = make(chan struct{},20)
func walkDir(dir string, n * sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir,entry.Name())
			go walkDir(subDir,n,fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) [] os.FileInfo {
	sema <- struct{}{}
	defer func() {<- sema}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr,"du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	//解析命令行的输入,尤其指各种不同的标志位
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string {"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root,&n,fileSizes)
	}
	//这里必须用另外一个协程，因为会被Wait挡住，然后
	//其余的协程则会因为无法往管道里写入数据而被阻塞,所以唯一的解决办法就是
	//另外起一个协程来监听程序的完成的情况
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	//对管道使用range，那么在管道为空时会阻塞，并且在管道关闭时会自动退出阻塞状态
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f GB\n",nfiles, float64(nbytes)/1e9)
}