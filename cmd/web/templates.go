package main

import "snippetbox.saulosilva.dev/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
