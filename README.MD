# Embed Vue.js file as String into your Go code

# Usage

main.go
```go
package main

import (
	"fmt"
	"github.com/StevenZack/vue"
	"net/http"
)

func main() {
	http.HandleFunc("/js/vue.js", vue.HandlerFn)
	http.HandleFunc("/", home)
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println("serve error:", e)
		return
	}

}
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
```

index.html
```html
<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
<div id="app">
    <p>{{message}}</p>
</div>
<script type="text/javascript" src="/js/vue.js"></script>
<script type="text/javascript">
    var app=new Vue({
        el:'#app',
        data:{
            message:'hello'
        }
    })
</script>
</body>
</html>
```