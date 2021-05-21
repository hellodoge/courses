package cached

import (
	"io/ioutil"
	"time"
)

type Cache map[string]cached

type cached struct {
	cachedAt time.Time
	file string
}

func (c Cache) Load(path string, lifetime time.Duration) (string, error) {
	if cached, ok := c[path]; ok {
		if cached.cachedAt.Add(lifetime).After(time.Now()) {
			return cached.file, nil
		}
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	c[path] = cached{
		cachedAt: time.Now(),
		file:     string(file),
	}

	return string(file), nil
}