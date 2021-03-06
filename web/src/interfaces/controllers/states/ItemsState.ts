import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { Item, ItemTemplate } from '../../../domain/model';
import { itemActions, itemTemplateActions } from '../actions';

export interface ItemsState {
  items: Item[];
  isItemsLoading: boolean;
  itemsLoadMessage?: string;

  itemTemplates: ItemTemplate[];
  isItemTemplatesLoading: boolean;
  itemTemplatesLoadMessage?: string;

  isShowModal: boolean; 
}

const initialState: ItemsState = {
  items: [],
  isItemsLoading: false,

  itemTemplates: [],
  isItemTemplatesLoading: false,

  isShowModal: false,
};

export const itemsReducer = reducerWithInitialState(initialState)
  .case(itemActions.startSearch, (state) => {
    return Object.assign({}, state, { isLoading: true });
  })
  .case(itemActions.failedSearch, (state, { error }) => {
    return Object.assign({}, state, {
      isLoading: false,
      message: error.message,
    });
  })
  .case(itemActions.doneSearch, (state, { result }) => {
    return Object.assign({}, state, { isLoading: false, items: result });
  })
  .case(itemTemplateActions.startSearch, (state) => {
    return Object.assign({}, state, { isItemTemplatesLoading: true });
  })
  .case(itemTemplateActions.failedSearch, (state, { error }) => {
    return Object.assign({}, state, {
      isItemTemplatesLoading: false,
      message: error.message,
    });
  })
  .case(itemTemplateActions.doneSearch, (state, { result }) => {
    return Object.assign({}, state, { isItemTemplatesLoading: false, itemTemplates: result });
  })
  ;
