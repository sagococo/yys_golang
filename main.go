package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

var shotSquare image.Rectangle

var pause time.Duration

func main() {
	t := time.Now().Unix()
	tupoT := time.Now().Unix()
	tupoNum := 0
	zd := 0
	shotSquare = image.Rect(1766, 54, 2555, 643)
	for {
		capture(shotSquare)
		for _, v := range []string{"call", "autoagree", "agree", "yunhunagree", "huodong", "dashe", "jixu", "attack", "users", "yuling", "yeyuanhuo", "tiaozhan"} {
			if templateFile(v) {
				if v == "jixu" && time.Now().Unix()-tupoT > 9 && templateFile("tupo") {
					tupoNum++
					tupoT = time.Now().Unix()
					if tupoNum >= 25 {
						go play(fmt.Sprintf("累计获得%v张突破券", tupoNum))
					}
				}
				if v == "users" {
					pause = 500 * time.Millisecond
				} else {
					pause = 200 * time.Millisecond
				}
				if time.Now().Unix()-t > 9 {
					zd++
					log.Printf("已经战斗%v次，本次战斗耗时%v秒，累计获得%v张突破券\n", zd, time.Now().Unix()-t, tupoNum)
					t = time.Now().Unix()
				}
				break
			}
		}
		time.Sleep(pause)
	}
}

func mains() {
	play("audios/alreadyStop.mp3")
	play("audios/10.mp3")
}
