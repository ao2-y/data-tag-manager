import React, { useEffect } from 'react';
import * as H from 'history';
import { ItemTemplatesState } from '../../interfaces/controllers/states';
import { ItemTemplatesActions } from '../container/ItemTemplatesContainer';
import ItemTemplateSearchForm from '../components/ItemTemplateSearchForm';
import { ItemTemplateReadComponent } from '../components/ItemTemplateCollection';
import { CButton } from '@coreui/react';
import CIcon from '@coreui/icons-react';
interface OwnProps {
  history: H.History;
}
type ItemTemplatesProps = OwnProps & ItemTemplatesState & ItemTemplatesActions;

const ItemTemplatesScreen: React.FC<ItemTemplatesProps> = (props) => {
  useEffect(() => { props.search(); }, []);

  return (
    <div className="">
      <div className="d-flex mb-4">
        <div className="flex-grow-1 ">
          <ItemTemplateSearchForm onChange={() => {}}></ItemTemplateSearchForm>
        </div>
        <div className="d-flex align-items-end flex-column">
          <CButton
            type="button"
            variant="outline"
            size="sm"
            className="ml-2 mt-auto"
            color="primary"
            onClick={() => {}}>
            <CIcon name="cil-plus" /> Add ItemTemplate
          </CButton>
        </div>
      </div>
      <ItemTemplateReadComponent itemTemplates={props.items}></ItemTemplateReadComponent>
    </div>
  );
};

export default ItemTemplatesScreen;
