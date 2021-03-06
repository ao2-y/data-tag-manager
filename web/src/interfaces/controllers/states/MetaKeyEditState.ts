import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { MetaKey } from '../../../domain/model';
import { metaKeyActions } from '../actions';

export interface MetaKeyEditState {
  metaKey?: MetaKey;
  isLoading: boolean;
  message?: string;
}

const initialState: MetaKeyEditState = {
  isLoading: false,
};

export const metaKeyEditReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(metaKeyActions.startFetch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(metaKeyActions.failedFetch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(metaKeyActions.doneFetch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, metaKey: result });
  })
  // Register
  .case(metaKeyActions.startRegister, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(metaKeyActions.failedRegister, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(metaKeyActions.doneRegister, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, metaKey: result });
  });
