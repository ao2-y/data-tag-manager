import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Item } from '../../../domain/model';
import { itemActions } from '../actions';

export interface ItemNewState {
  item?: Item;
  template?: number;
  isLoading: boolean;
  message?: string;
}

const initialState: ItemNewState = {
  isLoading: false,
};

export const itemNewReducer = reducerWithInitialState(initialState)
.case(itemActions.startFetchInitItem, (state) => {
  return Object.assign({}, state, { isLoading: true });
})
.case(itemActions.failedFetchInitItem, (state, { error }) => {
  return Object.assign({}, state, {
    isLoading: false,
    message: error.message,
  });
})
.case(itemActions.doneFetchInitItem, (state, { result }) => {
  return Object.assign({}, state, { isLoading: false, item: result });
})
;
