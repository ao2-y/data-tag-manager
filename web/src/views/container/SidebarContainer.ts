import { Action } from 'typescript-fsa';
import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { AppState } from '../../interfaces/controllers/stores/store';
import { layoutActions } from '../../interfaces/controllers/actions/LayoutActions';
import Sidebar from '../components/Sidebar';

export interface LayoutActions {
  showSidebar: (v: 'responsive' | boolean) => Action<'responsive' | boolean>;
}

function mapDispatchToProps(dispatch: Dispatch<Action<'responsive' | boolean>>) {
  return {
    showSidebar: (v: 'responsive' | boolean) => dispatch(layoutActions.showSidebar(v)),
  };
}

function mapStateToProps(appState: AppState) {
  return Object.assign({}, appState.layout);
}

export default connect(mapStateToProps, mapDispatchToProps)(Sidebar);
