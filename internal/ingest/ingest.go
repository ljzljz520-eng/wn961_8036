package ingest

import (
	"context"
	"fmt"
	"frontend_go/internal/model"
	"strings"
)

type Parser struct{ Allowed map[string]bool }

func New() *Parser {
	return &Parser{Allowed: map[string]bool{"pdf": true, "video": true, "manual": true}}
}
func (p *Parser) Parse(ctx context.Context, id, title, category string) (model.TrainingFile, error) {
	if e := ctx.Err(); e != nil {
		return model.TrainingFile{}, e
	}
	if id == "" || title == "" {
		return model.TrainingFile{}, fmt.Errorf("missing identity")
	}
	if !p.Allowed[strings.ToLower(category)] {
		return model.TrainingFile{}, fmt.Errorf("unsupported category")
	}
	return model.TrainingFile{ID: id, Title: title, Category: category, Version: "1", Tags: []string{category}}, nil
}
func (p *Parser) Validate(f model.TrainingFile) error {
	if f.ID == "" {
		return fmt.Errorf("id")
	}
	if f.Title == "" {
		return fmt.Errorf("title")
	}
	if len(f.Tags) == 0 {
		return fmt.Errorf("tags")
	}
	return nil
}
func (p *Parser) Normalize(f model.TrainingFile) model.TrainingFile {
	f.Title = strings.TrimSpace(f.Title)
	f.Category = strings.ToLower(strings.TrimSpace(f.Category))
	return f
}
func (p *Parser) IsSupported(category string) bool { return p.Allowed[strings.ToLower(category)] }
