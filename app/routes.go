package app

import (
	"bufio"
	"net"
	"strconv"

	"github.com/Singert/xjtu_cnlab/core/router"
)

// RegisterAppRoutes 注册应用路由
func RegisterAppRoutes(r *router.Router) {
	r.RegisterRoute("GET", "/hello", "discription", HandleHello)
	r.RegisterRoute("POST", "/upload", "discription", HandleUpload)

	r.RegisterGroupRoute("/admin", func(g *router.Group) {
		g.RegisterRoute("GET", "/reload", "discription", HandleAdminReload)
		g.RegisterRoute("GET", "/download-logs", "discription", HandleDownloadLogs)

	})
	r.RegisterRoute("GET", "logs", "discription", HandleLogs)
	r.RegisterGroupRoute("/debug", func(g *router.Group) {
		g.RegisterRoute("GET", "/", "discription", HandleDebugRoutes)
		g.RegisterRoute("GET", "/json", "discription", HandleDebugRoutesJSON)
		g.RegisterRoute("GET", "/routes", "discription", HandleDebugRoutesSmart)
		g.RegisterRoute("GET", "/update-route", "discription", HandleUpdateRoute)
		g.RegisterRoute("GET", "/dashboard", "discription", HandleDebugDashboard)
		g.RegisterRoute("GET", "/reload", "discription", HandleAdminReload)
		g.RegisterRoute("GET", "/download-logs", "discription", HandleDownloadLogs)
	})

}

func HandleHello(ctx *router.Context) {
	conn := ctx.Conn.(net.Conn)
	writer := bufio.NewWriter(conn)

	body := []byte("Hello World!")

	writer.WriteString("HTTP/1.1 200 OK\r\n")
	writer.WriteString("Content-Type: text/plain\r\n")
	writer.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	writer.WriteString("\r\n")
	writer.Write(body)
	writer.Flush()
}

func HandleUpload(ctx *router.Context) {
	conn := ctx.Conn.(net.Conn)
	writer := bufio.NewWriter(conn)

	body := []byte("Upload successful (dummy)")

	writer.WriteString("HTTP/1.1 200 OK\r\n")
	writer.WriteString("Content-Type: text/plain\r\n")
	writer.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	writer.WriteString("\r\n")
	writer.Write(body)
	writer.Flush()
}

func HandleLogin(ctx *router.Context) {
	conn := ctx.Conn.(net.Conn)
	writer := bufio.NewWriter(conn)

	body := []byte("Login successful (dummy)")

	writer.WriteString("HTTP/1.1 200 OK\r\n")
	writer.WriteString("Content-Type: text/plain\r\n")
	writer.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	writer.WriteString("\r\n")
	writer.Write(body)
	writer.Flush()
}

func HandleRegister(ctx *router.Context) {
	conn := ctx.Conn.(net.Conn)
	writer := bufio.NewWriter(conn)

	body := []byte("Register successful (dummy)")

	writer.WriteString("HTTP/1.1 200 OK\r\n")
	writer.WriteString("Content-Type: text/plain\r\n")
	writer.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	writer.WriteString("\r\n")
	writer.Write(body)
	writer.Flush()
}
