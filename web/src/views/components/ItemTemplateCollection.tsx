import React from 'react';
import { CBadge } from '@coreui/react';
import { ItemTemplate } from '../../domain/model';
import DataCollection from './DataCollection';

interface ComponentProps {
  itemTemplates: ItemTemplate[];
  isLoading?: boolean;
}

const makeHeader = (object: {}) => {
  const item = object as ItemTemplate;
  return (<span>{item.name}<small> ({( item.metas || []).length} keys) </small></span>)
};

const makeBody = (object: {}) => {
  const item = object as ItemTemplate;
  return item.metas?.map((meta) => (<CBadge key={meta.id} color="secondary" className="mr-1"> {` ${meta.metaKey.name} `}</CBadge>))
};

interface ReadComponentProps extends ComponentProps {
}
export const ItemTemplateReadComponent = React.memo((props: React.PropsWithChildren<ReadComponentProps>) => {
  return (
    <DataCollection
      datasource={props.itemTemplates}
      isLoading={props.isLoading}
      makeHeader={makeHeader}
      makeBody={makeBody}
      makeEditLink={(item) => `/item-templates/${(item as ItemTemplate).id}`}
    ></DataCollection>
  );
});

interface SelectComponentProps extends ComponentProps {
  onSelect: (item: ItemTemplate) => void;
}
export const ItemTemplateSelectComponent = React.memo((props: React.PropsWithChildren<SelectComponentProps>) => {
  return (
    <DataCollection
      datasource={props.itemTemplates}
      isLoading={props.isLoading}
      makeHeader={makeHeader}
      makeBody={makeBody}
      onSelect={(item) => { props.onSelect(item[0] as ItemTemplate); } }
    ></DataCollection>
  );
});
