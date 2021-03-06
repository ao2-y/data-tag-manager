import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { tagActions } from '../../interfaces/controllers/actions';
import { TagUsecase } from '../../usecase/TagUsecase';
import screen from '../screens/TagsScreen';

export interface TagsActions {
  search: () => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(TagUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    search() {
      dispatch(tagActions.startSearch({}));
      usecase
        .find()
        .then((result) => {
          dispatch(tagActions.doneSearch({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            tagActions.failedSearch({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.tags);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
