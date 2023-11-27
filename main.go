package main

func main() {

	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandleFunc("POST", "/", func(c *Context) {

	})

}
