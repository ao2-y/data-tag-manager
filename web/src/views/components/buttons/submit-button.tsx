import React from 'react';
import { CButton } from '@coreui/react';
import CIcon from '@coreui/icons-react';

interface ComponentProps {
  onClick: () => void;
}

export const SubmitButton: React.FC<ComponentProps> = (props) => {
  return (
    <CButton type="button" size="sm" color="success" onClick={props.onClick}>
      <CIcon name="cil-scrubber" /> Submit
    </CButton>
  );
};
