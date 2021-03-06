import { inject, injectable } from 'tsyringe';
import { Item, ItemMeta, ItemTemplate } from '../domain/model';
import { ItemRepository, ItemTemplateRepository } from '../domain/repository';

@injectable()
export class ItemUsecase {
  constructor(
    @inject('ItemRepository') private itemRepository: ItemRepository,
    @inject('ItemTemplateRepository') private itemTemplateRepository: ItemTemplateRepository,
  ) { }

  find(param?: {}) {
    return this.itemRepository.find(param);
  }

  fetch(id: number) {
    return this.itemRepository.fetchById(id);
  }

  create(item: Item) {
    return Promise.resolve<Item>(item);
  }

  update(item: Item) {
    return Promise.resolve<Item>(item);
  }

  fetchInitData(templateId?: number) {
    let loadItemTemplate: Promise<ItemTemplate|undefined>;
    if (templateId) {
      loadItemTemplate = this.itemTemplateRepository.fetchById(templateId);
    } else {
      loadItemTemplate = Promise.resolve(undefined);
    }
    return loadItemTemplate
    .then(results => results || { name: '', metas: [], id: 0} as ItemTemplate)
    .then(results => ({
      id: 0, name: '', description: '', 
      metaDatas: results.metas.map(m => ({ id: 0, metaKeyId: m.id, value: ''} as ItemMeta)),
      tags: [],
    } as Item))
  }

  validate() {

  }
}