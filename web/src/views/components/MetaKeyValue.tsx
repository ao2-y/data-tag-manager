import React, { useState } from 'react';
import {
  CLink,
  CBadge,
  CCard,
  CCardBody,
  CRow,
  CCol,
  CCardHeader,
  CCollapse,
  CFormGroup,
} from '@coreui/react';

interface OwnProps {
  key: number;
  value: number;
}

type ComponentProps = OwnProps;

const Component: React.FC<ComponentProps> = (props) => {
  return (
    <div>
      <small className="d-block">{props.key}</small>
      <div>{props.value}</div>
    </div>
  );
};

export default React.memo(Component);
