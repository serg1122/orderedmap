package orderedmap

type OrderedMap struct {
	firstKey string
	lastKey  string
	pairs    map[string]*Node
}

func CreateOrderedMap() *OrderedMap {

	return &OrderedMap{
		pairs: make(map[string]*Node),
	}
}

func (m OrderedMap) Length() int {

	return len(m.pairs)
}

func (m OrderedMap) HasKey(key string) bool {

	_, hasKey := m.pairs[key]

	return hasKey
}

func (m OrderedMap) FirstNode() (*Node, error) {

	if m.Length() == 0 {
		return nil, CreateErrorChainIsEmpty("")
	}

	return m.pairs[m.firstKey], nil
}

func (m OrderedMap) LastNode() (*Node, error) {

	if m.Length() == 0 {
		return nil, CreateErrorChainIsEmpty("")
	}

	return m.pairs[m.lastKey], nil
}

func (m OrderedMap) GetNode(key string) (*Node, bool) {

	node, isExist := m.pairs[key]

	if isExist {
		return node, true
	}

	return nil, false
}

func (m *OrderedMap) Set(key string, value interface{}) {

	if m.Length() == 0 {
		m.init(key, value)
		return
	}

	if m.HasKey(key) {
		m.pairs[key].SetValue(value)
		return
	}

	m.append(key, value)
}

func (m *OrderedMap) init(key string, value interface{}) {

	m.pairs[key] = &Node{
		prev:  nil,
		next:  nil,
		value: value,
	}
	m.lastKey = key
	m.firstKey = key
}

func (m *OrderedMap) append(key string, value interface{}) {

	m.pairs[key] = &Node{
		prev:  m.pairs[m.lastKey],
		next:  nil,
		value: value,
	}
	m.lastKey = key
}

func (m *OrderedMap) Get(key string) (interface{}, bool) {

	if m.HasKey(key) {
		return m.pairs[key].Value(), true
	}

	return nil, false
}

func (m *OrderedMap) Delete(key string) bool {

	hasKey := m.HasKey(key)

	if hasKey {
		newPrev := m.pairs[key].prev
		newNext := m.pairs[key].next

		newPrev.SetNext(newNext)
		newNext.SetPrev(newPrev)

		delete(m.pairs, key)

		return true
	}

	return hasKey
}
