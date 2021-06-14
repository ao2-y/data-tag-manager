import React from 'react';
import { CCol, CFormGroup, CBadge } from '@coreui/react';
import { Item } from '../../domain/model';
import MetaKeyLabel from '../components/MetaKeyLabel';
import DataCollection from './DataCollection';

interface ComponentProps {
  items: Item[];
  isLoading?: boolean;
}

const makeHeader = (object: {}) => {
  const item = object as Item;
  return <span>{item.name}</span>;
};

const makeBody = (object: {}) => {
  const item = object as Item;
  return (
    <div>
      {item.metaDatas?.map((value) => (
        <CFormGroup row key={`item-meta-item${value.id}`}>
          <CCol>
            <small className="d-block">
              <MetaKeyLabel metaKeyId={value.metaKeyId} />
            </small>
            <div>{value.value}</div>
          </CCol>
        </CFormGroup>
      ))}
      {item.tags?.map((value) => (
        <CBadge key={`item-tag-item${value.id}`} color="secondary" className="mr-1">
          {` ${value.tag.name} `}
        </CBadge>
      ))}
    </div>
  );
};

interface ReadComponentProps extends ComponentProps {}
export const ItemReadComponent = React.memo((props: React.PropsWithChildren<ReadComponentProps>) => {
  return (
    <DataCollection
      datasource={props.items}
      isLoading={props.isLoading}
      makeHeader={makeHeader}
      makeBody={makeBody}
      makeEditLink={(item) => `/items/${(item as Item).id}`}></DataCollection>
  );
});

interface SelectComponentProps extends ComponentProps {
  onSelect: (item: Item) => void;
}
export const ItemSelectComponent = React.memo((props: React.PropsWithChildren<SelectComponentProps>) => {
  return (
    <DataCollection
      datasource={props.items}
      isLoading={props.isLoading}
      makeHeader={makeHeader}
      makeBody={makeBody}
      onSelect={(item) => {
        props.onSelect(item[0] as Item);
      }}></DataCollection>
  );
});
