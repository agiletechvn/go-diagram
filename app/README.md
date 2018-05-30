# React.js Boilerplate

[![Dependency Status][dep-status-img]][dep-status-link] [![devDependency Status][dev-dep-status-img]][dev-dep-status-link] [![Build Status][ci-img]][ci] [![Join the chat at https://gitter.im/mxstbr/react-boilerplate](https://badges.gitter.im/mxstbr/react-boilerplate.svg)](https://gitter.im/mxstbr/react-boilerplate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Quick setup for new performance orientated, offline–first React.js applications featuring Redux, hot–reloading, PostCSS, react-router, ServiceWorker, AppCache, FontFaceObserver and Mocha.

Made with :heart: by [Max Stoiber](https://twitter.com/mxstbr)!

[dep-status-img]: https://david-dm.org/mxstbr/react-boilerplate.svg
[dep-status-link]: https://david-dm.org/mxstbr/react-boilerplate
[dev-dep-status-img]: https://david-dm.org/mxstbr/react-boilerplate/dev-status.svg
[dev-dep-status-link]: https://david-dm.org/mxstbr/react-boilerplate#info=devDependencies
[ci-img]: https://travis-ci.org/mxstbr/react-boilerplate.svg?branch=master
[ci]: https://travis-ci.org/mxstbr/react-boilerplate

-----

## Features

- Using [**react-transform-hmr**](https://github.com/gaearon/react-transform-hmr), your changes in the CSS and JS get reflected in the app instantly without refreshing the page. That means that the **current application state persists** even when you change something in the underlying code! For a very good explanation and demo, watch Dan Abramov himself [talking about it at react-europe](https://www.youtube.com/watch?v=xsSnOQynTHs).

- [**Redux**](https://github.com/rackt/redux) is a much better implementation of a flux–like, unidirectional data flow. Redux makes actions composable, reduces the boilerplate code and makes hot–reloading possible in the first place. For a good overview of redux, check out the talk linked above or the [official documentation](https://gaearon.github.io/redux/)!

- [**Babel**](http://babeljs.io/) is a modular JavaScript transpiler that helps to use next generation JavaScript and more, like transformation for JSX, hot loading, error catching etc. Babel has a [solid ecosystem of offical preset and plugins](https://github.com/babel/babel/tree/master/packages).

- [**PostCSS**](https://github.com/postcss/postcss) is like Sass, but modular and capable of much more. PostCSS is, in essence, just a wrapper for plugins which exposes an easy to use, but very powerful API. While it is possible to [replicate Sass features](https://github.com/jonathantneal/precss) with PostCSS, PostCSS has an [ecosystem of amazing plugins](http://postcss.parts) with functionalities Sass cannot even dream about having. See [this talk](http://mxs.is/pctalk) for a short introduction to PostCSS.

- **Unit tests** should be an important part of every web application developers toolchain. [Mocha](https://github.com/mochajs/mocha) checks your application is working exactly how it should without you lifting a single finger. Congratulations, you just won a First Class ticket to world domination, fasten your seat belt please!

- [**react-router**](https://github.com/rackt/react-router) is used for routing in this boilerplate. react-router makes routing easy to do and takes care of most of the work. (except you have to define your routes yourself) You can find the documentation right  [here](https://github.com/rackt/react-router/blob/master/doc/00%20Guides/0%20Overview.md).

- [**ServiceWorker**](http://www.html5rocks.com/en/tutorials/service-worker/introduction/) and [**AppCache**](http://www.html5rocks.com/en/tutorials/appcache/beginner/) make it possible to use your application offline. As soon as the website has been opened once, it is cached and available without a network connection. See [this talk](http://mxs.is/swtalk) for an explanation of the ServiceWorker used in this boilerplate. [**`manifest.json`**](https://developer.chrome.com/multidevice/android/installtohomescreen) is specifically for Chrome on Android. Users can add the website to the homescreen and use it like a native app!

## Getting started

1. Clone this repo using `git clone https://github.com/mxstbr/react-boilerplate.git`.

2. Delete the existing git repository by running `rm -rf .git`.

3. Initialize a new git repository with `git init`, `git add .` and `git commit -m "Initial commit"`.

4. Run `npm install` to install the dependencies.

5. Run `npm start` to start the local web server.

6. Go to `http://localhost:3000` and you should see the app running!

## Building & Deploying

1. Run `npm run build`, which will compile all the necessary files in a `build` folder.

2. Upload the contents of the `build` folder to your web server.

### Server Configuration

#### Apache

This boilerplate includes a `.htaccess` file that does two things:

1. Redirect all traffic to HTTPS because ServiceWorker only works for encrypted traffic

2. Rewrite all pages (e.g. yourdomain.com/subpage) to the `index.html` to let `react-router` take care of presenting the correct page

## CSS

The CSS modules found in the `css` subfolders all get imported into one big file (`main.css`), which gets transpiled with PostCSS.

See the [`css` folder README](css/README.md) for more information about the PostCSS plugins used and the CSS structure.

## JS

All files that are `import`ed/`require`d somewhere get compiled into one big file at build time. (`build/bundle.js`) Webpack automatically optimizes your JavaScript with `UglifyJS`, so you don't have to worry about that.

See the [`js` folder README](js/README.md) for more information about the JS structure.

## Testing

### Browser testing

With `npm run serve` you can start a server that's accessible in the entire local network and shows the version of the app that's in the `build` folder. Useful for testing on different devices!

### Unit testing

Unit tests live in the `test` directory, and are run with `npm test`. There are two files in there already, testing that the current AppActions and Reducers are working correctly. Mocha is used as a unit testing framework, check out the official [documentation](http://mochajs.org)!

## Files in the root folder

For a full explanation of all files in the root folder, see [`docs/FILES.md`](docs/FILES.md).

## Opinionated features

### Web Fonts

If you simply use web fonts in your project, the page will stay blank until these fonts are downloaded. That means a lot of waiting time in which users could already read the content.

[FontFaceObserver](https://github.com/bramstein/fontfaceobserver) adds a `js-<font-name>-loaded` class to the `body` when the fonts have loaded. You should specify an initial `font-family` with save fonts on the `body`, and a `.js-<font-name>-loaded` `font-family` which includes your web font. See [app.js](js/app.js#L17-L25) and [base.css](css/base/_base.css#L26-L32).

#### Adding a new font

1. Either add the `@font-face` declaration to `base/_fonts.css` or add a `<link>` tag to the [`index.html`](index.html). (Don't forget to remove the `<link>` for Open Sans from the [`index.html`](index.html))

2. In `base/_base.css`, specify your initial `font-family` in the `body` tag with only save fonts. In the `body.js-<font-name>-loaded` tag, specify your `font-family` stack with your web font.

3. In `js/app.js` add a `<font-name>Observer` for your font.

#### Removing performant web font loading

**Careful** about removing this, as perceived performance might be highly impacted.

To remove `FontFaceObserver`, don't import it in [`app.js`](js/app.js) and remove it from the [`package.json`](package.json).

### Offline access

Using a `ServiceWorker` and the `Application Cache`, your application is cached for offline usage.

#### Cache a new file

To cache a file, add it to the `urlsToCache` variable in the [`serviceworker.js`](serviceworker.js) file.

#### Removing offline access

**Careful** about removing this, as there is no real downside to having your application available when the users network connection isn't perfect.

To remove offline capability, delete [`serviceworker.js`](serviceworker.js), remove the import in [`app.js`](js/app.js), remove `AppCachePlugin` in [`makewebpackconfig.js`](makewebpackconfig.js) and remove the `manifest` attribute of the `<html>` tag in [`index.html`](index.html).

### Add To Homescreen

On Chrome for Android (soon hopefully more browsers), users can add a webpage to the homescreen. Combined with offline caching, this means your web app can be used exactly like a native application.

The name and icon to be displayed are set in the `manifest.json` file. Change them to your project name and icon, and try it!

#### Removing add to homescreen functionality

Delete [`manifest.json`](manifest.json) and remove the `<link rel="manifest" href="manifest.json">` tag from the [`index.html`](index.html).

## Gotchas

These are some things to be aware of when using this boilerplate.

### Images in the HTML file(s)

Adding images to the HTML is a bit of a pain right now as webpack only goes through the JavaScript file. Add the image to your HTML file how you always would:

```HTML
<!-- Normal Image -->
<img src="img/yourimg.png" />
<!-- Meta tags -->
<meta property="og:image" content="img/yourimg.png" />
<!-- ... -->
```

If you simply do this, webpack will not transfer the images to the build folder. To get webpack to transfer them, you have to import them with the file loader in your JavaScript somewhere, e.g.:

```JavaScript
import 'file?name=[name].[ext]!../img/yourimg.png';
```

Then webpack will correctly transfer the image to the build folder.

## How to contribute

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for infos on how you can help!

## License

This project is licensed under the MIT license, Copyright (c) 2016 Maximilian Stoiber. For more information see `LICENSE.md`.
