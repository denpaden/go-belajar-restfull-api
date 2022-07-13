package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

// provider untuk membuat pointer SimpleRepository
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// SimpleService butuh dependecy pointer SimpleRepository
type SimpleService struct {
	*SimpleRepository
}

// provider untuk membuat pointer SimpleService
func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("failed created service")
	} else {
		return &SimpleService{SimpleRepository: simpleRepository}, nil
	}

}
