import React, { useEffect } from 'react';
import * as H from 'history';
import { TagsState } from '../../interfaces/controllers/states';
import { TagsActions } from '../container/TagsContainer';
import { TagReadComponent } from '../components/TagCollection';
import TagSearchForm from '../components/TagSearchForm';
import { CButton } from '@coreui/react';
import CIcon from '@coreui/icons-react';

interface OwnProps {
  history: H.History;
}
type TagsProps = OwnProps & TagsState & TagsActions;

const TagsScreen: React.FC<TagsProps> = (props) => {
  useEffect(() => {
    props.search();
  }, []);

  return (
    <div className="">
      <div className="d-flex mb-4">
        <div className="flex-grow-1 ">
          <TagSearchForm onChange={() => {}}></TagSearchForm>
        </div>
        <div className="d-flex align-items-end flex-column">
          <CButton
            type="button"
            variant="outline"
            size="sm"
            className="ml-2 mt-auto"
            color="primary"
            onClick={() => {
              props.history.push('/tags/new');
            }}>
            <CIcon name="cil-plus" /> Add Tag
          </CButton>
        </div>
      </div>
      <TagReadComponent tags={props.tags}></TagReadComponent>
    </div>
  );
};

export default TagsScreen;
