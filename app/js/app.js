/**
 *
 * app.js
 *
 * This is the entry file for the application, mostly just setup and boilerplate
 * code. Routes are configured at the end of this file!
 *
 */

// Import all the third party stuff
import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import FontFaceObserver from 'fontfaceobserver';
import createHistory from 'history/lib/createBrowserHistory';

// Observer loading of Open Sans (to remove open sans, remove the <link> tag in the index.html file and this observer)
const openSansObserver = new FontFaceObserver('Open Sans', {});

// When Open Sans is loaded, add the js-open-sans-loaded class to the body
openSansObserver.check().then(() => {
  document.body.classList.add('js-open-sans-loaded');
}, () => {
  document.body.classList.remove('js-open-sans-loaded');
});

// Import the pages
import HomePage from './components/pages/HomePage.js';
import ReadmePage from './components/pages/ReadmePage.js';
import NotFoundPage from './components/pages/NotFound.js';
import App from './components/App.js';

// Import the CSS file, which HtmlWebpackPlugin transfers to the build folder
import '../css/index.styl';

// Create the store with the redux-thunk middleware, which allows us
// to do asynchronous things in the actions
import rootReducer from './reducers/rootReducer';
const createStoreWithMiddleware = applyMiddleware(thunk)(createStore);
const store = createStoreWithMiddleware(rootReducer);

// Make reducers hot reloadable, see http://stackoverflow.com/questions/34243684/make-redux-reducers-and-other-non-components-hot-loadable
if (module.hot) {
  module.hot.accept('./reducers/rootReducer', () => {
    const nextRootReducer = require('./reducers/rootReducer').default;
    store.replaceReducer(nextRootReducer);
  });
}

// Mostly boilerplate, except for the Routes. These are the pages you can go to,
// which are all wrapped in the App component, which contains the navigation etc
ReactDOM.render(
  <Provider store={store}>
    <Router history={createHistory()}>
      <Route component={App}>
        <Route path="/" component={HomePage} />
        <Route path="/readme" component={ReadmePage} />
        <Route path="*" component={NotFoundPage} />
      </Route>
    </Router>
  </Provider>,
  document.getElementById('ReactApp')
);

// TODO: Websocket integration
// (function() {
//     console.log('Initializing websockets...');
//     var conn = new WebSocket("ws://localhost:8080/ws?lastMod=143918dd9ce16851");
//     conn.onclose = function(evt) {
//         console.log('Connection closed');
//     }
//     conn.onmessage = function(evt) {
//         console.log('file updated', evt.data);
//     }
// })();
