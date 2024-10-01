package xgin
import (
	"log"
	"fmt"
)

func Recovery()HandlerFunc{
	return func (c *Context){
		defer func(){
			if err := recover();err != nil{
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", message)
			}
		}()
		c.Next()
	}
}