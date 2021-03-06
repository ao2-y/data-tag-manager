import React, { useEffect } from 'react';
import * as H from 'history';
import { MetaKeysState } from '../../interfaces/controllers/states';
import { MetaKeysActions } from '../container/MetaKeysContainer';
import { MetaKeyReadComponent } from '../components/MetaKeyCollection';
import MetaKeySearchForm from '../components/MetaKeySearchForm';
import { CButton } from '@coreui/react';
import CIcon from '@coreui/icons-react';
interface OwnProps {
  history: H.History;
}
type MetaKeysProps = OwnProps & MetaKeysState & MetaKeysActions;

const MetaKeysScreen: React.FC<MetaKeysProps> = (props) => {
  useEffect(() => {
    props.search();
  }, []);

  return (
    <div className="">
      <div className="d-flex mb-4">
        <div className="flex-grow-1 ">
          <MetaKeySearchForm onChange={() => {}}></MetaKeySearchForm>
        </div>
        <div className="d-flex align-items-end flex-column">
          <CButton
            type="button"
            variant="outline"
            size="sm"
            className="ml-2 mt-auto"
            color="primary"
            onClick={() => {
              props.history.push('/meta-keys/new');
            }}>
            <CIcon name="cil-plus" /> Add MetaKey
          </CButton>
        </div>
      </div>
      <MetaKeyReadComponent metaKeys={props.metaKeys}></MetaKeyReadComponent>
    </div>
  );
};

export default MetaKeysScreen;
