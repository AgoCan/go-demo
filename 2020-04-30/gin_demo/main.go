package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/icza/session"
)

func middle(c *gin.Context) {
	m := map[string]interface{}{}
	r := c.Request
	w := c.Writer
	var templ = template.Must(template.New("").Parse(page))
	sess := session.Get(r)
	if sess != nil {
		// Already logged in
		if r.FormValue("Logout") != "" {
			session.Remove(sess, w) // Logout user
			sess = nil
		} else {
			sess.SetAttr("Count", sess.Attr("Count").(int)+1)
		}
	} else {
		// Not logged in
		if r.FormValue("Login") != "" {
			if userName := r.FormValue("UserName"); userName != "" && r.FormValue("Password") == "aDeexsdfeq234" {
				// Successful login. New session with initial constant and variable attributes:
				sess = session.NewSessionOptions(&session.SessOptions{
					CAttrs: map[string]interface{}{"UserName": userName},
					Attrs:  map[string]interface{}{"Count": 1},
				})
				session.Add(sess, w)
			} else {
				m["InvalidLogin"] = true
			}
		}
	}

	if sess != nil {
		m["UserName"] = sess.CAttr("UserName")
		m["Count"] = sess.Attr("Count")
	}
	if err := templ.Execute(w, m); err != nil {
		log.Println("Error:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"name": 1,
	})
}

func main() {
	session.Global.Close()
	session.Global = session.NewCookieManagerOptions(session.NewInMemStore(), &session.CookieMngrOptions{AllowHTTP: true})
	r := gin.Default()
	r.GET("/index", middle)
	r.POST("/index", middle)
	r.POST("login", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"he": 1,
		})
	})
	r.Run(":9000")
}

const page = `<html><body>
{{if .InvalidLogin}}<p style="color:red">Invalid user name or password!</p>{{end}}

{{if .UserName}}
	<p>Hello <b>{{.UserName}}</b>! Since login you visited <b>{{.Count}}</b> times! <a href="/index">Refresh!</a></p>
{{end}}

<form method="post" action="/index">
	{{if .UserName}}
		<input type="submit" name="Logout" value="Logout">
	{{else}}
		<label for="UserNameId" style="width:100px; display: inline-block">User name:</label>
		<input type="text" name="UserName" id="UserNameId"><br>
		<label for="PasswordId" style="width:100px; display: inline-block">Password:</label>
		<input type="password" name="Password" id="PasswordId">
		<span style="font-style:italic; font-size: 90%">Tip: use 'a' to login ;)</span><br>
		<input type="submit" name="Login" value="Login">
	{{end}}
</form>
</body></html>`
