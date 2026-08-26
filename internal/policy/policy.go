package policy

import (
	"fmt"
	"frontend_go/internal/model"
)

type Rule struct {
	Role       string
	CanArchive bool
	CanReview  bool
}

func Default(role string) Rule {
	r := Rule{Role: role}
	if role == "admin" {
		r.CanArchive = true
		r.CanReview = true
	}
	if role == "reviewer" {
		r.CanReview = true
	}
	return r
}
func AuthorizeArchive(r Rule) error {
	if !r.CanArchive {
		return fmt.Errorf("archive denied")
	}
	return nil
}
func AuthorizeReview(r Rule) error {
	if !r.CanReview {
		return fmt.Errorf("review denied")
	}
	return nil
}
func CanView(r model.Record, student string) bool {
	return r.StudentID == student || student == "admin"
}
func ValidRole(role string) bool { return role == "admin" || role == "reviewer" || role == "student" }
