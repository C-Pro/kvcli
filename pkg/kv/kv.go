package kv

type TX struct {
	Parent *TX
	Data   map[string]string
}

func NewTX(parent *TX) *TX {
	tx := &TX{
		Data:   make(map[string]string),
		Parent: parent,
	}

	// Insanly ineffective way, but will work for CLI use case
	if parent != nil {
		for k, v := range parent.Data {
			tx.Data[k] = v
		}
	}

	return tx
}

func (tx *TX) Set(key, value string) {
	tx.Data[key] = value
}

func (tx *TX) Get(key string) *string {
	val, ok := tx.Data[key]
	if ok {
		return &val
	}

	return nil
}

func (tx *TX) Delete(key string) {
	delete(tx.Data, key)
}

func (tx *TX) Count(value string) int {
	cnt := 0
	for _, v := range tx.Data {
		if v == value {
			cnt++
		}
	}

	return cnt
}
