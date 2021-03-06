import { ItemMeta } from './ItemMeta';
import { ItemTag } from './ItemTag';

export interface Item {
  id: number;
  name: string;
  description: string;
  metaDatas: ItemMeta[];
  tags: ItemTag[];
}
