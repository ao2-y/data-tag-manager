import { reducerWithInitialState } from 'typescript-fsa-reducers';
import { layoutActions } from '../actions/LayoutActions';

export interface LayoutState {
  sidebarShow: 'responsive' | boolean;
}

const initialState: LayoutState = {
  sidebarShow: 'responsive',
};

export const layoutReducer = reducerWithInitialState(initialState)
  .case(layoutActions.showSidebar, (state, sidebarShow) => {
    return Object.assign({}, state, { sidebarShow });
  });
