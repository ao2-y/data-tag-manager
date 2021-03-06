import React, { useState, useEffect } from 'react';
import { container } from 'tsyringe';
import { SketchPicker } from 'react-color';
import {
  CFormGroup,
  CCol,
  CInput,
  CForm,
  CButton,
  CCard,
  CCardBody,
  CCardHeader,
  CLink,
} from '@coreui/react';
import CIcon from '@coreui/icons-react';
import { Tag } from '../../domain/model';
import { TagSelectModal } from './TagSelectModal';
import { TagUsecase } from '../../usecase/TagUsecase';

interface ComponentProps {
  tag: Tag;
  onSave: (item: Tag) => void;
}

const usecase = container.resolve(TagUsecase);

export const TagForm: React.FC<ComponentProps> = (props) => {
  const [parentTags, setParentTags] = useState<Tag[] | undefined>(undefined);
  const [displayColorPicker, setDisplayColorPicker] = useState<boolean>(false);
  const [tag, setTag] = useState<Tag>(props.tag);
  const [parentModal, setParentModal] = useState(false);

  useEffect(() => {
    usecase.find().then((v) => {
      setParentTags(v);
    });
  }, [true]);

  const updateString = (value: string, key: string) => {
    const target = JSON.parse(JSON.stringify(tag));
    target[key] = value;
    setTag(target);
  };
  const updateParent = (value?: Tag) => {
    const target = JSON.parse(JSON.stringify(tag)) as Tag;
    target.parent = value;
    setTag(target);
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
                  value={tag.name}
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
                <small className="d-block">Parent</small>
                {!tag.parent && (
                  <div
                    style={{
                      border: '2px dashed lightgray',
                      width: '100%',
                      height: '46px',
                      textAlign: 'center',
                      lineHeight: '46px',
                      color: 'gray',
                      cursor: 'pointer',
                    }}
                    onClick={() => {
                      setParentModal(true);
                    }}>
                    Please select parent tag
                  </div>
                )}
                {tag.parent && (
                  <CCard
                    style={{
                      borderLeft: `${tag.parent.color} 5px solid`,
                      marginBottom: '0',
                    }}
                    onClick={() => {}}>
                    <CCardHeader>
                      {tag.parent.name}
                      <div
                        className="card-header-actions"
                        onClick={(e) => {
                          e.stopPropagation();
                        }}>
                        <CLink
                          className="card-header-action"
                          onClick={() => {
                            updateParent();
                          }}>
                          <CIcon name="cil-trash" />
                        </CLink>
                      </div>
                    </CCardHeader>
                  </CCard>
                )}
                <TagSelectModal
                  isShow={parentModal}
                  isLoading={!parentTags}
                  tags={parentTags ? parentTags : []}
                  onSelect={(item) => {
                    updateParent(item);
                    setParentModal(false);
                  }}
                />
              </div>
            </CCol>
          </CFormGroup>
          <CFormGroup row>
            <CCol>
              <div>
                <small className="d-block">Color</small>
                <div
                  style={{
                    padding: '5px',
                    background: ' #fff',
                    borderRadius: '1px',
                    cursor: 'pointer',
                    display: 'inline-block',
                    boxShadow: '0 0 0 1px rgba(0,0,0,.1)',
                  }}
                  onClick={() => {
                    console.log('aaaa');
                    console.log(displayColorPicker);
                    setDisplayColorPicker(!displayColorPicker);
                  }}>
                  <div
                    style={{
                      width: '36px',
                      height: '14px',
                      borderRadius: '2px',
                      background: tag.color,
                    }}
                  />
                </div>
                {displayColorPicker && (
                  <div style={{ position: 'absolute', zIndex: 2 }}>
                    <div
                      style={{
                        position: 'fixed',
                      }}
                      onClick={() => setDisplayColorPicker(false)}>
                      <SketchPicker
                        color={tag.color}
                        onChange={(color) => {
                          updateString(color.hex, 'color');
                        }}
                      />
                    </div>
                  </div>
                )}
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
