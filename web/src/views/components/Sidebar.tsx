import React from 'react';
import {
  CCreateElement,
  CSidebar,
  CSidebarBrand,
  CSidebarNav,
  CSidebarNavDivider,
  CSidebarNavTitle,
  CSidebarMinimizer,
  CSidebarNavDropdown,
  CSidebarNavItem,
} from '@coreui/react';

import CIcon from '@coreui/icons-react';

import { LayoutState } from '../../interfaces/controllers/states/LayoutState';
import { LayoutActions } from '../container/LayoutContainer';

interface OwnProps {}

type LayoutProps = OwnProps & LayoutState & LayoutActions;

const navigation = [
  {
    _tag: 'CSidebarNavTitle',
    _children: ['Generals'],
  },
  {
    _tag: 'CSidebarNavItem',
    name: 'Items',
    to: '/items',
    exact: false,
    icon: 'cil-list',
  },
  {
    _tag: 'CSidebarNavTitle',
    _children: ['Settings'],
  },
  {
    _tag: 'CSidebarNavItem',
    name: 'ItemTemplate',
    to: '/item-templates',
    exact: false,
    icon: 'cil-spreadsheet',
  },
  {
    _tag: 'CSidebarNavItem',
    name: 'MetaKey',
    to: '/meta-keys',
    exact: false,
    icon: 'cil-info',
  },
  {
    _tag: 'CSidebarNavItem',
    name: 'Tag',
    to: '/tags',
    exact: false,
    icon: 'cil-tags',
  },
];

const Sidebar: React.FC<LayoutProps> = (props) => {
  return (
    <CSidebar
      show={props.sidebarShow}
      onShowChange={(val: boolean | 'responsive') => props.showSidebar(val)}>
      <CSidebarBrand className="d-md-down-none" to="/">
        <CIcon
          className="c-sidebar-brand-full"
          name="logo-negative"
          height={35}
        />
        <CIcon
          className="c-sidebar-brand-minimized"
          name="sygnet"
          height={35}
        />
      </CSidebarBrand>
      <CSidebarNav>
        <CCreateElement
          items={navigation}
          components={{
            CSidebarNavDivider,
            CSidebarNavDropdown,
            CSidebarNavItem,
            CSidebarNavTitle,
          }}
        />
      </CSidebarNav>
      <CSidebarMinimizer className="c-d-md-down-none" />
    </CSidebar>
  );
};

export default React.memo(Sidebar);
