package main

import (
	"context"
	"fmt"
	"time"
)

func monitor(ctx context.Context, num int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("监控器%v, 接收到通道值为：%v，监控结束。\n", v, num)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", num)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go monitor(ctx, 1)
	go monitor(ctx, 2)
	go monitor(ctx, 3)
	go monitor(ctx, 4)
	go monitor(ctx, 5)
	time.Sleep(time.Second)
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("main over")
}
