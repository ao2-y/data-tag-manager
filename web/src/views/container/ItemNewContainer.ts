import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemActions } from '../../interfaces/controllers/actions';
import { ItemUsecase } from '../../usecase/ItemUsecase';
import { Item } from '../../domain/model';
import screen from '../screens/ItemNewScreen';

export interface ItemNewActions {
  loadInitItem: (template?: number) => void;
  register: (v: Item) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(ItemUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    loadInitItem(template?: number) {
      dispatch(itemActions.startFetchInitItem({}));
      usecase
        .fetchInitData(template)
        .then((result) => {
          dispatch(itemActions.doneFetchInitItem({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            itemActions.failedFetchInitItem({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
    register(v: Item) {
      dispatch(itemActions.startRegister({}));
      usecase
        .create(v)
        .then((result) => {
          if (result) {
            dispatch(itemActions.doneRegister({ params: {}, result }));
          } else {
            throw new Error();
          }
        })
        .catch(() => {
          dispatch(
            itemActions.failedRegister({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.itemNew);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
