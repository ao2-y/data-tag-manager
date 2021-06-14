import { inject, injectable } from 'tsyringe';
import { MetaKey } from '../domain/model';
import { MetaKeyRepository } from '../domain/repository';

@injectable()
export class MetaKeyUsecase {
  constructor(@inject('MetaKeyRepository') private metaKeyRepository: MetaKeyRepository) {}

  find(param: { keyword?: string; excludes?: string[] } = {}) {
    return this.metaKeyRepository.find(param);
  }

  fetch(id: string) {
    return this.metaKeyRepository.fetchById(id);
  }

  create(item: MetaKey) {
    return this.metaKeyRepository.create(item);
  }

  update(item: MetaKey) {
    return this.metaKeyRepository.update(item);
  }

  fetchInitData() {
    return Promise.resolve({ id: '', name: '' } as MetaKey);
  }

  validate() {}
}
