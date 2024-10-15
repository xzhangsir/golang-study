package xgin

import (
	"fmt"
	"log"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			// recover可以捕获一个panic
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", message)
			}
		}()
		c.Next()
	}
}
