import React, { useEffect } from 'react';
import { TagNewState } from '../../interfaces/controllers/states';
import { TagNewActions } from '../container/TagNewContainer';
import { TagForm } from '../components/TagForm';
import { Tag } from '../../domain/model';
import Loading from '../components/Loading';

interface OwnProps {
  match: { params: { id: string } };
}
type TagNewProps = OwnProps & TagNewState & TagNewActions;

const TagFormScreen: React.FC<TagNewProps> = (props) => {
  useEffect(() => {
    props.loadInitData();
  }, [true]);
  console.log(props.tag);
  if (props.tag) {
    const updateItem = (value: Tag) => {
      const target = JSON.parse(JSON.stringify(value));
      console.log(target);
    };
    return (
      <TagForm
        tag={props.tag}
        onSave={(v) => {
          updateItem(v);
        }}></TagForm>
    );
  } else {
    return <Loading />;
  }
};

export default TagFormScreen;
