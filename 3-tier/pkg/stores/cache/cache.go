package cache

import (
	"pet-store/pkg/models"
	"sync"
)

type Cache struct {
	cache map[int]*models.Pet
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		cache: map[int]*models.Pet{},
		mu:    sync.Mutex{},
	}
}

func (c *Cache) Get(id int) (*models.Pet, bool) {
	if pet, ok := c.cache[id]; !ok {
		return nil, false
	} else {
		return pet, true
	}
}

func (c *Cache) Set(pet *models.Pet) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[pet.ID] = pet
}

func (c *Cache) GetAllPet() ([]*models.Pet, error) {
	var pets []*models.Pet
	for _, v := range c.cache {
		pets = append(pets, v)
	}
	return pets, nil
}

func (c *Cache) DeleteAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = map[int]*models.Pet{}
}

func (c *Cache) Delete(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, id)
}
