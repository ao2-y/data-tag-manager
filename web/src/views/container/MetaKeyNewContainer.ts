import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { metaKeyActions } from '../../interfaces/controllers/actions';
import { MetaKey } from '../../domain/model';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import screen from '../screens/MetaKeyNewScreen';

export interface MetaKeyNewActions {
  loadInitData: () => void;
  register: (v: MetaKey) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(MetaKeyUsecase);

function mapDispatchToProps(dispatch: Dispatch): MetaKeyNewActions {
  return {
    loadInitData() {
      dispatch(metaKeyActions.startFetchInitData({}));
      usecase
        .fetchInitData()
        .then((result) => {
          dispatch(metaKeyActions.doneFetchInitData({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            metaKeyActions.failedFetchInitData({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
    register(v: MetaKey) {
      dispatch(metaKeyActions.startRegister({}));
      const register = v.id === 0 ? usecase.create : usecase.update;
      register(v)
        .then((result) => {
          if (result) {
            dispatch(metaKeyActions.doneRegister({ params: {}, result }));
          } else {
            throw new Error();
          }
        })
        .catch(() => {
          dispatch(
            metaKeyActions.failedRegister({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.metaKeyNew);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
