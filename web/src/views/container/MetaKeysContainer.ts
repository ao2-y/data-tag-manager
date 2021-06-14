import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { metaKeyActions } from '../../interfaces/controllers/actions';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import screen from '../screens/MetaKeysScreen';

export interface MetaKeysActions {
  search: (keyword?: string) => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(MetaKeyUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    search(keyword?: string) {
      dispatch(metaKeyActions.startSearch({}));
      usecase
        .find({ keyword })
        .then((result) => {
          dispatch(metaKeyActions.doneSearch({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            metaKeyActions.failedSearch({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.metaKeys);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
