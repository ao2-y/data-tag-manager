import React from 'react';
import { MetaKey } from '../../domain/model';
import DataCollection from './DataCollection';
interface ComponentProps {
  metaKeys: MetaKey[];
  isLoading?: boolean;
}

const makeHeader = (object: {}) => {
  const item = object as MetaKey;
  return <span>{item.name}</span>;
};

interface ReadComponentProps extends ComponentProps {}
export const MetaKeyReadComponent = React.memo(
  (props: React.PropsWithChildren<ReadComponentProps>) => {
    return (
      <DataCollection
        datasource={props.metaKeys}
        isLoading={props.isLoading}
        makeHeader={makeHeader}
        makeEditLink={(item) =>
          `/meta-keys/${(item as MetaKey).id}`
        }></DataCollection>
    );
  }
);

interface SelectComponentProps extends ComponentProps {
  onSelect: (item: MetaKey) => void;
}
export const MetaKeySelectComponent = React.memo(
  (props: React.PropsWithChildren<SelectComponentProps>) => {
    return (
      <DataCollection
        datasource={props.metaKeys}
        isLoading={props.isLoading}
        makeHeader={makeHeader}
        onSelect={(item) => {
          props.onSelect(item[0] as MetaKey);
        }}></DataCollection>
    );
  }
);
