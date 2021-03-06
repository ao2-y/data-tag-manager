import { inject, injectable } from 'tsyringe';
import { Tag } from '../domain/model';
import { TagRepository } from '../domain/repository';

@injectable()
export class TagUsecase {
  constructor(@inject('TagRepository') private tagRepository: TagRepository) {}

  find(param?: {}) {
    return this.tagRepository.find(param);
  }

  fetch(id: number) {
    return this.tagRepository.fetchById(id);
  }

  create(item: Tag) {
    return Promise.resolve<Tag>(item);
  }

  update(item: Tag) {
    return Promise.resolve<Tag>(item);
  }

  fetchInitData() {
    const r = Math.floor(Math.random() * 255);
    const g = Math.floor(Math.random() * 255);
    const b = Math.floor(Math.random() * 255);
    const color = `rgb(${r},${g},${b})`;
    return Promise.resolve({ id: 0, name: '', level: 1, color } as Tag);
  }

  validate() {}
}
