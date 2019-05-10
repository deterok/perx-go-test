package main

import (
	"perx-go-test/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"perx-go-test/lib"
	"perx-go-test/resources"
	"perx-go-test/views"
)

type Server struct {
	config *Config

	codeAlphabet []rune
	codeSize     int

	e  *echo.Echo
	db *gorm.DB
}

func NewServer(c *Config) *Server {
	return &Server{
		config:       c,
		codeAlphabet: []rune(c.Code.Alphabet),
		codeSize:     c.Code.Size,
	}
}

func (s *Server) Start() error {
	if err := s.initDB(); err != nil {
		panic(errors.Wrap(err, "db connect failed"))
	}

	if err := s.initWebServer(); err != nil {
		panic(errors.Wrap(err, "db connect failed"))
	}

	if err := s.e.Start(s.config.ListenAddress); err != nil {
		panic(errors.Wrap(err, "web server start failed"))
	}

	return nil
}

func (s *Server) initDB() error {
	c := s.config
	db, err := gorm.Open(c.DB.Dialect, c.DB.DSN)
	if err != nil {
		return err
	}

	if err := models.InitAllModels(db); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Server) initWebServer() error {
	e := echo.New()
	e.Use(middleware.Logger())
	if err := s.mountRoutes(e); err != nil {
		return err
	}

	s.e = e
	return nil
}

func (s *Server) mountRoutes(e *echo.Echo) error {
	codesRes := resources.NewCodesResource(s.db, s.codeAlphabet)

	codesView := views.NewCodesView(s.codeSize, codesRes)
	lib.BindModel(e, codesView, "/codes")

	codesCheckView := views.NewCodesCheckView(codesRes)
	lib.BindModel(e, codesCheckView, "/codes/check")

	return nil
}
