import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemActions } from '../../interfaces/controllers/actions';
import { Item } from '../../domain/model';
import { ItemUsecase } from '../../usecase/ItemUsecase';
import screen from '../screens/ItemEditScreen';

export interface ItemEditActions {
  fetch: (id: number) => void;
  register: (v: Item) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(ItemUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    fetch(id: number) {
      dispatch(itemActions.startFetch({}));
      usecase
      .fetch(id)
      .then((result) => {
        console.log(result)
        if (result) {
          dispatch(itemActions.doneFetch({ params: {}, result }));
        } else {
          throw new Error();
        }
      })
      .catch(() => {
        dispatch(
          itemActions.failedFetch({
            params: {},
            error: { message: '' },
          })
        );
      });
    },
    register(v: Item) {
      dispatch(itemActions.startRegister({}));
      usecase.update(v)
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
  return Object.assign({}, appState.itemEdit);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
