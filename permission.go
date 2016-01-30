package gopermission

import (
	"net/http"
)

// Interface to check permission. Implement it to verify some rule
type Checker interface {
	HasPermission(request *http.Request) bool
}

// Main permission object. Create it and use in middleware
type Permission struct {
	checkers []Checker
}

func (p *Permission) AddChecker(checker Checker) {
	p.checkers = append(p.checkers, checker)
}

func (p *Permission) IsPermitted(request *http.Request) bool {
	for _, v := range p.checkers {
		if !v.HasPermission(request) {
			return false
		}
	}
	return true
}

func New(checkers ...Checker) Permission {
	return Permission{checkers}
}
