import 'reflect-metadata';
import React from 'react';
import ReactDOM from 'react-dom';
import { container } from 'tsyringe';
import './scss/index.scss';
import App from './App';
import * as serviceWorker from './serviceWorker';

import { icons } from './assets/icons';

import { Provider } from 'react-redux';
import store from './interfaces/controllers/stores/store';

import { ItemTemplateDataStore } from './infrastructure/datastore/ItemTemplateDataStore';
import { MetaKeyDataStore } from './infrastructure/datastore/MetaKeyDataStore';
import { TagDataStore } from './infrastructure/datastore/TagDataStore';
import { ItemDataStore } from './infrastructure/datastore/ItemDataStore';
container.register('ItemTemplateRepository', {
  useClass: ItemTemplateDataStore,
});
container.register('MetaKeyRepository', {
  useClass: MetaKeyDataStore,
});
container.register('TagRepository', {
  useClass: TagDataStore,
});
container.register('ItemRepository', {
  useClass: ItemDataStore,
});

// eslint-disable-next-line @typescript-eslint/no-explicit-any
(React as any).icons = icons;
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
