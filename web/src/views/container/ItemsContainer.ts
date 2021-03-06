import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemActions, itemTemplateActions } from '../../interfaces/controllers/actions';
import { ItemUsecase } from '../../usecase/ItemUsecase';
import { ItemTemplateUsecase } from '../../usecase/ItemTemplateUsecase';
import screen from '../screens/ItemsScreen';

export interface ItemsActions {
  searchItems: () => void;
  searchItemTemplates: () => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const itemUsecase = container.resolve(ItemUsecase);
const itemTemplateUsecase = container.resolve(ItemTemplateUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    searchItems() {
      dispatch(itemActions.startSearch({}));
      itemUsecase
        .find()
        .then((result) => {
          dispatch(itemActions.doneSearch({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            itemActions.failedSearch({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
    searchItemTemplates() {
      dispatch(itemTemplateActions.startSearch({}));
      itemTemplateUsecase
        .find()
        .then((result) => {
          dispatch(itemTemplateActions.doneSearch({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            itemTemplateActions.failedSearch({
              params: {},
              error: { message: '' },
            })
          );
        });
    }
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.items);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
