import React, { useEffect, useState } from 'react';
import * as H from 'history';
import { ItemsState } from '../../interfaces/controllers/states';
import { ItemsActions } from '../container/ItemsContainer';
import { ItemTemplateSelectModal } from '../components/ItemTemplateSelectModal';
import ItemSearchForm from '../components/ItemSearchForm';
import { CButton } from '@coreui/react';
import { ItemReadComponent } from '../components/ItemCollection';
import CIcon from '@coreui/icons-react';
interface OwnProps {
  history: H.History;
}

type ItemsProps = OwnProps & ItemsState & ItemsActions;

const ItemsScreen: React.FC<ItemsProps> = (props) => {
  useEffect(() => {
    props.searchItems();
  }, [true]);
  const [modal, setModal] = useState(false);
  return (
    <div>
      <ItemTemplateSelectModal
        isLoading={props.isItemTemplatesLoading}
        itemTemplates={props.itemTemplates}
        isShow={modal}
        onSelect={(item) => {
          setModal(false);
          console.log(item);
          if (item) {
            console.log(item.id);
            props.history.push({
              pathname: '/items/new',
              state: { template: item.id },
            });
          }
        }}
      />
      <div className="d-flex mb-4">
        <div className="flex-grow-1 ">
          <ItemSearchForm onChange={() => {}}></ItemSearchForm>
        </div>
        <div className="d-flex align-items-end flex-column">
          <CButton
            type="button"
            variant="outline"
            size="sm"
            className="ml-2 mt-auto"
            color="primary"
            onClick={() => {
              console.log('hoge');
              props.searchItemTemplates();
              setModal(!modal);
            }}>
            <CIcon name="cil-plus" /> Add Item
          </CButton>
        </div>
      </div>
      <ItemReadComponent {...props}></ItemReadComponent>
    </div>
  );
};

export default ItemsScreen;
