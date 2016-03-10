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
    <h4>Secured API endpoints for the <code>secureweb</code> application.</h4>
    <ul>
      <li>GET <code>/api/get</code></li>
      <li>POST <code>/api/post</code></li>
    </ul>
  </body>
</html>
`
	fmt.Fprintln(res, index)
}
