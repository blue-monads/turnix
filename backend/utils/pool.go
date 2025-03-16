package xutils

import (
	"sync"

	"github.com/rs/zerolog"
)

// Element is the interface that all poolable elements must implement
type Element interface {
	Reset()
	GetID() int64
}

// Pool is a generic pool for reusable elements
type Pool[T Element] struct {
	logger zerolog.Logger
	pool   sync.Pool

	idCacheIndex map[int64]T
	currentCount int64
	maxCount     int64
	mutex        sync.Mutex
}

// New creates a new Pool with the specified factory function
func NewPool[T Element](logger zerolog.Logger, factory func() T, maxCount int64) *Pool[T] {
	return &Pool[T]{
		logger: logger,
		pool: sync.Pool{
			New: func() interface{} {
				return factory()
			},
		},
		idCacheIndex: make(map[int64]T),
		currentCount: 0,
		maxCount:     maxCount,
		mutex:        sync.Mutex{},
	}
}

// Get retrieves an element from the pool
func (p *Pool[T]) Get(id int64, cacheOnly bool) T {
	p.logger.Info().Int64("id", id).Msg("Pool.Get")

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if we have the element cached by ID
	elem, exists := p.idCacheIndex[id]
	if exists {
		p.logger.Info().Msg("Pool.Get/idCacheIndex")
		delete(p.idCacheIndex, id)
		return elem
	}

	if cacheOnly {
		p.logger.Info().Msg("Pool.Get/cacheOnly")
		var zero T
		return zero
	}

	// Try to get from the general pool
	maybeElem := p.pool.Get()
	if maybeElem != nil {
		p.logger.Info().Msg("Pool.Get/pool")
		elem := maybeElem.(T)
		elem.Reset() // Reset the element state
		return elem
	}

	// Check if we've reached the maximum number of elements
	if p.maxCount <= p.currentCount {
		p.logger.Info().Int64("current_count", p.currentCount).Msg("Pool.Get/maxReached")
		var zero T
		return zero
	}

	// Create a new element
	p.currentCount = p.currentCount + 1
	p.logger.Info().Int64("current_count", p.currentCount).Msg("Pool.Get/New")

	// Use the factory function from sync.Pool
	return p.pool.New().(T)
}

// Put returns an element to the pool
func (p *Pool[T]) Put(elem T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	id := elem.GetID()
	_, exists := p.idCacheIndex[id]
	if !exists {
		p.idCacheIndex[id] = elem
		return
	}

	p.pool.Put(elem)
}
