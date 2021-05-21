package repository

import (
	cachedStatic "github.com/hellodoge/courses-tg-bot/pkg/static/cached"
	"path/filepath"
	"time"
)

const (
	queriesFolder = "queries"

	cachedQueryLifetime = 5 * time.Second
)

var (
	queriesCache = cachedStatic.Cache{}
)

func getQuery(folder, filename string) (string, error) {
	path := filepath.Join(queriesFolder, folder, filename)
	query, err := queriesCache.Load(path, cachedQueryLifetime)
	return query, err
}
