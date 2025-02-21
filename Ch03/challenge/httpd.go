package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	loginHTML = `<!DOCTYPE html>
<html>
	<body>
		<form method="post">
			<h2>Please Login</h2>
			User: <input name="user"> <br/>
			Password: <input type="password" name="passwd"> <br/>
			<input type="submit"/>
		</form>
	<body>
</html>
`
	loginHTMLTemplate = template.Must(template.New("login").Parse(loginHTML))

	statusHTML = `<!DOCTYPE html>
<html>
	<body>
		<h2>Status</h2>
		%s
	</body>
</html>
`

	statusHTMLTemplate = template.Must(template.New("status").Parse(statusHTML))
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method != http.MethodPost {
		fmt.Fprint(w, loginHTMLTemplate)
		return
	}

	user, passwd := r.FormValue("user"), r.FormValue("passwd")
	if !authUser(user, passwd) {
		http.Error(w, fmt.Sprintf("%s:%s - bad login", user, passwd), http.StatusUnauthorized)
		return
	}

	statusHTMLTemplate.Execute(w, statusHTML)
	fmt.Fprintf(w, "end")
}

func main() {
	http.HandleFunc("/status", statusHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
