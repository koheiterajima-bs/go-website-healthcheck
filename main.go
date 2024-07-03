package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "siteHealthChecker",             // プログラムの名前
		Usage: "use to check web site health.", // プログラムについての記述
		Flags: []cli.Flag{ // 解析するフラグのリスト
			&cli.StringFlag{ // 文字列を受け取るフラグ
				Name:     "url",         // フラグ名
				Aliases:  []string{"u"}, // フラグ名の別名
				Usage:    "url name",    // 使用用途
				Required: true,          // このフラグが必須かどうか
			},
		},
		// Actionフィールドは、CLIアプリケーションでフラグや引数が解析された後に実行される関数を指定するために使用される
		Action: func(c *cli.Context) error {
			url := c.String("url") // cli.StringFlagのNameを指定

			port := "80"

			// タイムアウト時間を設定
			timeout := 5 * time.Second

			// TCP接続を試みる
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(url, port), timeout)
			if err != nil {
				fmt.Printf("接続エラー：%v\n", err)
			}
			defer conn.Close()

			// ローカルアドレスについて
			localAddr := conn.LocalAddr().String()
			// リモートアドレスについて
			remoteAddr := conn.RemoteAddr().String()

			fmt.Printf("接続成功：あなたの入力したホスト名は「%s」\nローカルアドレスは%v\nリモートアドレスは%v", url, localAddr, remoteAddr)
			return nil
		},
	}

	err := app.Run(os.Args) // appの実行
	if err != nil {
		log.Fatal(err)
	}
}
