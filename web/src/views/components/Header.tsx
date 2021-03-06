import React from 'react';
import {
  CHeader,
  CToggler,
  CHeaderBrand,
  CHeaderNav,
  CHeaderNavItem,
  CHeaderNavLink,
  CSubheader,
  CBreadcrumbRouter,
} from '@coreui/react';
import CIcon from '@coreui/icons-react';

// routes config
import routes from '../../infrastructure/router/routes';
import { LayoutState } from '../../interfaces/controllers/states/LayoutState';
import { LayoutActions } from '../container/LayoutContainer';

interface OwnProps {}

type LayoutProps = OwnProps & LayoutState & LayoutActions;

const Header: React.FC<LayoutProps> = (props) => {
  const toggleSidebar = () => {
    const val = [true, 'responsive'].includes(props.sidebarShow)
      ? false
      : 'responsive';
    props.showSidebar(val);
  };

  const toggleSidebarMobile = () => {
    const val = [false, 'responsive'].includes(props.sidebarShow)
      ? true
      : 'responsive';
    props.showSidebar(val);
  };

  return (
    <CHeader withSubheader>
      <CToggler
        inHeader
        className="ml-md-3 d-lg-none"
        onClick={toggleSidebarMobile}
      />
      <CToggler
        inHeader
        className="ml-3 d-md-down-none"
        onClick={toggleSidebar}
      />
      <CHeaderBrand className="mx-auto d-lg-none" to="/">
        <CIcon name="logo" height="48" alt="Logo" />
      </CHeaderBrand>
      <CHeaderNav className="d-md-down-none mr-auto">
        <CHeaderNavItem className="px-3">
          <CHeaderNavLink to="/items">Item</CHeaderNavLink>
        </CHeaderNavItem>
        <CHeaderNavItem className="px-3">
          <CHeaderNavLink to="/item-templates">Item Template</CHeaderNavLink>
        </CHeaderNavItem>
        <CHeaderNavItem className="px-3">
          <CHeaderNavLink to="/tags">Tag</CHeaderNavLink>
        </CHeaderNavItem>
        <CHeaderNavItem className="px-3">
          <CHeaderNavLink to="/meta-keys">MetaKey</CHeaderNavLink>
        </CHeaderNavItem>
      </CHeaderNav>

      <CSubheader className="px-3 justify-content-between">
        <CBreadcrumbRouter
          className="border-0 c-subheader-nav m-0 px-0 px-md-3"
          routes={routes}
        />
        <div className="d-md-down-none mfe-2 c-subheader-nav">
          {/* <CLink className="c-subheader-nav-link" href="#">
            <CIcon name="cil-speech" alt="Settings" />
          </CLink> */}
          {/* <CLink className="c-subheader-nav-link" href="#">
            <CIcon name="cil-settings" alt="Settings" />
            &nbsp;Settings
          </CLink> */}
        </div>
      </CSubheader>
    </CHeader>
  );
};

export default React.memo(Header);
