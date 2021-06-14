import React, { useEffect, useState } from 'react';
import { MetaKeyEditState } from '../../interfaces/controllers/states';
import { MetaKeyEditActions } from '../container/MetaKeyEditContainer';
import Loading from '../components/Loading';
import { SubmitButton } from '../components/buttons';
import { MetaKeyForm } from '../components/MetaKeyForm';

interface OwnProps {
  match: { params: { id: string } };
}
type MetaKeyFormProps = OwnProps & MetaKeyEditState & MetaKeyEditActions;

const MetaKeyFormScreen: React.FC<MetaKeyFormProps> = (props) => {
  const [name, setName] = useState(props.metaKey?.name);
  useEffect(() => props.fetch(props.match.params.id), [props.match.params.id]);

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
      {!props.isLoading && props.metaKey && (
        <SubmitButton onClick={() => name && props.register({ id: props.match.params.id, name })} />
      )}
    </div>
  );
};

export default MetaKeyFormScreen;
