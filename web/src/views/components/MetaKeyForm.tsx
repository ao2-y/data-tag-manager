import React, { useState } from 'react';
import { CFormGroup, CCol, CInput, CForm, CCard, CCardBody, CCardHeader } from '@coreui/react';

interface ComponentProps {
  name: string;
  onChange: (changes: { [key: string]: string }) => void;
}

export const MetaKeyForm: React.FC<ComponentProps> = (props) => {
  const [name, setName] = useState(props.name);

  const updateString = (key: string, value: string, callback: (key: string) => void) => {
    callback(value);
    props.onChange({ [key]: value });
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
                  value={name}
                  placeholder="Please entry name..."
                  onChange={(e) => updateString('name', (e.target as HTMLInputElement).value, setName)}
                />
              </div>
            </CCol>
          </CFormGroup>
        </CCardBody>
      </CCard>
    </CForm>
  );
};
