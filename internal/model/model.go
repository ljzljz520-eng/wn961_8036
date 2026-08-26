package model

import "time"

type Record struct {
	ID, FileID, StudentID, Action, Notes string
	CreatedAt                            time.Time
	Archived                             bool
}
type Profile struct {
	ID, Name, Department, Level string
	Active                      bool
	UpdatedAt                   time.Time
}
type Event struct {
	ID, RecordID, Kind, Actor, Payload string
	At                                 time.Time
}
type Audit struct {
	ID, Entity, EntityID, Operation, Actor string
	At                                     time.Time
	Detail                                 string
}
type TrainingFile struct {
	ID, Title, Category, Version string
	Published                    bool
	Tags                         []string
}

func NewRecord(id, file, student, action string) Record {
	return Record{ID: id, FileID: file, StudentID: student, Action: action, CreatedAt: time.Now().UTC()}
}
func NewProfile(id, name string) Profile {
	return Profile{ID: id, Name: name, Active: true, UpdatedAt: time.Now().UTC()}
}
func NewEvent(id, record, kind, actor string) Event {
	return Event{ID: id, RecordID: record, Kind: kind, Actor: actor, At: time.Now().UTC()}
}
func NewAudit(id, entity, entityID, op, actor string) Audit {
	return Audit{ID: id, Entity: entity, EntityID: entityID, Operation: op, Actor: actor, At: time.Now().UTC()}
}
func (r Record) Valid() bool {
	return r.ID != "" && r.FileID != "" && r.StudentID != "" && r.Action != ""
}
func (p Profile) Valid() bool { return p.ID != "" && p.Name != "" }
func (e Event) Valid() bool   { return e.ID != "" && e.RecordID != "" && e.Kind != "" }
func (a Audit) Valid() bool   { return a.ID != "" && a.Entity != "" && a.EntityID != "" }
func (f TrainingFile) SearchText() string {
	return f.ID + " " + f.Title + " " + f.Category + " " + f.Version
}
func (f TrainingFile) HasTag(tag string) bool {
	for _, t := range f.Tags {
		if t == tag {
			return true
		}
	}
	return false
}
