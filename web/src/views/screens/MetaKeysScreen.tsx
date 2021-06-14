import React, { useEffect } from 'react';
import { History, Location } from 'history';
import { toast } from 'react-toastify';
import { MetaKeysState } from '../../interfaces/controllers/states';
import { MetaKeysActions } from '../container/MetaKeysContainer';
import { MetaKeyReadComponent } from '../components/MetaKeyCollection';
import MetaKeySearchForm from '../components/MetaKeySearchForm';
import { CButton } from '@coreui/react';
import CIcon from '@coreui/icons-react';
interface OwnProps {
  history: History;
  location: Location<{ registerResultMessage: string } | undefined>;
}
type MetaKeysProps = OwnProps & MetaKeysState & MetaKeysActions;

const MetaKeysScreen: React.FC<MetaKeysProps> = (props) => {
  useEffect(() => props.search(), []);
  useEffect(() => {
    if (props.location.state?.registerResultMessage) {
      toast.success(props.location.state?.registerResultMessage, {
        position: 'top-right',
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
      });
    }
  }, [props.location.state?.registerResultMessage]);
  return (
    <div className="">
      <div className="d-flex mb-4">
        <div className="flex-grow-1 ">
          <MetaKeySearchForm onNameChange={props.search}></MetaKeySearchForm>
        </div>
        <div className="d-flex align-items-end flex-column">
          <CButton
            style={{ width: '150px', whiteSpace: 'nowrap' }}
            type="button"
            variant="outline"
            size="sm"
            className="ml-2 mt-auto"
            color="primary"
            onClick={() => props.history.push('/meta-keys/new')}>
            <CIcon name="cil-plus" /> Add MetaKey
          </CButton>
        </div>
      </div>
      <MetaKeyReadComponent metaKeys={props.metaKeys}></MetaKeyReadComponent>
    </div>
  );
};

export default MetaKeysScreen;
