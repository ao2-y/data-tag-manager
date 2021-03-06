import { createStore, combineReducers } from 'redux';
import { layoutReducer, LayoutState } from '../states/LayoutState';
import {
  itemTemplatesReducer,
  ItemTemplatesState,
  itemTemplateEditReducer,
  ItemTemplateEditState,
  itemTemplateNewReducer,
  ItemTemplateNewState,
  itemsReducer,
  ItemsState,
  itemEditReducer,
  ItemEditState,
  itemNewReducer,
  ItemNewState,
  tagsReducer,
  TagsState,
  tagEditReducer,
  TagEditState,
  tagNewReducer,
  TagNewState,
  metaKeysReducer,
  MetaKeysState,
  metaKeyEditReducer,
  MetaKeyEditState,
  metaKeyNewReducer,
  MetaKeyNewState,
} from '../states';

export type AppState = {
  items: ItemsState;
  itemEdit: ItemEditState;
  itemNew: ItemNewState;

  itemTemplates: ItemTemplatesState;
  itemTemplateEdit: ItemTemplateEditState;
  itemTemplateNew: ItemTemplateNewState;

  tags: TagsState;
  tagEdit: TagEditState;
  tagNew: TagNewState;

  metaKeys: MetaKeysState;
  metaKeyEdit: MetaKeyEditState;
  metaKeyNew: MetaKeyNewState;

  layout: LayoutState;
};

const store = createStore(
  combineReducers<AppState>({
    items: itemsReducer,
    itemEdit: itemEditReducer,
    itemNew: itemNewReducer,

    itemTemplates: itemTemplatesReducer,
    itemTemplateNew: itemTemplateNewReducer,
    itemTemplateEdit: itemTemplateEditReducer,

    tags: tagsReducer,
    tagNew: tagNewReducer,
    tagEdit: tagEditReducer,

    metaKeys: metaKeysReducer,
    metaKeyNew: metaKeyNewReducer,
    metaKeyEdit: metaKeyEditReducer,

    layout: layoutReducer,
  })
);

export default store;
