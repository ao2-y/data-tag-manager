import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { MetaKey } from '../../../domain/model';
import { metaKeyActions } from '../actions';

export interface MetaKeysState {
  metaKeys: MetaKey[];
  isLoading: boolean;
  message?: string;
}

const initialState: MetaKeysState = {
  metaKeys: [],
  isLoading: false,
};

export const metaKeysReducer = reducerWithInitialState(initialState)
  .case(metaKeyActions.startSearch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(metaKeyActions.failedSearch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(metaKeyActions.doneSearch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, metaKeys: result });
  });
