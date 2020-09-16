package main

func main() {
	kazakh := &kazakh{}
	cache := initCache(kazakh)
	cache.add("a", "1")
	cache.add("b","2")
	cache.add("c","3")
	english := &english{}
	cache.setSpeakLang(english)
	cache.add("d","4")
	russian := &russian{}
	cache.setSpeakLang(russian)
	cache.add("e","5")
}
