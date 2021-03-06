import React, { useState } from 'react';
import { ItemTemplateSelectComponent } from '../components/ItemTemplateCollection';
import { CButton, CModal, CModalHeader, CModalTitle, CModalBody, CModalFooter } from '@coreui/react';
import { ItemTemplate } from '../../domain/model';
import ItemTemplateSearchForm from './ItemTemplateSearchForm';

interface ComponentProps {
  title?: string;
  itemTemplates: ItemTemplate[];
  isLoading?: boolean;
  isShow: boolean;
  onSelect: (itemTemplate?: ItemTemplate) => void;
}

export const ItemTemplateSelectModal: React.FC<ComponentProps> = React.memo((props) => {
  const [selected, setSelected] = useState<ItemTemplate|undefined>();
  return (
    <CModal show={props.isShow} onClose={() => {props.onSelect(selected)}} size="xl">
      <CModalHeader>
        <CModalTitle>{props.title || 'Please select item template.'}</CModalTitle>
      </CModalHeader>
      <CModalBody>
        <div className="mb-4">
          <ItemTemplateSearchForm onChange={() => {}}></ItemTemplateSearchForm>
        </div>
        <ItemTemplateSelectComponent 
          itemTemplates={props.itemTemplates} 
          isLoading={props.isLoading} 
          onSelect={(value) => { console.log(value); setSelected(value); }}
        />
      </CModalBody>
      <CModalFooter>
        <CButton color="secondary" onClick={() => { props.onSelect(); }}>Close</CButton>
        <CButton color="primary" onClick={() => { props.onSelect(selected); }}>Done</CButton>
      </CModalFooter>
    </CModal>
  );
});
