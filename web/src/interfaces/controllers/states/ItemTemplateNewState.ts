import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { ItemTemplate } from '../../../domain/model';
import { itemTemplateActions } from '../actions';

export interface ItemTemplateNewState {
  itemTemplate?: ItemTemplate;
  isLoading: boolean;
  message?: string;
}

const initialState: ItemTemplateNewState = {
  isLoading: false,
};

export const itemTemplateNewReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(itemTemplateActions.startFetch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemTemplateActions.failedFetch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemTemplateActions.doneFetch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, item: result });
  })
  // Register
  .case(itemTemplateActions.startRegister, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemTemplateActions.failedRegister, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemTemplateActions.doneRegister, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, item: result });
  });
