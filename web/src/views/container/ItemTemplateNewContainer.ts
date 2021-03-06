import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemTemplateActions } from '../../interfaces/controllers/actions';
import { ItemTemplate } from '../../domain/model';
import { ItemTemplateUsecase } from '../../usecase/ItemTemplateUsecase';
import screen from '../screens/ItemTemplateNewScreen';

export interface ItemTemplateNewActions {
  loadInitData: () => void;
  register: (v: ItemTemplate) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(ItemTemplateUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    loadInitData() {
      dispatch(itemTemplateActions.startFetchInitData({}));
      usecase
        .fetchInitData()
        .then((result) => {
          dispatch(
            itemTemplateActions.doneFetchInitData({ params: {}, result })
          );
        })
        .catch(() => {
          dispatch(
            itemTemplateActions.failedFetchInitData({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
    register(v: ItemTemplate) {
      dispatch(itemTemplateActions.startRegister({}));
      const register = v.id === 0 ? usecase.create : usecase.update;
      register(v)
        .then((result) => {
          if (result) {
            dispatch(itemTemplateActions.doneRegister({ params: {}, result }));
          } else {
            throw new Error();
          }
        })
        .catch(() => {
          dispatch(
            itemTemplateActions.failedRegister({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.itemTemplateNew);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
