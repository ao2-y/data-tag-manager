import React from 'react';
import { RouteComponentProps, RouteProps } from 'react-router-dom';
interface IRouteProps extends RouteProps {
  name: string;
  component:
    | React.ComponentType<RouteComponentProps<{}>>
    | React.ComponentType<{}>;
}
const routes = [
  { path: '/', exact: true, name: 'Home' },
  // ItemTemplate
  {
    path: '/item-templates/new',
    name: 'New',
    component: React.lazy(
      () => import('../../views/container/ItemTemplateNewContainer')
    ),
    exact: true,
  },
  {
    path: '/item-templates/:id',
    name: 'Edit',
    component: React.lazy(
      () => import('../../views/container/ItemTemplateEditContainer')
    ),
    exact: true,
  },
  {
    path: '/item-templates',
    name: 'Item Templates',
    component: React.lazy(
      () => import('../../views/container/ItemTemplatesContainer')
    ),
    exact: true,
  },
  // Tag
  {
    path: '/tags/new',
    name: 'New',
    component: React.lazy(
      () => import('../../views/container/TagNewContainer')
    ),
    exact: true,
  },
  {
    path: '/tags/:id',
    name: 'Edit',
    component: React.lazy(
      () => import('../../views/container/TagEditContainer')
    ),
    exact: true,
  },
  {
    path: '/tags',
    name: 'Tags',
    component: React.lazy(() => import('../../views/container/TagsContainer')),
    exact: true,
  },
  // Item
  {
    path: '/items/new',
    name: 'New',
    component: React.lazy(
      () => import('../../views/container/ItemNewContainer')
    ),
    exact: true,
  },
  {
    path: '/items/:id',
    name: 'Edit',
    component: React.lazy(
      () => import('../../views/container/ItemEditContainer')
    ),
    exact: true,
  },
  {
    path: '/items',
    name: 'Item',
    component: React.lazy(() => import('../../views/container/ItemsContainer')),
    exact: true,
  },
  // MetaKey
  {
    path: '/meta-keys/new',
    name: 'New',
    component: React.lazy(
      () => import('../../views/container/MetaKeyNewContainer')
    ),
    exact: true,
  },
  {
    path: '/meta-keys/:id',
    name: 'Edit',
    component: React.lazy(
      () => import('../../views/container/MetaKeyEditContainer')
    ),
    exact: true,
  },
  {
    path: '/meta-keys',
    name: 'MetaKeys',
    component: React.lazy(
      () => import('../../views/container/MetaKeysContainer')
    ),
    exact: true,
  },
] as IRouteProps[];

export default routes;
