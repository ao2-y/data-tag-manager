import React, { useState } from 'react';
import {
  CFormGroup,
  CCol,
  CInput,
  CForm,
  CButton,
  CCard,
  CCardBody,
  CCardHeader,
} from '@coreui/react';
import CIcon from '@coreui/icons-react';
import { ItemTemplate } from '../../domain/model';

interface ComponentProps {
  itemTemplate: ItemTemplate;
  onSave: (item: ItemTemplate) => void;
}

export const ItemTemplateForm: React.FC<ComponentProps> = (props) => {
  const [itemTemplate, setItemTemplate] = useState<ItemTemplate>(
    props.itemTemplate
  );
  const updateString = (value: string, key: string) => {
    const target = JSON.parse(JSON.stringify(itemTemplate));
    target[key] = value;
    setItemTemplate(target);
  };
  return (
    <CForm>
      <CCard>
        <CCardHeader>General</CCardHeader>
        <CCardBody>
          <CFormGroup row>
            <CCol>
              <div>
                <small className="d-block">Name</small>
                <CInput
                  type="text"
                  value={itemTemplate.name}
                  placeholder="Please entry name..."
                  onChange={(e) => {
                    updateString((e.target as HTMLInputElement).value, 'name');
                  }}
                />
              </div>
            </CCol>
          </CFormGroup>
          <CFormGroup row>
            <CCol>
              <div>
                <small className="d-block">Meta</small>
                {/* <InputMeta
                  value={item.metas}
                  onChange={(v) => {
                    setMetasValue(v);
                  }}></InputMeta> */}
              </div>
            </CCol>
          </CFormGroup>
        </CCardBody>
      </CCard>
      <CButton
        type="button"
        size="sm"
        color="success"
        onClick={() => console.log(name)}>
        <CIcon name="cil-scrubber" /> Submit
      </CButton>
    </CForm>
  );
};
