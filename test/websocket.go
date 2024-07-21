/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-21 11:12:08
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 11:17:15
 * @FilePath: /easy-chat/test/websocket.go
 * @Description:  learn
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理websocket(基于http升级)
func serverWs(w http.ResponseWriter, r *http.Request) {
	// 进行升级
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", serverWs)
	fmt.Println("start websocket")
	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}
