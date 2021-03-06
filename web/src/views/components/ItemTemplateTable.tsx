import React from 'react'
import { CDataTable, CBadge, CCard, CCardBody } from '@coreui/react';
import { ItemTemplate } from '../../domain/model';

interface OwnProps { 
  items: ItemTemplate[];
  onRowClick?: (item: ItemTemplate) => void; 
}

type ItemTemplateTableProps = OwnProps;

const fields = [
  { key: 'id', _style: {} },
  { key: 'name', _style: {} },
  { key: 'metas', _style: {} },
];

const makeIDColumn = (item: ItemTemplate) => (<td>{item.id}</td>);
const makeNameColumn = (item: ItemTemplate) => (<td>{item.name}</td>);
const makeMetasColumn = (item: ItemTemplate) => (<td>{item.metas.map(meta => (<CBadge key={meta.id} color="secondary" className="mr-1"> {meta.metaKey.name}</CBadge>))}</td>);

const ItemTemplateTable: React.FC<ItemTemplateTableProps> = (props) => {
  return (
    <CCard>
      <CCardBody>
        <CDataTable
        onRowClick={props.onRowClick}
        items={props.items}
        fields={fields}
        border={true}
        hover={true}
        striped
        itemsPerPage={20}
        pagination
        scopedSlots={{
          'id': makeIDColumn, 'name': makeNameColumn, 'metas': makeMetasColumn,
        }}
      />
      </CCardBody>
    </CCard>
  );
};

export default React.memo(ItemTemplateTable)
