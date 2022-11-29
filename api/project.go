package api

import (
	"github.com/SoufianeRep/tscit/db"
	"github.com/jackc/pgtype"
)

type createProjectRequest struct {
	ID uint `json:"id" binding:"required,min=1"`
}

type projectResponse struct {
	ID         uint         `json:"id"`
	Name       string       `json:"name"`
	Language   string       `json:"langauge"`
	Length     uint         `json:"length"`
	Transcript pgtype.JSONB `json:"transcript"`
}

func ProjectResponse(project db.Project) projectResponse {
	return projectResponse{
		ID:         project.ID,
		Name:       project.Name,
		Language:   project.Language,
		Length:     project.Length,
		Transcript: project.Transcript,
	}
}
