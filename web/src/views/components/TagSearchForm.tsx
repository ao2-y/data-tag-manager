import React, { useState } from 'react';
import {
  CCard,
  CCardBody,
  CFormGroup,
  CForm,
  CCol,
  CInput,
  CRow,
} from '@coreui/react';
import { MetaKey } from '../../domain/model';

interface OwnProps {
  name?: string;
  onChange: (value: MetaKey[]) => void;
}
type ComponentProps = OwnProps;

const Component: React.FC<ComponentProps> = (props) => {
  const [name, setNameValue] = useState<string>(props.name || '');
  return (
    <CCard accentColor="primary" className="mb-0">
      <CCardBody>
        <CRow>
          <CCol>
            <h4 id="traffic" className="card-title mb-0">
              Search Condition
            </h4>
            <div className="small text-muted">
              Please enter search conditions
            </div>
          </CCol>
        </CRow>
        <CRow className="mt-2">
          <CCol>
            <CForm>
              <CFormGroup row>
                <CCol>
                  <div>
                    <small className="d-block">Name</small>
                    <CInput
                      type="text"
                      value={name}
                      placeholder="Please entry name..."
                      onChange={(e) => {
                        setNameValue((e.target as HTMLInputElement).value);
                      }}
                    />
                  </div>
                </CCol>
              </CFormGroup>
            </CForm>
          </CCol>
        </CRow>
      </CCardBody>
    </CCard>
  );
};

export default Component;
