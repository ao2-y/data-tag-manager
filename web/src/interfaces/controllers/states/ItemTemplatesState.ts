import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { ItemTemplate } from '../../../domain/model';
import { itemTemplateActions } from '../actions';

export interface ItemTemplatesState {
  items: ItemTemplate[];
  isLoading: boolean;
  message?: string;
}

const initialState: ItemTemplatesState = {
  items: [],
  isLoading: false,
};

export const itemTemplatesReducer = reducerWithInitialState(initialState)
  .case(itemTemplateActions.startSearch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemTemplateActions.failedSearch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemTemplateActions.doneSearch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, items: result });
  });
