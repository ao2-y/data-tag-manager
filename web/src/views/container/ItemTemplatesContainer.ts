import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { container } from 'tsyringe';
import { AppState } from '../../interfaces/controllers/stores/store';
import { itemTemplateActions } from '../../interfaces/controllers/actions';
import { ItemTemplateUsecase } from '../../usecase/ItemTemplateUsecase';
import screen from '../screens/ItemTemplatesScreen';

export interface ItemTemplatesActions {
  search: () => void;
}

// このクラスをinterfaceに移設するか、Controllerでやるか検討する
const usecase = container.resolve(ItemTemplateUsecase);

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    search() {
      dispatch(itemTemplateActions.startSearch({}));
      usecase
        .find()
        .then((result) => {
          dispatch(itemTemplateActions.doneSearch({ params: {}, result }));
        })
        .catch(() => {
          dispatch(
            itemTemplateActions.failedSearch({
              params: {},
              error: { message: '' },
            })
          );
        });
    },
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.itemTemplates);
}

export default connect(mapStateToProps, mapDispatchToProps)(screen);
