import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { tagActions } from '../../interfaces/controllers/actions';
import { Tag } from '../../domain/model';
import { TagUsecase } from '../../usecase/TagUsecase';
import screen from '../screens/TagEditScreen';

export interface TagEditActions {
  fetch: (id: number) => void;
  register: (v: Tag) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(TagUsecase);

function mapDispatchToProps(dispatch: Dispatch): TagEditActions {
  return {
    fetch(id: number) {
      dispatch(tagActions.startFetch({}));
      if (id) {
        usecase
          .fetch(id)
          .then((result) => {
            if (result) {
              dispatch(tagActions.doneFetch({ params: {}, result }));
            } else {
              throw new Error();
            }
          })
          .catch(() => {
            dispatch(
              tagActions.failedFetch({
                params: {},
                error: { message: '' },
              })
            );
          });
      } else {
        setTimeout(() => {
          dispatch(
            tagActions.doneFetch({
              params: {},
              result: {
                id: 0,
                name: '',
                level: 0,
                parent: undefined,
                color: '#a52a2a',
              },
            })
          );
        });
      }
    },
    register(v: Tag) {
      dispatch(tagActions.startRegister({}));
      const register = v.id === 0 ? usecase.create : usecase.update;
      register(v)
        .then((result) => {
          if (result) {
            dispatch(tagActions.doneRegister({ params: {}, result }));
          } else {
            throw new Error();
          }
        })
        .catch(() => {
          dispatch(
            tagActions.failedRegister({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.tagEdit);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
