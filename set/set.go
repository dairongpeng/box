package set

type Item struct {
	Key   string
	Value interface{}
}

type Set struct {
	data map[string]interface{}
}

// NewSet 初始化一个集合
func NewSet(items []*Item) *Set {
	s := &Set{make(map[string]interface{})}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add 往集合中追加一个元素
func (s *Set) Add(item *Item) {
	s.data[item.Key] = item.Value
}

// Remove 移除集合中当前key的元素
func (s *Set) Remove(key string) {
	delete(s.data, key)
}

// Contains 检查集合中是否包含当前元素
func (s *Set) Contains(key string) bool {
	_, exists := s.data[key]
	return exists
}

// Any 获取集合中任意一个元素
func (s *Set) Any() string {
	for item := range s.data {
		return item
	}
	return ""
}

// Intersection 两个集合做交集
func Intersection(s1, s2 *Set) *Set {
	intersection := &Set{make(map[string]interface{})}
	for k := range s1.data {
		if s2.Contains(k) {
			itm := &Item{
				Key:   k,
				Value: s2.data[k],
			}
			intersection.Add(itm)
		}
	}
	return intersection
}

// Union 两个集合做并集
func Union(s1, s2 *Set) *Set {
	union := NewSet(nil)
	for key, value := range s1.data {
		itm := &Item{
			Key:   key,
			Value: value,
		}
		union.Add(itm)
	}
	for key, value := range s2.data {
		itm := &Item{
			Key:   key,
			Value: value,
		}
		union.Add(itm)
	}
	return union
}

// Difference 两个集合做差集。(s1 - s2)
func Difference(s1, s2 *Set) *Set {
	difference := &Set{make(map[string]interface{})}
	for key, value := range s1.data {
		if !s2.Contains(key) {
			itm := &Item{
				Key:   key,
				Value: value,
			}
			difference.Add(itm)
		}
	}
	return difference
}
