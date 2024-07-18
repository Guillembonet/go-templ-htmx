package bunetz

//go:generate go run github.com/a-h/templ/cmd/templ@v0.2.731 generate

//go:generate npx tailwindcss -c ./tailwind.config.js -i ./views/assets/css/input.css -o ./views/assets/css/tailwind.css --minify
