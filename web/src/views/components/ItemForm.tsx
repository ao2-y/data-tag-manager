import React, { useState } from 'react';
import {
  CFormGroup,
  CCol,
  CInput,
  CForm,
  CButton,
  CTextarea,
  CCard,
  CCardBody,
  CCardHeader,
  CInputGroup,
  CInputGroupAppend,
} from '@coreui/react';
import CIcon from '@coreui/icons-react';
import MetaKeyLabel from '../components/MetaKeyLabel';
import InputTag from '../components/InputTag';
import { Item, Tag } from '../../domain/model';

interface ComponentProps {
  item: Item;
  onSave: (item: Item) => void;
}

export const ItemForm: React.FC<ComponentProps> = (props) => {
  const [item, setItem] = useState<Item>(props.item);
  const updateString = (value: string, key: string) => {
    const target = JSON.parse(JSON.stringify(item));
    target[key] = value;
    setItem(target);
  };
  const updateMetaKey = (value: string, metaKeyId: string) => {
    const target = JSON.parse(JSON.stringify(item)) as Item;
    const meta = target.metaDatas?.find((v) => v.metaKeyId);
    if (meta) {
      meta.value = value;
    }
    setItem(target);
  };
  const updateTag = (value: Tag[]) => {
    // const target = JSON.parse(JSON.stringify(item)) as Item;
    // target.tags = value;
    // setItem(target);
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
                  value={item.name}
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
                <small className="d-block">Description</small>
                <CTextarea
                  type="text"
                  value={item.description}
                  placeholder="Please entry description..."
                  onChange={(e) => {
                    updateString((e.target as HTMLTextAreaElement).value, 'description');
                  }}
                />
              </div>
            </CCol>
          </CFormGroup>
          <CFormGroup row>
            <CCol>
              <div>
                <small className="d-block">Tags</small>
                <InputTag value={[]} onChange={(v) => {}} />
              </div>
            </CCol>
          </CFormGroup>
        </CCardBody>
      </CCard>
      <CCard>
        <CCardHeader>Metas</CCardHeader>
        <CCardBody>
          {props.item?.metaDatas.map((meta) => (
            <CFormGroup row key={`item-meta-input${meta.metaKeyId})`}>
              <CCol>
                <div>
                  <small className="d-block">
                    <MetaKeyLabel metaKeyId={meta.metaKeyId} />
                  </small>
                  <CInputGroup>
                    <CInput
                      type="text"
                      value={meta.value}
                      placeholder="Please entry name..."
                      onChange={(e) => {
                        updateMetaKey((e.target as HTMLInputElement).value, meta.metaKeyId);
                      }}
                    />
                    <CInputGroupAppend>
                      <CButton type="button" color="danger">
                        <CIcon name="cil-trash" />
                      </CButton>
                    </CInputGroupAppend>
                  </CInputGroup>
                </div>
              </CCol>
            </CFormGroup>
          ))}
          <CButton type="button" size="sm" color="info" onClick={() => console.log()}>
            <CIcon name="cil-plus" /> Add
          </CButton>
        </CCardBody>
      </CCard>

      <CButton type="button" size="sm" color="success" onClick={() => console.log()}>
        <CIcon name="cil-scrubber" /> Submit
      </CButton>
    </CForm>
  );
};
