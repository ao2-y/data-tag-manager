import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { History } from 'history';
import { AppState } from '../../interfaces/controllers/stores/store';
import { metaKeyActions } from '../../interfaces/controllers/actions';
import { MetaKey } from '../../domain/model';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import screen from '../screens/MetaKeyEditScreen';

export interface MetaKeyEditActions {
  fetch: (id: string) => void;
  register: ({ id, name }: { id: string; name: string }) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(MetaKeyUsecase);

function mapDispatchToProps(dispatch: Dispatch, ownProps: { history: History }) {
  return {
    fetch(id: string) {
      dispatch(metaKeyActions.startFetch({}));
      usecase
        .fetch(id)
        .then((result) => dispatch(metaKeyActions.doneFetch({ params: {}, result })))
        .catch(() => dispatch(metaKeyActions.failedFetch({ params: {}, error: { message: '' } })));
    },
    register({ id, name }: { id: string; name: string }) {
      dispatch(metaKeyActions.startRegister({}));
      const v = new MetaKey(id, name);
      usecase
        .update(v)
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
  return Object.assign({}, appState.metaKeyEdit);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
