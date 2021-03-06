import React, { useState } from 'react';
import {
  CLink,
  CCard,
  CCardBody,
  CRow,
  CCol,
  CCardHeader,
  CCollapse,
} from '@coreui/react';
import CIcon from '@coreui/icons-react';
import Loading from '../components/Loading';

type DataObject = { id: number };

// FIXME: Genericsがうまく動かないのでanyで一旦実装する
// eslint-disable-next-line @typescript-eslint/no-explicit-any
interface OwnProps<T extends {}> {
  datasource: T[];
  isLoading?: boolean;
  makeHeader: (item: T) => string | React.ReactElement;
  makeBody?: (item: T) => React.ReactNode;
  makeEditLink?: (item: T) => string;
  mutiple?: boolean;
  onSelect?: (item: T[]) => void;
  makeStyle?: (item: T) => React.CSSProperties;
  makeBodyStyle?: (item: T) => React.CSSProperties;
}

type ComponentProps = OwnProps<{}>;

const findItem = (item: DataObject, source: DataObject[]) => {
  return source.find((v) => v.id === item.id);
};

const Component: React.FC<ComponentProps> = (props) => {
  const [collapsed, setCollapsed] = useState<{ [key: number]: boolean }>({});
  const [keyPrefix] = useState<string>(Math.random().toString(32).substring(2));
  const [selectes, setSelects] = useState<DataObject[]>([]);

  const multipleSelect = (
    item: DataObject,
    source: DataObject[],
    callback: (items: DataObject[]) => void
  ) => {
    const findIndex = source.indexOf(item);
    let results: DataObject[] = [];
    if (findIndex === -1) {
      results = source.concat(item);
    } else {
      results = source.filter((v) => v !== item);
    }
    callback(results);
    if (props.onSelect) {
      props.onSelect(results);
    }
  };
  const singleSelect = (
    item: DataObject,
    source: DataObject[],
    callback: (items: DataObject[]) => void
  ) => {
    const results: DataObject[] = [item];
    callback(results);
    if (props.onSelect) {
      props.onSelect(results);
    }
  };
  return (
    <CRow>
      {props.isLoading && <Loading />}
      {!props.isLoading &&
        props.datasource.map((object) => {
          const item = object as DataObject;
          let cardClassName = '';
          props.onSelect && (cardClassName += `selectable`);
          props.onSelect &&
            (cardClassName += findItem(item, selectes) ? ` text-white` : ``);
          return (
            <CCol key={`${keyPrefix}${item.id}`} md={12} lg={6} xl={4}>
              <CCard
                style={props.makeStyle ? props.makeStyle(item) : undefined}
                className={cardClassName}
                color={findItem(item, selectes) ? 'info' : ''}
                onClick={() => {
                  props.onSelect &&
                    (props.mutiple === true
                      ? multipleSelect(item, selectes, setSelects)
                      : singleSelect(item, selectes, setSelects));
                }}>
                <CCardHeader>
                  {props.makeHeader(item)}
                  <div
                    className="card-header-actions"
                    onClick={(e) => {
                      e.stopPropagation();
                    }}>
                    {props.makeEditLink && (
                      <CLink
                        className="card-header-action"
                        to={props.makeEditLink(item)}>
                        <CIcon name="cil-pencil" />
                      </CLink>
                    )}
                    {props.makeBody && (
                      <CLink
                        className="card-header-action"
                        onClick={() => {
                          collapsed[item.id] = !collapsed[item.id];
                          setCollapsed(JSON.parse(JSON.stringify(collapsed)));
                        }}>
                        <CIcon
                          name={
                            collapsed[item.id]
                              ? 'cil-chevron-bottom'
                              : 'cil-chevron-top'
                          }
                        />
                      </CLink>
                    )}
                  </div>
                </CCardHeader>
                <CCollapse show={collapsed[item.id]}>
                  <CCardBody
                    className="py-2"
                    style={
                      props.makeBodyStyle
                        ? props.makeBodyStyle(item)
                        : undefined
                    }>
                    {props.makeBody && props.makeBody(item)}
                  </CCardBody>
                </CCollapse>
              </CCard>
            </CCol>
          );
        })}
    </CRow>
  );
};

export default React.memo(Component);
