import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Item } from '../../../domain/model';
import { itemActions } from '../actions';

export interface ItemEditState {
  item?: Item;
  isLoading: boolean;
  message?: string;
}

const initialState: ItemEditState = {
  isLoading: false,
};

export const itemEditReducer = reducerWithInitialState(initialState)
  // Fetch
  .case(itemActions.startFetch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemActions.failedFetch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemActions.doneFetch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, item: result });
  })
  // Register
  .case(itemActions.startRegister, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemActions.failedRegister, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemActions.doneRegister, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, item: result });
  });
