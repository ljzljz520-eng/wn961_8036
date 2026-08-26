package catalog

import (
	"frontend_go/internal/model"
	"sort"
	"strings"
)

type Catalog struct{ items map[string]model.TrainingFile }

func New() *Catalog { return &Catalog{items: map[string]model.TrainingFile{}} }
func (c *Catalog) Add(f model.TrainingFile) bool {
	if f.ID == "" || f.Title == "" {
		return false
	}
	c.items[f.ID] = f
	return true
}
func (c *Catalog) Remove(id string) bool {
	if _, ok := c.items[id]; !ok {
		return false
	}
	delete(c.items, id)
	return true
}
func (c *Catalog) Get(id string) (model.TrainingFile, bool) { f, ok := c.items[id]; return f, ok }
func (c *Catalog) Published() []model.TrainingFile {
	out := []model.TrainingFile{}
	for _, f := range c.items {
		if f.Published {
			out = append(out, f)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func (c *Catalog) Search(term string) []model.TrainingFile {
	out := []model.TrainingFile{}
	term = strings.ToLower(term)
	for _, f := range c.items {
		if strings.Contains(strings.ToLower(f.SearchText()), term) {
			out = append(out, f)
		}
	}
	return out
}
func (c *Catalog) Tag(id, tag string) bool {
	f, ok := c.items[id]
	if !ok {
		return false
	}
	if !f.HasTag(tag) {
		f.Tags = append(f.Tags, tag)
	}
	c.items[id] = f
	return true
}
func (c *Catalog) Publish(id string) bool {
	f, ok := c.items[id]
	if !ok {
		return false
	}
	f.Published = true
	c.items[id] = f
	return true
}
