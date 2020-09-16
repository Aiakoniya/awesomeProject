package main

type cache struct{
	storage  map[string] string
	speakLang speakLang
	capacity int
	maxCapacity int
}

func initCache(e speakLang) *cache{
	storage := make(map[string]string)
	return &cache{
		storage: storage,
		speakLang: e,
		capacity: 0,
		maxCapacity: 2,
	}
}

func (c *cache) setSpeakLang(e speakLang){
	c.speakLang = e
}
func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity{
		c.speak()
	}
	c.capacity++
	c.storage [key] = value
}
func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) speak() {
	c.speakLang.speak(c)
	c.capacity--
}
