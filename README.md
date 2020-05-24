# go_auth_server

## introduction 

    This is an test project for build an auth server with gin framework

## how to use

### 1 install package
```golang===
go install
```
### 2 setup .env
```yaml===
PORT=${SERVER_PORT}
```

### 3 run
```golang===
go run main.go
```

### 4 build
```golang===
go build
```

### gracdful shutdown features

在 go 1.8之後

gin 提供一個 graceful shutown的method可以讓

開發者在關閉server連線之前

等待處理完使用中的連線 才中斷連線的功能

實做如下：

假設 有一個要處理 5s的 router

我們把listenAndServe放到一個goroutine 處理 避免ctl+C結束 mainthread時 影響連線

建立一個 quit channel監聽 os signal

直到 發出 syscall.SIGINT, syscall.SIGTERM

server等待處理完才 執行shutdown// WithTimeout 作用

***要注意的是*** 這邊 context.WithTimeout的第二個參數 時間

是server最多等多久的時間 因此建議可以設定長一點



```golang===
PORT := os.Getenv("PORT")
	// setup Default Router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome Gin Server",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// setup quit channel to receive system shutdown
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
```