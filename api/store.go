package api

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type notFoundError struct {
	Err error
}
type conflictError struct {
	Err error
}

func (e *notFoundError) Error() string {
	return e.Err.Error()
}

func (e *conflictError) Error() string {
	return e.Err.Error()
}

// global store, implement mutex when writing to it
type Store struct {
	mutex   sync.Mutex
	advices []string
}

func NewStore() *Store {
	return &Store{
		mutex:   sync.Mutex{},
		advices: []string{},
	}
}

func (s *Store) add(val string) error {
	if s.contains(val) {
		return &conflictError{
			Err: errors.New("that advice already exists"),
		}
	}

	// add a mutex lock around this
	s.mutex.Lock()
	s.advices = append(s.advices, val)
	s.mutex.Unlock()
	return nil
}

func (s *Store) update(val string, updatedVal string) error {
	if !s.contains(val) {
		return &notFoundError{
			Err: errors.New("no item found with provided value"),
		}
	}

	// what if one already exists with that name? (The one you are updating to)
	if s.contains(updatedVal) {
		return &conflictError{
			Err: errors.New("the value you are trying to update to already exists"),
		}
	}

	// remove from the slice
	err := s.remove(val)
	if err != nil {
		return err
	}

	// add to the slice - throw away the error because we removed it just before
	_ = s.add(updatedVal)
	return nil
}

func (s *Store) remove(val string) error {
	i := s.getItemIndexValue(val)

	if !s.contains(val) {
		// we don't have the item
		return &notFoundError{
			Err: errors.New("no item found with provided value"),
		}
	}

	s.mutex.Lock()
	s.advices = append(s.advices[:i], s.advices[i+1:]...)
	s.mutex.Unlock()

	return nil
}

func (s *Store) getItemIndexValue(val string) int {
	for i, v := range s.advices {
		if v == val {
			return i
		}
	}
	return -1
}

func (s *Store) PopulateStore(amount int) error {
	// this will populate the store with random advice slips (it will retrieve the amount)
	for i := 0; i < amount; i++ {
		rAdvice, err := retrieveAdviceSlip()
		if err != nil {
			return err
		}

		if s.contains(rAdvice.Slip.Advice) {
			// we already contain the item, call ourselves again
			s.PopulateStore(amount)
		}

		s.add(rAdvice.Slip.Advice)
		time.Sleep(time.Second * 1)
	}

	return nil
}

func (s *Store) retrieveRandomAdvice() (string, error) {
	// If we have no advices to give, we need to error
	if len(s.advices) < 1 {
		return "", &notFoundError{
			Err: errors.New("there were no advices to provide"),
		}
	}

	rand.Seed(time.Now().Unix())
	return s.advices[rand.Intn(len(s.advices))], nil
}

func (s *Store) contains(containStr string) bool {
	for _, val := range s.advices {
		if val == containStr {
			return true
		}
	}

	return false
}
