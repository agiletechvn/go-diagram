# go-diagram

A UML diagram editor for Golang projects

## Setup

Make sure your `PATH` includes `:$GOPATH/bin`

```sh
glide install

cd app
yarn
yarn build
# dev mode
yarn start

cd ..
go run server.go <directory name>
```

## Tech Stack

*   Golang for the server (uses goparser, gorilla websockets)
*   React + Redux frontend
    *   Webpack
    *   Babel
    *   Stylus
