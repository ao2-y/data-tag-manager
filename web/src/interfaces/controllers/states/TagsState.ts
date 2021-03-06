import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Tag } from '../../../domain/model';
import { tagActions } from '../actions';

export interface TagsState {
  tags: Tag[];
  isLoading: boolean;
  message?: string;
}

const initialState: TagsState = {
  tags: [],
  isLoading: false,
};

export const tagsReducer = reducerWithInitialState(initialState)
  .case(tagActions.startSearch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(tagActions.failedSearch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(tagActions.doneSearch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, tags: result });
  });
