# Command Line Commands

## Development

```Shell
$ npm start
```

Starts the development server and makes your application accessible at `localhost:3000`. Changes in the application code will be hot-reloaded.

## Building

```Shell
$ npm run build
```

Gets your application ready for deployment. Optimizes and minifies all files, and generates a folder called `build`. `build` includes all files you need for your application. Upload the contents of `build` to your web server to see it live!

## Browser testing

```Shell
$ npm run serve
```

This will run a server that's accessible in the entire local network and shows the version of the app that's in the `build` folder. Useful for testing on different devices!

> Note: This assumes you have done a build with `npm run build` before. The changes you make in your application won't be reflected in the application unless you run `npm run build` again.

## Unit testing

```Shell
$ npm run lint
```

Checks your JavaScript coding styles for potential errors and suspicious usage of included functions.

```Shell
$ npm run test
```

Lints JavaScript files and tests your application with the unit tests specified in the `test` folder.

```Shell
$ npm run test:watch
```

Runs unit tests everytime you change something in your JavaScript code