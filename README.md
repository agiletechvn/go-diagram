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
# install client dependencies
yarn install-client
# build client
yarn build-client

# start server in hot-reload mode
yarn start <directory name>
# run client in dev mode
yarn client
```

## Tech Stack

*   Golang for the server (uses goparser, gorilla websockets)
*   React + Redux frontend
    *   Webpack
    *   Babel
    *   Stylus
