import { ItemRepository } from '../../domain/repository';
import { Item } from '../../domain/model';
import { ITEM_DATA } from './mock';

export class ItemDataStore implements ItemRepository {
  find(param?: {}): Promise<Item[]> {
    console.log(`Item - Find`, JSON.stringify(param));
    return Promise.resolve(ITEM_DATA);
  }
  fetchById(id: number): Promise<Item| undefined> {
    console.log(`Item - fetchById`, id);
    return Promise.resolve(ITEM_DATA.find(v => v.id === id));
  }
}