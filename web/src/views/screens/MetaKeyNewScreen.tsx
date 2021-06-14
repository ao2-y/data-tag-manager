import React, { useEffect, useState } from 'react';
import { MetaKeyNewState } from '../../interfaces/controllers/states';
import { MetaKeyNewActions } from '../container/MetaKeyNewContainer';
import Loading from '../components/Loading';
import { SubmitButton } from '../components/buttons';
import { MetaKeyForm } from '../components/MetaKeyForm';

interface OwnProps {}
type MetaKeyNewProps = OwnProps & MetaKeyNewState & MetaKeyNewActions;

const MetaKeyFormScreen: React.FC<MetaKeyNewProps> = (props) => {
  const [name, setName] = useState(props.metaKey?.name);
  useEffect(() => props.loadInitData(), [true]);

  const mapper: { [key: string]: (val: string) => void } = { name: setName };
  return (
    <div>
      {props.isLoading && <Loading />}
      {!props.isLoading && props.metaKey && (
        <MetaKeyForm
          name={props.metaKey.name}
          onChange={(v) => Object.keys(v).forEach((k) => mapper[k](v[k]))}></MetaKeyForm>
      )}
      {!props.isLoading && !props.metaKey && <h1>Not found meta.</h1>}
      {!props.isLoading && props.metaKey && <SubmitButton onClick={() => name && props.register({ name })} />}
    </div>
  );
};

export default MetaKeyFormScreen;
