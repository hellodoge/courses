package cached

import (
	"io/ioutil"
	"sync"
	"time"
)

type Cache struct {
	cached map[string]cached
	mutex  sync.Mutex
}

type cached struct {
	cachedAt time.Time
	file     string
}

func (c *Cache) Load(path string, lifetime time.Duration) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if cached, ok := c.cached[path]; ok {
		if cached.cachedAt.Add(lifetime).After(time.Now()) {
			return cached.file, nil
		}
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	c.cached[path] = cached{
		cachedAt: time.Now(),
		file:     string(file),
	}

	return string(file), nil
}
