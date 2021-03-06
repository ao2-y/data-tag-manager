import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Tag } from '../../../domain/model';
import { tagActions } from '../actions';

export interface TagNewState {
  tag?: Tag;
  isLoading: boolean;
  message?: string;
}

const initialState: TagNewState = {
  isLoading: false,
};

export const tagNewReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(tagActions.startFetchInitData, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(tagActions.failedFetchInitData, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(tagActions.doneFetchInitData, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, tag: result });
  })
  // Register
  .case(tagActions.startRegister, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(tagActions.failedRegister, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(tagActions.doneRegister, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, tag: result });
  });
