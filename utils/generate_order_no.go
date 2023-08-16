package utils

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// GenerateOrderNo 生成商户订单号 流水号
func GenerateOrderNo() string {
	var t = time.Now()
	var num int64
	s := t.Format("20060102150405")              // 年 月 天 时 分 秒
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3 // 毫秒 部分
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// sup 将一个整数转换为指定长度的字符串
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
