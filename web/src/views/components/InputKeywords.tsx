import React, { useEffect, useState, useRef } from 'react';

import { container } from 'tsyringe';
import { TagUsecase } from '../../usecase/TagUsecase';
import { MetaKeyUsecase } from '../../usecase/MetaKeyUsecase';
import { ItemUsecase } from '../../usecase/ItemUsecase';
import Loading from './Loading';
import CIcon from '@coreui/icons-react';

interface OwnProps {
  placeholder?: string;
}
type ComponentProps = OwnProps;

interface IDatasource {
  metaKeys: IAutoCompleteItem[];
  tags: IAutoCompleteItem[];
  descriptions: IAutoCompleteItem[];
  names: IAutoCompleteItem[];
}
interface IAutoCompleteItem {
  text: string;
  value: string;
  forKey: string;
}
interface IAutoCompleteCategoryItem extends IAutoCompleteItem {
  makeAutoComplete: (value: string, datasouce: IDatasource, ignoreKeys: string[]) => IAutoCompleteItem[];
}

const filter = (value: string, datasource: IAutoCompleteItem[], ignoreKeys: string[]) => {
  let ds = datasource;
  if (value !== '') {
    ds = datasource.filter((v) => v.value.indexOf(value) > -1);
  }
  return ds.filter((v) => ignoreKeys.indexOf(v.forKey) === -1);
};

enum CategoryName {
  Tag = 'Tag',
  MetaKey = 'MetaKey',
  Description = 'Description',
  Name = 'Name',
}

const CATEGORY_ITEMS: IAutoCompleteCategoryItem[] = [
  {
    text: CategoryName.Tag,
    value: `${CategoryName.Tag}:`,
    forKey: 'category-tag',
    makeAutoComplete: (value, datasource, ignoreKeys) => filter(value, datasource.tags, ignoreKeys),
  },
  {
    text: CategoryName.MetaKey,
    value: `${CategoryName.MetaKey}:`,
    forKey: 'category-metakey',
    makeAutoComplete: (value, datasource, ignoreKeys) => filter(value, datasource.metaKeys, ignoreKeys),
  },
  {
    text: CategoryName.Description,
    value: `${CategoryName.Description}:`,
    forKey: 'category-description',
    makeAutoComplete: (value, datasource, ignoreKeys) => filter(value, datasource.descriptions, ignoreKeys),
  },
  {
    text: CategoryName.Name,
    value: `${CategoryName.Name}:`,
    forKey: 'category-name',
    makeAutoComplete: (value, datasource, ignoreKeys) => filter(value, datasource.names, ignoreKeys),
  },
];

const tagUsecase = container.resolve(TagUsecase);
const metaKeysUsecase = container.resolve(MetaKeyUsecase);
const itemUsecase = container.resolve(ItemUsecase);

const fetchDatasource = async () => {
  const values = await Promise.all([tagUsecase.find(), metaKeysUsecase.find(), itemUsecase.find()]);
  return {
    tags: values[0].map((m) => ({
      text: m.name,
      value: `${CategoryName.Tag}:${m.name}`,
      forKey: `item-tag-${m.name}`,
    })),
    metaKeys: values[1].map((m) => ({
      text: m.name,
      value: `${CategoryName.MetaKey}:${m.name}`,
      forKey: `item-metakey-${m.name}`,
    })),
    names: values[2]
      .filter((v, i, self) => self.findIndex((e) => e.name === v.name) === i)
      .map((m) => ({ text: m.name, value: `${CategoryName.Name}:${m.name}`, forKey: `item-name-${m.name}` })),
    descriptions: values[2]
      .filter((v, i, self) => self.findIndex((e) => e.description === v.description) === i)
      .filter((v) => v.description.length > 0)
      .map((m) => ({
        text: m.description,
        value: `${CategoryName.Description}:${m.description}`,
        forKey: `item-description-${m.description}`,
      })),
  } as IDatasource;
};

const makeAutocomplete = (
  value: string,
  inputeds: IAutoCompleteItem[],
  datasource?: IDatasource
): IAutoCompleteItem[] => {
  const [category, val] = value.split(':');
  if (val == null) {
    if (category.length === 0) {
      return CATEGORY_ITEMS;
    }
    return CATEGORY_ITEMS.filter((v) => v.text.indexOf(category) > -1);
  }

  // 該当するカテゴリを検索してオートコンプリートを生成
  const found = CATEGORY_ITEMS.find((v) => value.startsWith(v.text));
  if (found && datasource) {
    return found.makeAutoComplete(
      val,
      datasource,
      inputeds.map((v) => v.forKey)
    );
  }

  return [];
};

const Component: React.FC<ComponentProps> = (props) => {
  const [, setInputValue] = useState<string>('');
  const [values, setValues] = useState<IAutoCompleteItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [datasource, setDatasource] = useState<IDatasource | undefined>();
  const [isInputFocus, setIsInputFocus] = useState<boolean>(false);
  const inputElement = useRef<HTMLInputElement>(null);

  const onInputTextChange = () => {
    setInputValue(inputElement.current?.value || '');
  };

  const onAutocompleteSelect = (e: React.MouseEvent<HTMLDivElement>, object: IAutoCompleteItem) => {
    if ((object as IAutoCompleteCategoryItem).makeAutoComplete) {
      if (inputElement.current) {
        inputElement.current.value = object.value;
        onInputTextChange();
      }
    } else {
      if (inputElement.current) {
        inputElement.current.value = '';
        onInputTextChange();
      }
      setValues(values.concat([object]));
    }
    setTimeout(() => inputElement.current?.focus());
  };

  const onInputKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      if (!inputElement.current?.value.length) {
        return;
      }
      const value = inputElement.current?.value || '';
      const separatorIndex = value.indexOf(':');
      if (separatorIndex === -1) {
        const key = `manual-item-${value}`;
        // 重複してなければ追加。
        const found = values.find((v) => v.forKey === key);
        if (!found) {
          inputElement.current.value = '';
          onInputTextChange();
          setValues(values.concat([{ text: value, value, forKey: key }]));
        }
      } else if (separatorIndex === value.length - 1) {
        // 最後尾にあるので無視する
      } else {
        // オートコンプリートの中で該当するものがないか検索
        const found = makeAutocomplete(value, values, datasource).find((v) => v.value === value);
        if (found) {
          inputElement.current.value = '';
          onInputTextChange();
          setValues(values.concat([found]));
        } else {
          const key = `manual-item-${value}`;
          // 重複してなければ追加。
          const found = values.find((v) => v.forKey === key);
          if (!found) {
            inputElement.current.value = '';
            onInputTextChange();
            setValues(values.concat([{ text: value, value, forKey: key }]));
          }
        }
      }
      return;
    }
  };

  useEffect(() => {
    fetchDatasource()
      .then(setDatasource)
      .then(() => setIsLoading(false));
  }, [isLoading]);

  const autocompletes = makeAutocomplete(inputElement.current?.value || '', values, datasource);
  const message = 'Not found autocomplete.';
  return (
    <div>
      <style>{`
      .input-keywords { }
      .input-keywords > .input-container { border-radius: 0.25rem; background: rgb(255, 255, 255); padding: 0.25rem 0.75rem 0 0.75rem; position: relative; border: 1px solid rgb(216, 219, 224); font-size: 1em; line-height: 0.7; cursor: text;}
      .input-keywords > .input-container.focused { border-color: #958bef; outline: 0; box-shadow: 0 0 0 0.2rem rgba(50, 31, 219, 0.25);}
      .input-keywords > .input-container > * {display: inline-block; position: relative;}
      .input-keywords > .input-container > .keyword-item {display: inline-block; margin: 0 6px 2px 0; padding: 6px 8px; border: 1px solid #D1D1D1; border-radius: 2px; background: #F1F1F1; position: relative; }
      .input-keywords > .input-container > .keyword-item > * { white-space: nowrap; }
      .input-keywords > .input-container > .keyword-item > .close {position: absolute; right: 3px; top: 4px; color: #AAA}
      .input-keywords > .input-container > .input-wrapper { padding: 3px 2px; margin: 0 0 2px 0; max-width: 100% }
      .input-keywords > .input-container > .input-wrapper > input { max-width: 100%; margin: 0 0 0 0; padding: 0; outline: none; border: 0; }
      .autocomplete-row{line-height: 1.5em; padding: 2px 20px;}
      .autocomplete-row.clickable{cursor: pointer;}
      .autocomplete-row.clickable:hover{background: rgb(55,55,55); color: rgb(255,153,0)}
      `}</style>
      <div className="input-keywords" onClick={() => inputElement.current?.focus()}>
        <div className={`input-container ${isInputFocus ? 'focused' : ''}`}>
          {values.map((val, index) => (
            <button
              key={val.forKey}
              type="button"
              className="keyword-item"
              onClick={() => setValues(values.splice(index, 1) && [...values])}>
              <span className="mr-3">{val.value}</span>
              <CIcon name="cil-x" size="sm" className="close" />
            </button>
          ))}
          <div className="input-wrapper">
            <input
              ref={inputElement}
              placeholder={props.placeholder}
              onKeyPress={onInputKeyPress}
              onChange={onInputTextChange}
              onFocus={() => setIsInputFocus(true)}
              onBlur={() => setIsInputFocus(false)}
            />
          </div>
        </div>
        <div
          style={{
            marginTop: '0.25rem',
            borderRadius: '0.25rem',
            borderColor: '#d8dbe0',
            color: '#fff',
            background: '#444',
            padding: '5px 0',
            display: document.activeElement === inputElement.current ? undefined : 'none',
          }}>
          <div>
            {isLoading && (
              <div className="autocomplete-row">
                <Loading />
              </div>
            )}
            {!isLoading && message && <div className="autocomplete-row">{message}</div>}
            {!isLoading &&
              autocompletes.length > 0 &&
              autocompletes.map((object) => (
                <div
                  className="autocomplete-row clickable"
                  key={object.forKey}
                  onMouseDown={(e) => onAutocompleteSelect(e, object)}>
                  {object.text}
                </div>
              ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Component;
