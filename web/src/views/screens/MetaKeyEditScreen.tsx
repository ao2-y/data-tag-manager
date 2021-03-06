import React, { useEffect } from 'react';
import { MetaKeyEditState } from '../../interfaces/controllers/states';
import { MetaKeyEditActions } from '../container/MetaKeyEditContainer';
import Loading from '../components/Loading';
import { MetaKey } from '../../domain/model';
import { MetaKeyForm } from '../components/MetaKeyForm';

interface OwnProps {
  match: { params: { id: string } };
}
type MetaKeyFormProps = OwnProps & MetaKeyEditState & MetaKeyEditActions;

const MetaKeyFormScreen: React.FC<MetaKeyFormProps> = (props) => {
  useEffect(() => {
    props.fetch(Number(props.match.params.id));
  }, [true]);

  if (props.metaKey) {
    const updateItem = (value: MetaKey) => {
      const target = JSON.parse(JSON.stringify(value));
      console.log(target);
    };
    return (
      <MetaKeyForm
        metaKey={props.metaKey}
        onSave={(v) => {
          updateItem(v);
        }}></MetaKeyForm>
    );
  } else {
    return <Loading />;
  }
};

export default MetaKeyFormScreen;
