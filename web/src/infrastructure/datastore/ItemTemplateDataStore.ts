import { ItemTemplateRepository } from '../../domain/repository';
import { ItemTemplate } from '../../domain/model';
import { ITEM_TEMPLATE_DATA } from './mock';

export class ItemTemplateDataStore implements ItemTemplateRepository {
  find(param?: {}): Promise<ItemTemplate[]> {
    console.log(`ItemTemplate - Find`, JSON.stringify(param));
    return Promise.resolve(ITEM_TEMPLATE_DATA);
  }
  fetchById(id: number): Promise<ItemTemplate | undefined> {
    console.log(`ItemTemplate - fetchById`, id);
    return Promise.resolve(ITEM_TEMPLATE_DATA.find(v => v.id === id));
  }

}