package grest

import (
	"io"
	"net/http"
)

const (
	get    = "GET"
	post   = "POST"
	put    = "PUT"
	delete = "DELETE"
	patch  = "PATCH"
)

// IndexAction Displays all resources
// path /resources [method GET]
type IndexAction interface {
	Index(w http.ResponseWriter, req *http.Request)
}

// ShowAction Displays a specific resource
// path /resources/:id [method GET]
type ShowAction interface {
	Show(w http.ResponseWriter, req *http.Request)
}

// UpdateAction Updates a specific resource
// path /resources/:id [method PUT or PATH]
type UpdateAction interface {
	Update(w http.ResponseWriter, req *http.Request)
}

// CreateAction Creates a new resource
// path /resources [method POST]
type CreateAction interface {
	Create(w http.ResponseWriter, req *http.Request)
}

// DestroyAction Deletes a specific resource
// path /resources/:id [method DELETE]
type DestroyAction interface {
	Destroy(w http.ResponseWriter, req *http.Request)
}

//Controller implement default actions
type Controller struct {
}

//New Controller
func New() *Controller {
	controller := Controller{}
	return &controller
}

func getID(pattern string, path string) (id string, hasID bool) {
	if pattern == path || pattern+"/" == path {
		return "", false
	}
	pattern += "/"
	n := len(pattern)
	return path[n:len(path)], true
}

func (c *Controller) resourcesHandler(actions interface{}, pattern string) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		path := req.URL.Path

		var handler http.HandlerFunc
		id, hasID := getID(pattern, path)
		switch req.Method {
		case get:
			// Index or Show
			if action, ok := actions.(IndexAction); ok && !hasID {
				handler = action.Index
			} else if action, ok := actions.(ShowAction); ok && hasID {
				req.Form.Set("id", id)
				handler = action.Show
			}
		case post:
			if action, ok := actions.(CreateAction); ok && !hasID {
				handler = action.Create
			}
		case put:
			if action, ok := actions.(UpdateAction); ok && hasID {
				req.Form.Set("id", id)
				handler = action.Update
			}
		case patch:
			if action, ok := actions.(UpdateAction); ok && hasID {
				req.Form.Set("id", id)
				handler = action.Update
			}
		case delete:
			if action, ok := actions.(DestroyAction); ok {
				req.Form.Set("id", id)
				handler = action.Destroy
			}
		}

		if handler == nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(w, req)
	}
}

//Resources add default actions routes (Index, Show, Update, Create, Delete)
func (c *Controller) Resources(path string, actions interface{}) {
	n := len(path)
	if path[n-1] == '/' {
		path = path[0 : n-1]
	}

	http.HandleFunc(path, c.resourcesHandler(actions, path))
	http.HandleFunc(path+"/", c.resourcesHandler(actions, path))
}

//este sería la raíz "/"
func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "controller: "+r.Method)
}
