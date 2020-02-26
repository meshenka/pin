package pin

import "log"

// Generator service definition
type Generator interface {
	Generate() string
}

type generatorService struct {
	logger *log.Logger
}

func (s generatorService) Generate() string {
	pin := Generate()
	s.logger.Printf("Service generate pin %s", pin)
	return pin
}

// NewGeneratorService instanciate a generator
func NewGeneratorService(logger *log.Logger) Generator {
	return generatorService{logger}
}
