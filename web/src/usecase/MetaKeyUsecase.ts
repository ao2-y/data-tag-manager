import { inject, injectable } from 'tsyringe';
import { MetaKey } from '../domain/model';
import { MetaKeyRepository } from '../domain/repository';

@injectable()
export class MetaKeyUsecase {
  constructor(
    @inject('MetaKeyRepository') private metaKeyRepository: MetaKeyRepository
  ) {}

  find(param?: { keyword: string; excludes?: string[] }) {
    return this.metaKeyRepository.find(param);
  }

  fetch(id: number) {
    return this.metaKeyRepository.fetchById(id);
  }

  create(item: MetaKey) {
    return Promise.resolve<MetaKey>(item);
  }

  update(item: MetaKey) {
    return Promise.resolve<MetaKey>(item);
  }

  fetchInitData() {
    return Promise.resolve({ id: 0, name: '' } as MetaKey);
  }

  validate() {}
}
