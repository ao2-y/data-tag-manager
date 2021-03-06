import React from 'react';
import { container } from 'tsyringe';
import ReactTags, { Tag as ReactTagAutoCompleteTag } from 'react-tag-autocomplete';
import { Tag } from '../../domain/model';
import { TagUsecase } from '../../usecase/TagUsecase';

interface Props {
  value: Tag[];
  onChange: (value: Tag[]) => void;
}

interface State {
  tags: ReactTagAutoCompleteTag[];
  suggestions: ReactTagAutoCompleteTag[];
}

const usecase = container.resolve(TagUsecase)

export default class InputTag extends React.Component<Props, State> {
  private reactTags: React.RefObject<ReactTags>;
  constructor (props: Props) {
    super(props)
    const tags = props.value.map(item => ({ id: item.name, name: item.name }));
    this.state = { tags, suggestions: [{id: 1, name: 'Banner'}] };
    this.reactTags = React.createRef()
  }

  onDelete (i: number) {
    const tags = this.state.tags.slice(0);
    tags.splice(i, 1);
    this.setState({ tags });
  }

  onAddition (tag: ReactTagAutoCompleteTag) {
    const tags = ([] as ReactTagAutoCompleteTag[]).concat(this.state.tags, tag);
    this.setState({ tags });
    this.props.onChange(tags as Tag[]);
  }

  onInput(query: string) {
    if (query.length < 3) { return; } 
    usecase.find({keyword: query, excludes: this.state.tags.map(v => v.name)})
    .then(suggestions => { this.setState({ suggestions }) })
  }

  render () {
    return (
      <ReactTags
        ref={this.reactTags}
        tags={this.state.tags}
        allowNew={true}
        placeholderText="Please entry tags..."
        suggestions={this.state.suggestions}
        onDelete={this.onDelete.bind(this)}
        onInput={this.onInput.bind(this)}
        onAddition={this.onAddition.bind(this)} />
    )
  }
}


// interface OwnProps { 
//   tags: []
// };
// type InputMetaProps = OwnProps;

// const ItemTemplateFormScreen: React.FC<InputMetaProps> = (props) => {
//   return (
//     <ReactTags
//     ref={this.reactTags}
//     tags={props.tags}
//     suggestions={[
//       { id: 3, name: "Bananas" },
//       { id: 4, name: "Mangos" },
//       { id: 5, name: "Lemons" },
//       { id: 6, name: "Apricots" }
//     ]}
//     onDelete={this.onDelete.bind(this)}
//     onAddition={this.onAddition.bind(this)} />
//   )
// }

// export default ItemTemplateFormScreen;