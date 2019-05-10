package lib

import "github.com/labstack/echo"

type Getter interface {
	Get(c echo.Context) error
}

type Poster interface {
	Post(c echo.Context) error
}

type Puter interface {
	Put(c echo.Context) error
}

type Deleter interface {
	Delete(c echo.Context) error
}

type Optioner interface {
	Options(c echo.Context) error
}

func BindModel(e *echo.Echo, view interface{}, path string) {
	if v, ok := view.(Getter); ok {
		e.GET(path, v.Get)
		return
	}

	if v, ok := view.(Poster); ok {
		e.POST(path, v.Post)
		return
	}

	if v, ok := view.(Puter); ok {
		e.PUT(path, v.Put)
		return
	}

	if v, ok := view.(Deleter); ok {
		e.DELETE(path, v.Delete)
		return
	}

	if v, ok := view.(Optioner); ok {
		e.OPTIONS(path, v.Options)
		return
	}

	panic("view does not implement for any interface")
}
