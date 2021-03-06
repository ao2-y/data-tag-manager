import React, { useState } from 'react';
import { TagSelectComponent } from './TagCollection';
import {
  CButton,
  CModal,
  CModalHeader,
  CModalTitle,
  CModalBody,
  CModalFooter,
} from '@coreui/react';
import { Tag } from '../../domain/model';
import TagSearchForm from './TagSearchForm';

interface ComponentProps {
  title?: string;
  tags: Tag[];
  isLoading?: boolean;
  isShow: boolean;
  onSelect: (tag?: Tag) => void;
}

export const TagSelectModal: React.FC<ComponentProps> = React.memo((props) => {
  const [selected, setSelected] = useState<Tag | undefined>();
  return (
    <CModal
      show={props.isShow}
      onClose={() => {
        props.onSelect(selected);
      }}
      size="xl">
      <CModalHeader>
        <CModalTitle>{props.title || 'Please select tag.'}</CModalTitle>
      </CModalHeader>
      <CModalBody>
        <div className="mb-4">
          <TagSearchForm onChange={() => {}} />
        </div>
        <TagSelectComponent
          tags={props.tags}
          isLoading={props.isLoading}
          onSelect={(value) => {
            setSelected(value);
          }}
        />
      </CModalBody>
      <CModalFooter>
        <CButton
          color="secondary"
          onClick={() => {
            props.onSelect();
          }}>
          Close
        </CButton>
        <CButton
          color="primary"
          onClick={() => {
            props.onSelect(selected);
          }}>
          Done
        </CButton>
      </CModalFooter>
    </CModal>
  );
});
