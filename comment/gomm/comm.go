package main

import (
	"flag"
	"github.com/MWangxj/util/comment"
	"log"
	"os"
	"sync"
)

// main
func main() {
	pkg := flag.String("dir", ".", "USAGE [-dir = .] to choose package file")
	flag.Parse()
	if *pkg == "" {
		os.Exit(233)
	}
	fs, err := comment.GetGoFileListByDir(*pkg)
	if err != nil {
		log.Println(err.Error())
		os.Exit(234)
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(fs))
	for _, fn := range fs {
		go func(fn string) {
			if err := comment.GetFileLine(fn); err != nil {
				log.Println(err.Error())
			}
			wg.Done()
		}(fn)
	}
	wg.Wait()
	log.Println("ok,done")
}
