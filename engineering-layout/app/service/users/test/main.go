
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/wire"
	"sync"
	"time"
)

type Data struct {
	name string
}

func (d *Data) Get(){
	fmt.Println("11111")
}


type Biz struct {
	DataRepo
}

type Service struct {
	*Biz
}

type DataRepo interface {
	Get()
}

func NewData (name,age string) *Data{
	fmt.Println(age)
	return &Data{name}
}


func NewBiz(data DataRepo) (*Biz,error){
	return  &Biz{data},errors.New("create err")
}

func NewS(biz *Biz) *Service {
	return  &Service{biz}
}

var MockUserRepoSet = wire.NewSet(NewData,wire.Bind(new(DataRepo), new(*Data)))





func shutdown(){
	ch := make(chan struct{})

	c, cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()

	lock := sync.WaitGroup{}


	lock.Add(1)
	go func() {
		defer lock.Done()
		time.Sleep(time.Second*15)
	}()


	go func(){
		lock.Wait()
		fmt.Println("处理完了1")
		close(ch)
		fmt.Println("处理完了2")
	}()

	select {
	case <- c.Done():
		fmt.Println("超时")
		return
	case <- ch:
		fmt.Println("成功")
		return
	}

}


func ticker(){
	tick := time.NewTicker(2*time.Second)


	for {
		select {
			case <- tick.C:
				fmt.Println("1秒钟了")
		}

	}

}


func main() {
	ticker()
}