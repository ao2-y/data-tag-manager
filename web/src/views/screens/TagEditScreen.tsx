import React, { useEffect } from 'react';
import { TagEditState } from '../../interfaces/controllers/states';
import { TagEditActions } from '../container/TagEditContainer';
import Loading from '../components/Loading';
import { Tag } from '../../domain/model';
import { TagForm } from '../components/TagForm';

interface OwnProps {
  match: { params: { id: string } };
}
type TagEditProps = OwnProps & TagEditState & TagEditActions;

const TagEditScreen: React.FC<TagEditProps> = (props) => {
  useEffect(() => {
    props.fetch(Number(props.match.params.id));
  }, [true]);

  if (props.tag) {
    const updateTag = (value: Tag) => {
      const target = JSON.parse(JSON.stringify(value));
      console.log(target);
    };
    return (
      <TagForm
        tag={props.tag}
        onSave={(v) => {
          updateTag(v);
        }}></TagForm>
    );
  } else {
    return <Loading />;
  }
};

export default TagEditScreen;
