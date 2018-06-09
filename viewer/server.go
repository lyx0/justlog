package viewer

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type templateRenderer struct {
	templates *template.Template
}

// Server api server
type Server struct {
	port string
}

// NewServer create Server
func NewServer() Server {
	return Server{
		port: ":8030",
	}
}

// Init api server
func (s *Server) Init() {

	e := echo.New()
	e.HideBanner = true

	renderer := &templateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tpl")),
	}
	e.Renderer = renderer

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "something.tpl", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	fmt.Println("Starting viewer on port " + s.port)
	e.Logger.Fatal(e.Start(s.port))
}

// Render renders a template document
func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
