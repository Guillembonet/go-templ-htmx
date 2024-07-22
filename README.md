# go-templ-htmx

This is a simple template of a server-side rendered Go web application using the [templ](https://github.com/a-h/templ) library for HTML templating and the [htmx](https://htmx.org/) Javascript library.

It also contains the [Alpine.js](https://alpinejs.dev/) library for some component interactivity and the [toastify JS](https://apvarun.github.io/toastify-js/) library for toast support.

## Usage

To run the application with live reload, install [air](https://github.com/air-verse/air) and use the `air` command.

Instead, you can also run the application with the `go run cmd/main.go` command and use `go generate ./...` to re-generate the templ files and tailwind css file.

## Contributing

Feel free to contribute to this project by creating a pull request. An idea for improvement is adding middlewares for authentication.

You can also contribute ideas by creating new issues.
