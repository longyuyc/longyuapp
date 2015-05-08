package main

import (
	//"github.com/go-martini/martini"
	"longyuapp/api/createQRCodes"
	//"longyuapp/api/exportExcel"
	//"longyuapp/api/sendEmail"
	//"longyuapp/api/user"
	//"net/http"
)

func main() {
	createQRCodes.GetQrCode()
	/*
		//创建一个典型的martini实例
		m := martini.Classic()
		m.Use(martini.Static("/api/public"))
		m.Get("/", LoginFunc)
		//收对\的GET方法请求，第二个参数是对一请求的处理方法。
		m.Group("/api", func(r martini.Router) {
			r.Group("/user", user.UserApiRoute)
			r.Group("/sendEmail", sendEmail.SendEmailApiRoute)
			r.Group("/export", exportExcel.ExportExcelApiRoute)
		})
		//运行服务器
		m.Run()
	*/
}

/*
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "html/ace/login.html", http.StatusFound)
	}
}
*/
