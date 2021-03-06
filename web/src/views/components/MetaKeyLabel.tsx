import React, { useEffect, useState } from 'react';
import { container } from 'tsyringe';
import { MetaKey } from '../../domain/model';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';

interface OwnProps {
  metaKeyId: number;
  metaKey?: MetaKey;
}

type ComponentProps = OwnProps;

const usecase = container.resolve(MetaKeyUsecase)

const Component: React.FC<ComponentProps> = (props) => {
  const [metaKey, setMetaKeyValue] = useState<MetaKey>();
  useEffect(() => { usecase.fetch(props.metaKeyId).then(setMetaKeyValue).then(console.log); }, []);
  return (
    <span>{metaKey?.name || 'Loading...'}</span>
  );
};

export default Component;
