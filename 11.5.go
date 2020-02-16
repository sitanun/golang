package main

import (
	"fmt"
	"sync"
	"time"
)

func say(txt string, sleep time.Duration, wg)
  defer wg.Done()
  fmt.Println(txt)
  time.sleep(time.Millisecond * sleep)