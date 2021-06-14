import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { History } from 'history';
import { AppState } from '../../interfaces/controllers/stores/store';
import { metaKeyActions } from '../../interfaces/controllers/actions';
import { MetaKey } from '../../domain/model';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import screen from '../screens/MetaKeyNewScreen';

export interface MetaKeyNewActions {
  loadInitData: () => void;
  register: ({ name }: { name: string }) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(MetaKeyUsecase);

function mapDispatchToProps(dispatch: Dispatch, ownProps: { history: History }): MetaKeyNewActions {
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
    register({ name }: { name: string }) {
      dispatch(metaKeyActions.startRegister({}));
      const v = new MetaKey('', name);
      usecase
        .create(v)
        .then((result) => dispatch(metaKeyActions.doneRegister({ params: {}, result })))
        .then(() =>
          ownProps.history.push({
            pathname: '/meta-keys',
            state: { registerResultMessage: '🎉 Registered MetaKey successfully' },
          })
        )
        .catch(() => dispatch(metaKeyActions.failedRegister({ params: {}, error: { message: '' } })));
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.metaKeyNew);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
