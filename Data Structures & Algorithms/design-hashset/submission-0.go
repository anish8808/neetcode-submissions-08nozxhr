type MyHashSet struct {
	myMap map[int]bool
}

func Constructor() MyHashSet {
    maps := make(map[int]bool)
	return MyHashSet{myMap : maps}
}

func (this *MyHashSet) Add(key int) {
    this.myMap[key] = true
}

func (this *MyHashSet) Remove(key int) {
	ok := this.myMap[key]
	if (ok){
		delete(this.myMap , key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
    ok := this.myMap[key]
	if (ok){
		return true
	}
	

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 