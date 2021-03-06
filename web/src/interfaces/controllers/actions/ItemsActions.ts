import actionCreatorFactory from 'typescript-fsa';
import { Item } from '../../../domain/model';

const actionCreator = actionCreatorFactory();

interface IParam {
  [key: string]: string;
}
interface IError {
  message: string;
}

const searchItem = actionCreator.async<IParam, Item[], IError>(
  'ACTIONS_ITEM_SEARCH'
);
const fetchItem = actionCreator.async<IParam, Item, IError>(
  'ACTIONS_ITEM_FETCH'
);
const fetchInitItem = actionCreator.async<IParam, Item, IError>(
  'ACTIONS_ITEM_FETCH_INIT'
);
const registerItem = actionCreator.async<IParam, Item, IError>(
  'ACTIONS_ITEM_REGSITER'
);

export const itemActions = {
  startSearch: searchItem.started,
  failedSearch: searchItem.failed,
  doneSearch: searchItem.done,

  startFetch: fetchItem.started,
  failedFetch: fetchItem.failed,
  doneFetch: fetchItem.done,

  startFetchInitItem: fetchInitItem.started,
  failedFetchInitItem: fetchInitItem.failed,
  doneFetchInitItem: fetchInitItem.done,

  startRegister: registerItem.started,
  failedRegister: registerItem.failed,
  doneRegister: registerItem.done,
};
