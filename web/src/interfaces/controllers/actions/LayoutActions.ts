import actionCreatorFactory from 'typescript-fsa';

const actionCreator = actionCreatorFactory();

export const layoutActions = {
  showSidebar: actionCreator<'responsive' | boolean>('ACTIONS_SHOW_SIDEBAR'),
};
