import React from 'react';
import { Tag } from '../../domain/model';
import DataCollection from './DataCollection';
import { CLink, CCard, CCardHeader } from '@coreui/react';
import CIcon from '@coreui/icons-react';

interface ComponentProps {
  tags: Tag[];
  isLoading?: boolean;
}

const makeHeader = (object: {}) => {
  const item = object as Tag;
  return <span>{item.name}</span>;
};

const makeBody = (object: {}) => {
  const item = object as Tag;
  return (
    <CCard
      style={{
        borderLeft: `${(item as Tag).color} 5px solid`,
        marginBottom: '0',
      }}
      onClick={() => {}}>
      <CCardHeader>
        {item.name}
        <div
          className="card-header-actions"
          onClick={(e) => {
            e.stopPropagation();
          }}>
          <CLink className="card-header-action" to={`/tags/${item.id}`}>
            <CIcon name="cil-pencil" />
          </CLink>
        </div>
      </CCardHeader>
    </CCard>
  );
};

const makeBodyStyle = () => {
  return { paddingRight: '5px' };
};

interface ReadComponentProps extends ComponentProps {}
export const TagReadComponent = React.memo(
  (props: React.PropsWithChildren<ReadComponentProps>) => {
    return (
      <DataCollection
        datasource={props.tags}
        isLoading={props.isLoading}
        makeHeader={makeHeader}
        makeBody={makeBody}
        makeEditLink={(item) => `/tags/${(item as Tag).id}`}
        makeStyle={(item) => ({
          borderLeft: `${(item as Tag).color} 5px solid`,
        })}
        makeBodyStyle={makeBodyStyle}
      />
    );
  }
);

interface SelectComponentProps extends ComponentProps {
  onSelect: (item: Tag) => void;
}
export const TagSelectComponent = React.memo(
  (props: React.PropsWithChildren<SelectComponentProps>) => {
    return (
      <DataCollection
        datasource={props.tags}
        isLoading={props.isLoading}
        makeHeader={makeHeader}
        makeStyle={(item) => ({
          borderLeft: `${(item as Tag).color} 5px solid`,
        })}
        onSelect={(item) => {
          props.onSelect(item[0] as Tag);
        }}
      />
    );
  }
);
