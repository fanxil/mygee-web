package main

import (
	"gee"
	"net/http"
)

// var wg sync.WaitGroup
//
//	func download(url string) {
//		fmt.Println("start downloand", url)
//		time.Sleep(time.Second)
//		wg.Done()
//	}
//
// var ch = make(chan string, 10) //创建大小为10的缓冲信道
//
//	func download2(url string) {
//		fmt.Println("start download2", url)
//		time.Sleep(time.Second)
//		ch <- url // 将url发送给信道，即队列
//
// }
//type Engine struct {
//}

// 定义Engine的方法： *：指针传递
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Printf("URL.Path = %q\n", req.URL.Path)
//		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//	}
//}

func main() {
	//obj = map[string]interface{}{
	//	"name": "geektutu",
	//	"password": "1234",
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//encoder := json.NewEncoder(w)
	//if err := encoder.Encode(obj); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}

	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

	//vmap := make(map[string]string, 1)
	//vmap["x"] = "xxx"
	//vmap["xx"] = "xxx"
	//vmap["xxx"] = "xxxx"
	//fmt.Println(len(vmap))
	//fmt.Println(vmap["x"])
	//
	//r := gee.New()
	//r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	//	for k, v := range req.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//
	//r.Run(":9999")

	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go download("a.com/" + string(i+'0'))
	//}
	//wg.Wait()
	//fmt.Println("Done!")
	//
	//for i := 0; i < 3; i++ {
	//	go download2("a.com/" + string(i+'0'))
	//}
	//for i := 0; i < 3; i++ {
	//	msg := <-ch // 等待信道返回消息。
	//	fmt.Println("finish2", msg)
	//}
	//fmt.Println("Done2!")
	////
	//engine := new(Engine)
	//log.Fatal(http.ListenAndServe(":9999", engine))

	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/hello", helloHandler)
	////第一个参数是地址，:9999表示在 9999 端口监听。而第二个参数则代表处理所有的HTTP请求的实例，
	////nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口
	//log.Fatal(http.ListenAndServe(":9999", nil))
}

//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//	}
//}
