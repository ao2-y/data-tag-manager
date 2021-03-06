import actionCreatorFactory from 'typescript-fsa';
import { Tag } from '../../../domain/model';

const actionCreator = actionCreatorFactory();

interface IParam {
  [key: string]: string;
}
interface IError {
  message: string;
}

const search = actionCreator.async<IParam, Tag[], IError>('ACTIONS_TAG_SEARCH');
const fetch = actionCreator.async<IParam, Tag, IError>('ACTIONS_TAG_FETCH');
const fetchInitData = actionCreator.async<IParam, Tag, IError>(
  'ACTIONS_TAG_FETCH_INIT'
);
const register = actionCreator.async<IParam, Tag, IError>(
  'ACTIONS_TAG_REGSITER'
);

export const tagActions = {
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
