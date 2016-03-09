package server

import (
	"fmt"
	"net/http"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	var index = `
<html>
  <head>
    <title>Secured APIs</title>
  </head>
  <body>
    <h1>Secured APIs</h1>
    <ul>
      <li>GET <code>/api/get</code></li>
      <li>POST <code>/api/post</code></li>
  </body>
</html>
`
	fmt.Fprintln(res, index)
}
