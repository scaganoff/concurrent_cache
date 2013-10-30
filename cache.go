// Package for Authentication, Authorization and Access Control
package cache

const (
	ADD = 1
	DEL = 2
	GET = 3
)

// The cache has in input channel for requests and an output channel for responses
// All activity is serialised through the input channel effectively acting as a mutex.
type ConcurrentCache struct {
	req   chan request
	resp  chan interface{}
	cache map[string]interface{}
}

// Request is what comes in on the input channel.
type request struct {
	action int
	key    string
	value  interface{}
}

// Constructer spawns a new goroutine with the main processing loop and
// returns a pointer to the cache.
func New() *ConcurrentCache {
	d := &ConcurrentCache{
		req:   make(chan request, 10),     // buffered input channel
		resp:  make(chan interface{}, 10), // buffered output channel for GET
		cache: make(map[string]interface{}),
	}
	go d.run()
	return d
}

// Public wrapper to add an entry to the cache via the input channel.
func (d *ConcurrentCache) Add(k string, v interface{}) {
	r := request{
		action: ADD,
		key:    k,
		value:  v,
	}
	d.req <- r
}

// Public wrapper to get an entry from the cache via the input/output channel.
func (d *ConcurrentCache) Get(k string) interface{} {
	r := request{
		action: GET,
		key:    k,
	}
	d.req <- r
	return <-d.resp
}

// Public wrapper to delete an entry from the cache via the input channel.
func (d *ConcurrentCache) Delete(k string) {
	r := request{
		action: DEL,
		key:    k,
	}
	d.req <- r
}

// Main message loop waits on the input channel and then executes the request
// when it arrives
func (d *ConcurrentCache) run() {
	for {
		req := <-d.req

		switch req.action {
		case ADD:
			d.cache[req.key] = req.value
		case DEL:
			delete(d.cache, req.key)
		case GET:
			d.resp <- d.cache[req.key]
		}
	}
}
