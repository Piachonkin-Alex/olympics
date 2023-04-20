package rendener

import (
	"bytes"
	"context"
	"html/template"
	"io/fs"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type cacheRendener struct {
	mx    sync.RWMutex
	fs    fs.FS
	cache map[string]*template.Template
}

var _ Renderer = (*cacheRendener)(nil)

func NewCacheRendener(fs fs.FS) Renderer {
	return &cacheRendener{
		fs:    fs,
		cache: make(map[string]*template.Template),
	}
}

func (r *cacheRendener) loadTemplate(cacheKey string, loadFilesNames []string) (*template.Template, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	var resTpl *template.Template
	for _, file := range loadFilesNames {
		fileData, err := fs.ReadFile(r.fs, file)
		if err != nil {
			return nil, errors.Errorf("no file %s found", file)
		}

		tpl := template.New(file)
		if resTpl == nil {
			resTpl = template.New(file)
		}

		if _, err = tpl.Parse(string(fileData)); err != nil {
			return nil, err
		}

		resTpl = resTpl.New(file)
	}

	r.cache[cacheKey] = resTpl
	return resTpl, nil
}

func (r *cacheRendener) Render(ctx context.Context, data any, key string, addKeys ...string) (string, error) {
	keys := append(addKeys, key)
	mapKey := strings.Join(keys, "\n")
	var err error
	r.mx.RLock()
	tpl, ok := r.cache[mapKey]
	r.mx.RUnlock()
	if !ok {
		if tpl, err = r.loadTemplate(mapKey, keys); err != nil {
			return "", err
		}
	}

	var buf bytes.Buffer

	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", errors.Errorf("couldn't render template: %v", err)
	}

	return buf.String(), nil

}
