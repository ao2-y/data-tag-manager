import React from 'react'
import TheFooter from '../components/Footer';
// import TheSidebar from '../components/Sidebar';
import TheContent from '../components/Content';
// import TheHeader from '../components/Header';
import { LayoutState } from '../../interfaces/controllers/states/LayoutState';
import { LayoutActions } from '../container/LayoutContainer';
import TheHeader from '../container/HeaderContainer';
import TheSidebar from '../container/SidebarContainer';
interface OwnProps { }

type LayoutProps = OwnProps & LayoutState & LayoutActions;

const Layout: React.FC<LayoutProps> = (props: LayoutProps) => { 
  props;
  return (
    <div className="c-app c-default-layout">
      <TheSidebar />
      <div className="c-wrapper">
        <TheHeader />
        <div className="c-body">
          <TheContent />
        </div>
        <TheFooter />
      </div>
    </div>
  )
}

export default Layout
