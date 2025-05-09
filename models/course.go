package models

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Course represents a learning course in the platform
type Course struct {
	Id          int       `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Content     string    `json:"content" bson:"content"`
	Category    string    `json:"category" bson:"category"`
	Lessons     int       `json:"lessons" bson:"lessons"`
	Sections    int       `json:"sections" bson:"sections"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	CreatedBy   int       `json:"created_by" bson:"created_by"`
	Image       string    `json:"image,omitempty" bson:"image,omitempty"`
	Thumbnail   string    `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	Videos      []Video   `json:"videos,omitempty" bson:"videos,omitempty"`
	Level       string    `json:"level,omitempty" bson:"level,omitempty"`
	Duration    string    `json:"duration,omitempty" bson:"duration,omitempty"`
	Instructor  string    `json:"instructor,omitempty" bson:"instructor,omitempty"`
}

// Video represents a lesson video in a course
type Video struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Url         string `json:"url" bson:"url"`
	Duration    string `json:"duration" bson:"duration"`
}

func (t *Course) Valid() bool {
	// Basic validation: Title and at least one of Description or Content must be present
	if t.Title == "" {
		return false
	}
	return true
}

func (v *Validator) CourseFromBody() *Course {
	b, err := ioutil.ReadAll(v.r.Body)
	if err != nil {
		v.Error("body", err.Error())
		return nil
	}

	ret := &Course{}
	err = json.Unmarshal(b, ret)
	if err != nil {
		v.Error("body", err.Error())
		return nil
	}

	if !ret.Valid() {
		v.Error("body", "Invalid Course")
		return nil
	}

	return ret
}