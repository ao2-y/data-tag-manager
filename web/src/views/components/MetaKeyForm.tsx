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
import { MetaKey } from '../../domain/model';

interface ComponentProps {
  metaKey: MetaKey;
  onSave: (metaKey: MetaKey) => void;
}

export const MetaKeyForm: React.FC<ComponentProps> = (props) => {
  const [metaKey, setMetaKey] = useState<MetaKey>(props.metaKey);
  const updateString = (value: string, key: string) => {
    const target = JSON.parse(JSON.stringify(metaKey));
    target[key] = value;
    setMetaKey(target);
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
                  value={metaKey.name}
                  placeholder="Please entry name..."
                  onChange={(e) => {
                    updateString((e.target as HTMLInputElement).value, 'name');
                  }}
                />
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
