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
import InputMeta from '../components/InputMeta';

interface OwnProps {
  name?: string;
  metas?: MetaKey[];
  onChange: (value: MetaKey[]) => void;
}
type TagSearchFormProps = OwnProps;

const Component: React.FC<TagSearchFormProps> = (props) => {
  const [name, setNameValue] = useState<string>(props.name || '');
  const [metas, setMetasValue] = useState<MetaKey[]>(props.metas || []);
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
                <CCol>
                  <div>
                    <small className="d-block">Meta</small>
                    <InputMeta
                      value={metas}
                      onChange={(v) => {
                        setMetasValue(v);
                      }}></InputMeta>
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
