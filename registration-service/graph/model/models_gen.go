// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Mutation struct {
}

type Query struct {
}

type Registration struct {
	ID         string   `json:"id"`
	StudentID  string   `json:"studentId"`
	CourseID   string   `json:"courseId"`
	Term       string   `json:"term"`
	Grade      *float64 `json:"grade,omitempty"`
	Status     string   `json:"status"`
	EnrolledAt string   `json:"enrolledAt"`
	UpdatedAt  string   `json:"updatedAt"`
}

type RegistrationInput struct {
	StudentID string   `json:"studentId"`
	CourseID  string   `json:"courseId"`
	Term      string   `json:"term"`
	Grade     *float64 `json:"grade,omitempty"`
	Status    string   `json:"status"`
}
