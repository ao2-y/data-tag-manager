import React, { useEffect } from 'react';
import { MetaKeyNewState } from '../../interfaces/controllers/states';
import { MetaKeyNewActions } from '../container/MetaKeyNewContainer';
import { MetaKeyForm } from '../components/MetaKeyForm';
import { MetaKey } from '../../domain/model';
import Loading from '../components/Loading';

interface OwnProps {
  match: { params: { id: string } };
}
type MetaKeyNewProps = OwnProps & MetaKeyNewState & MetaKeyNewActions;

const MetaKeyFormScreen: React.FC<MetaKeyNewProps> = (props) => {
  useEffect(() => {
    props.loadInitData();
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
