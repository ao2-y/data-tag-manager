import React, { useEffect } from 'react';
import { ItemTemplateEditState } from '../../interfaces/controllers/states';
import { ItemTemplateEditActions } from '../container/ItemTemplateEditContainer';
import { ItemTemplateForm } from '../components/ItemTemplateForm';
import { ItemTemplate } from '../../domain/model';
import Loading from '../components/Loading';

interface OwnProps {
  match: { params: { id: string } };
}
type ItemTemplateEditProps = OwnProps & ItemTemplateEditState & ItemTemplateEditActions;

const ItemTemplateEditScreen: React.FC<ItemTemplateEditProps> = (props) => {
  useEffect(() => {
    props.fetch(Number(props.match?.params.id));
  }, [true]);
  if (props.itemTemplate) {
    const updateItem = (value: ItemTemplate) => {
      const target = JSON.parse(JSON.stringify(value));
      console.log(target);
    };
    return (
      <ItemTemplateForm
        itemTemplate={props.itemTemplate}
        onSave={(v) => {
          updateItem(v);
        }}></ItemTemplateForm>
    );
  } else {
    return <Loading />;
  }
};

export default ItemTemplateEditScreen;
