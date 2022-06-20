package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandlder(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ (tpl Template) http.HandlerFunc {
	questions := []struct{
		Question string
		Answer template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer: "Yes.",
		},
		{
			Question: "What are your support hours?",
			Answer: "24/7.",
		},
		{
			Question: "How do I contact support?",
			Answer: `Via email - <a href="mailto:email@email.com">support@email.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer: "Our entire team is remote.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}