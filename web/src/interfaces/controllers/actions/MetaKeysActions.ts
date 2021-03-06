import actionCreatorFactory from 'typescript-fsa';
import { MetaKey } from '../../../domain/model';

const actionCreator = actionCreatorFactory();

interface IParam {
  [key: string]: string;
}
interface IError {
  message: string;
}

const search = actionCreator.async<IParam, MetaKey[], IError>(
  'ACTIONS_METAKEY_SEARCH'
);
const fetch = actionCreator.async<IParam, MetaKey, IError>(
  'ACTIONS_METAKEY_FETCH'
);
const fetchInitData = actionCreator.async<IParam, MetaKey, IError>(
  'ACTIONS_METAKEY_FETCH_INIT'
);
const register = actionCreator.async<IParam, MetaKey, IError>(
  'ACTIONS_METAKEY_REGSITER'
);

export const metaKeyActions = {
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
