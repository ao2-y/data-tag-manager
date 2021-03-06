import { Item } from '../model';

export interface ItemRepository  {

  find(param?: {}): Promise<Item[]>;
  fetchById(id: number): Promise<Item | undefined>;

}