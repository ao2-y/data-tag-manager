import React, { useEffect } from 'react';
import { ItemEditState } from '../../interfaces/controllers/states';
import { ItemEditActions } from '../container/ItemEditContainer';
import Loading from '../components/Loading';
import { Item } from '../../domain/model';
import { ItemForm } from '../components/ItemForm';

interface OwnProps {
  match: { params: { id: string } };
}
type ItemEditProps = OwnProps & ItemEditState & ItemEditActions;

const ItemEditScreen: React.FC<ItemEditProps> = (props) => {
  useEffect(() => {
    props.fetch(Number(props.match.params.id));
  }, [true]);

  if (props.item) {
    const updateItem = (value: Item) => {
      const target = JSON.parse(JSON.stringify(value));
      console.log(target);
    };
    return (
      <ItemForm
        item={props.item}
        onSave={(v) => {
          updateItem(v);
        }}></ItemForm>
    );
  } else {
    return <Loading />;
  }
};

export default ItemEditScreen;
