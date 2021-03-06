import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { tagActions } from '../../interfaces/controllers/actions';
import { Tag } from '../../domain/model';
import { TagUsecase } from '../../usecase/TagUsecase';
import screen from '../screens/TagNewScreen';

export interface TagNewActions {
  loadInitData: () => void;
  register: (v: Tag) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(TagUsecase);

function mapDispatchToProps(dispatch: Dispatch): TagNewActions {
  return {
    loadInitData() {
      dispatch(tagActions.startFetchInitData({}));
      usecase
        .fetchInitData()
        .then((result) => {
          dispatch(tagActions.doneFetchInitData({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            tagActions.failedFetchInitData({
              params: {},
              error: { message: '' },
            })
          );
        });
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
  return Object.assign({}, appState.tagNew);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
