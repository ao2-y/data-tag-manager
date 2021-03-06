import actionCreatorFactory from 'typescript-fsa';
import { ItemTemplate } from '../../../domain/model';

const actionCreator = actionCreatorFactory();

interface IParam {
  [key: string]: string;
}
interface IError {
  message: string;
}

const search = actionCreator.async<IParam, ItemTemplate[], IError>(
  'ACTIONS_ITEMTEMPLATE_SEARCH'
);
const fetch = actionCreator.async<IParam, ItemTemplate, IError>(
  'ACTIONS_ITEMTEMPLATE_FETCH'
);
const fetchInitData = actionCreator.async<IParam, ItemTemplate, IError>(
  'ACTIONS_ITEMTEMPLATE_FETCH_INIT'
);
const register = actionCreator.async<IParam, ItemTemplate, IError>(
  'ACTIONS_ITEMTEMPLATE_REGSITER'
);

export const itemTemplateActions = {
  startSearch: search.started,
  failedSearch: search.failed,
  doneSearch: search.done,

  startFetch: fetch.started,
  failedFetch: fetch.failed,
  doneFetch: fetch.done,

  startFetchInitData: fetchInitData.started,
  failedFetchInitData: fetchInitData.failed,
  doneFetchInitData: fetchInitData.done,

  startRegister: register.started,
  failedRegister: register.failed,
  doneRegister: register.done,
};
