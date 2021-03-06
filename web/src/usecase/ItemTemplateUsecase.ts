import { inject, injectable } from 'tsyringe';
import { ItemTemplate } from '../domain/model';
import { ItemTemplateRepository } from '../domain/repository';

@injectable()
export class ItemTemplateUsecase {
  constructor(
    @inject('ItemTemplateRepository')
    private itemTemplateRepository: ItemTemplateRepository
  ) {}

  find(param?: {}) {
    return this.itemTemplateRepository.find(param);
  }

  fetch(id: number) {
    return this.itemTemplateRepository.fetchById(id);
  }

  create(item: ItemTemplate) {
    return Promise.resolve<ItemTemplate>(item);
  }

  update(item: ItemTemplate) {
    return Promise.resolve<ItemTemplate>(item);
  }

  fetchInitData() {
    return Promise.resolve({ id: 0, name: '', metas: [] } as ItemTemplate);
  }

  validate() {}
}
