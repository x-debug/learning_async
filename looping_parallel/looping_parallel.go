package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func main() {
	filenames := []string{"./imgs/img01.jpg", "./imgs/img02.jpg", "./imgs/img03.jpg"}
	//files, err := makeThumbnails5(filenames)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println(files)

	var wg sync.WaitGroup
	ch := make(chan string)
	result := make(chan int64)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, f := range filenames {
			ch <- f
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		size := makeThumbnails6(ch)
		result <- size
	}()

	log.Println(<-result)
	wg.Wait()
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(file string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(file)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
