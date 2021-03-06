import React, { useEffect } from 'react';
import { ItemTemplateNewState } from '../../interfaces/controllers/states';
import { ItemTemplateNewActions } from '../container/ItemTemplateNewContainer';
import { ItemTemplateForm } from '../components/ItemTemplateForm';
import { ItemTemplate } from '../../domain/model';
import Loading from '../components/Loading';

interface OwnProps {
  match: { params: { id: string } };
}
type ItemTemplateNewProps = OwnProps &
  ItemTemplateNewState &
  ItemTemplateNewActions;

const ItemTemplateFormScreen: React.FC<ItemTemplateNewProps> = (props) => {
  useEffect(() => {
    props.loadInitData();
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

export default ItemTemplateFormScreen;
