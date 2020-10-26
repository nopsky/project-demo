/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/25 10:08
 */
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.router.Use(middleware...)
}

func (s *Server) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.router.Group(relativePath, handlers...)
}

// POST is a shortcut for router.Handle("POST", path, handle).
func (s *Server) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (s *Server) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (s *Server) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (s *Server) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (s *Server) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (s *Server) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (s *Server) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.HEAD(relativePath, handlers...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (s *Server) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.router.Any(relativePath, handlers...)
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (s *Server) StaticFile(relativePath, filepath string) gin.IRoutes {
	return s.router.StaticFile(relativePath, filepath)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func (s *Server) Static(relativePath, root string) gin.IRoutes {
	return s.router.Static(relativePath, root)
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (s *Server) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return s.router.StaticFS(relativePath, fs)
}
