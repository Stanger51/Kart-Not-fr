package response

import (
	"encoding/json"
	"sync"
)

type ResponseItem struct {
	OldID int64 `json:"oldId"`
	NewID int64 `json:"newId"`
}

type Response struct {
	cache       []ResponseItem
	mutex       sync.RWMutex
	onItemAdded func(i ResponseItem)
}

func New(onItemAdded ...func(i ResponseItem)) *Response {
	var callback func(i ResponseItem)
	if len(onItemAdded) > 0 {
		callback = onItemAdded[0]
	}

	return &Response{
		cache:       make([]ResponseItem, 0),
		onItemAdded: callback,
	}
}

func (r *Response) AddItem(i ResponseItem) {
	r.mutex.Lock()

	r.cache = append(r.cache, i)
	if r.onItemAdded != nil {
		go r.onItemAdded(i)
	}

	r.mutex.Unlock()
}

func (r *Response) Clear() {
	r.mutex.Lock()
	r.cache = make([]ResponseItem, 0)
	r.mutex.Unlock()
}

func (r *Response) EncodeJSON(e *json.Encoder) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return e.Encode(r.cache)
}

func (r *Response) Len() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.cache)
}
