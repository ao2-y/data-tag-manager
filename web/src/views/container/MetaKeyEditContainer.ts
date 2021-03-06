import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { metaKeyActions } from '../../interfaces/controllers/actions';
import { MetaKey } from '../../domain/model';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import screen from '../screens/MetaKeyEditScreen';

export interface MetaKeyEditActions {
  fetch: (id: number) => void;
  register: (v: MetaKey) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(MetaKeyUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    fetch(id: number) {
      dispatch(metaKeyActions.startFetch({}));
      if (id) {
        usecase
          .fetch(id)
          .then((result) => {
            if (result) {
              dispatch(metaKeyActions.doneFetch({ params: {}, result }));
            } else {
              throw new Error();
            }
          })
          .catch(() => {
            dispatch(
              metaKeyActions.failedFetch({
                params: {},
                error: { message: '' },
              })
            );
          });
      } else {
        setTimeout(() => {
          dispatch(
            metaKeyActions.doneFetch({
              params: {},
              result: {
                id: 0,
                name: '',
              },
            })
          );
        });
      }
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
  return Object.assign({}, appState.metaKeyEdit);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
