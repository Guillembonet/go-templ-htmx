//go:build tools

package tools

import (
	_ "github.com/a-h/templ/cmd/templ"
)

//go:generate go run github.com/a-h/templ/cmd/templ generate

//go:generate npx tailwindcss -c ./tailwind.config.js -i ./views/assets/css/input.css -o ./views/assets/css/tailwind.css --minify
