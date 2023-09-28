package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

func InitTemplate(templates []string) *Template {
	// Templates
	t := &Template{
		templates: ParseTemplates(templates),
	}
	return t
}

type Template struct {
	templates *template.Template
}

func ParseTemplates(paths []string) *template.Template {
	templ := template.New("")
	var err error
	for _, path := range paths {
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, ".html") {
				_, err = templ.ParseFiles(path)
				if err != nil {
					log.Println(err)
				}
			}
			return err
		})
	}
	if err != nil {
		panic(err)
	}
	return templ
}

var TData *TemplateData

type TemplateData struct {
	Data map[string]interface{}
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		Data: make(map[string]interface{}),
	}
}

func (td *TemplateData) AddTplData(data map[string]interface{}) {
	for k, v := range data {
		td.Data[k] = v
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var buf bytes.Buffer
	// add to data TData
	TData.AddTplData(data.(map[string]interface{}))
	err := t.templates.ExecuteTemplate(&buf, name, TData.Data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = buf.WriteTo(w)
	return err
}
