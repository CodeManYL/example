package main

import (
	"context"
	"example/gomicro/biz"
	"example/gomicro/data"
	"fmt"
	"github.com/xormplus/xorm"
	"net/http"
	"sync"
	"time"
)
import 	_ "github.com/go-sql-driver/mysql"


//导航
type WebNewNavigate struct {
	Index string `xorm:"varchar(32) pk notnull 'Index'"`
	Title string `xorm:"varchar(32) pk notnull 'Title'"`
	Icon string `xorm:"varchar(32)  notnull 'Icon'"`
}

var engine *xorm.Engine

func InitModEngine(dbName, dbAddress string) (err error) {
	engine, err = xorm.NewEngine(dbName, dbAddress)
	if err != nil {
		return
	}
	return nil
}


func Server (add string,handler http.Handler,stop <- chan struct{},wg *sync.WaitGroup) error {
	s := &http.Server{Addr: add,Handler: handler}

	go func() {
		defer wg.Done()
		<- stop
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		s.Shutdown(ctx)
	}()

	return s.ListenAndServe()
}


//func main (){
//	done := make(chan error,2)
//	stop := make(chan struct{})
//
//	r := gin.Default()
//	r.GET("/", func(c *gin.Context) {
//		//模拟有未处理完的请求
//		fmt.Println("/ 进来请求了")
//		time.Sleep(5*time.Second)
//		fmt.Println("/ 结束了")
//	})
//
//
//	tmp := sync.WaitGroup{}
//	tmp.Add(1)
//	go func() {
//		done <- Server(":8081",r,stop,&tmp)
//	}()
//
//	go func() {
//		//模拟其中一个服务挂了
//		time.Sleep(5*time.Second)
//		done <- errors.New("create errr")
//		//done <- Server(":8081",r,stop)
//	}()
//
//	var stopped bool
//	for i := 0; i < cap(done); i ++ {
//		err := <- done
//		if err != nil {
//			fmt.Println("服务挂了",err)
//		}
//		if !stopped {
//			stopped = true
//			close(stop)
//		}
//	}
//	tmp.Wait()
//}



type Tracker struct {
	ch chan string
	stop chan struct {}
}

func NewTracker() *Tracker {
	return &Tracker{
		make(chan string,10),
		make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context,data string) error {

	select {
	case t.ch <- data:
		//数据满了可以报错
		fmt.Println("加入数据")
		return nil
	case <- ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1*time.Second)
		fmt.Println("res",data) //业务处理逻辑
	}

	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		//这里就可以感知到run方法退出了
		fmt.Println("感知run已经退出")
	case <- ctx.Done():
		//如果Run一直没退出 这这里做超时控制
	}
}

//func main (){
//	tracker := NewTracker()
//	go tracker.Run()
//
//	for i := 0 ;i <3;i++{
//		if err := tracker.Event(context.Background(),"111");err != nil {
//			fmt.Println(err)
//		}
//	}
//	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(2*time.Second))
//	defer cancel()
//
//	//一瞬间就执行到这里了
//	tracker.Shutdown(ctx)
//}

type Goods struct {
	Name string
	Num int
}


func main (){
	//engine, err := xorm.NewEngine("mysql", "root:33css@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	//if err != nil {
	//	panic(err)
	//}
	//
	//var categories []*Goods
	//
	//err = engine.Find(&categories)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _,i := range categories {
	//	fmt.Println(*i)
	//}
	//fmt.Println(categories)
	//

	orderRepo := data.NewOrderRepo()
	orderUsercase := biz.NewOrderUsercase(orderRepo)
	fmt.Println(orderUsercase)
}