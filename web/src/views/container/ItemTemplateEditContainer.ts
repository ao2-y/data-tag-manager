import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemTemplateActions } from '../../interfaces/controllers/actions';
import { ItemTemplate } from '../../domain/model';
import { ItemTemplateUsecase } from '../../usecase/ItemTemplateUsecase';
import screen from '../screens/ItemTemplateEditScreen';

export interface ItemTemplateEditActions {
  fetch: (id: number) => void;
  register: (v: ItemTemplate) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(ItemTemplateUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    fetch(id: number) {
      dispatch(itemTemplateActions.startFetch({}));
      if (id) {
        usecase
          .fetch(id)
          .then((result) => {
            if (result) {
              dispatch(itemTemplateActions.doneFetch({ params: {}, result }));
            } else {
              throw new Error();
            }
          })
          .catch(() => {
            dispatch(
              itemTemplateActions.failedFetch({
                params: {},
                error: { message: '' },
              })
            );
          });
      } else {
        setTimeout(() => {
          dispatch(
            itemTemplateActions.doneFetch({
              params: {},
              result: {
                id: 0,
                metas: [],
                name: '',
              },
            })
          );
        });
      }
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
  return Object.assign({}, appState.itemTemplateEdit);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
