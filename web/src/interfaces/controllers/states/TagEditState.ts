import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Tag } from '../../../domain/model';
import { tagActions } from '../actions';

export interface TagEditState {
  tag?: Tag;
  isLoading: boolean;
  message?: string;
}

const initialState: TagEditState = {
  isLoading: false,
};

export const tagEditReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(tagActions.startFetch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(tagActions.failedFetch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(tagActions.doneFetch, (state, { result }) => {
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
