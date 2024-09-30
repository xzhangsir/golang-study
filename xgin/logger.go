package xgin

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("middle-[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
