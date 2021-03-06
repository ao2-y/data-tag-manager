import React, { useEffect } from 'react';
import { ItemNewState } from '../../interfaces/controllers/states';
import { ItemNewActions } from '../container/ItemNewContainer';
import Loading from '../components/Loading';
import { Item } from '../../domain/model';
import { ItemForm } from '../components/ItemForm';

interface OwnProps {}
type ItemNewProps = OwnProps & ItemNewState & ItemNewActions;

const ItemNewScreen: React.FC<ItemNewProps> = (props) => {
  useEffect(() => {
    props.loadInitItem(props.template);
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

export default ItemNewScreen;
