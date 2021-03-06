import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { MetaKey } from '../../../domain/model';
import { metaKeyActions } from '../actions';

export interface MetaKeyNewState {
  metaKey?: MetaKey;
  isLoading: boolean;
  message?: string;
}

const initialState: MetaKeyNewState = {
  isLoading: false,
};

export const metaKeyNewReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(metaKeyActions.startFetchInitData, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(metaKeyActions.failedFetchInitData, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(metaKeyActions.doneFetchInitData, (state, { result }) => {
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
