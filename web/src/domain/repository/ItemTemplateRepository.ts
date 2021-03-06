import { ItemTemplate } from '../model';

export interface ItemTemplateRepository { 

  find(param?: {}): Promise<ItemTemplate[]>;
  fetchById(id: number): Promise<ItemTemplate | undefined>;

}