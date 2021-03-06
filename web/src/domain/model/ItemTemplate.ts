import { ItemTemplateMetaKey } from './ItemTemplateMetaKey';

export interface ItemTemplate {
  id: number;
  name: string;
  metas: ItemTemplateMetaKey[];
}
