# go-diagram

A UML diagram editor for Golang projects

## **Install the requirements**

**Install nodemon**

```sh
yarn global add nodemon
```

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
yarn start <directory name>
# run client in dev mode
yarn client
# build client
yarn build-client
```

## Tech Stack

*   Golang for the server (uses goparser, gorilla websockets)
*   React + Redux frontend
    *   Webpack
    *   Babel
    *   Stylus
