import { TagRepository } from '../../domain/repository';
import { Tag } from '../../domain/model';
import { TAG_DATA } from './mock';

export class TagDataStore implements TagRepository {
  find(param?: {}): Promise<Tag[]> {
    console.log(`Tag - Find`, JSON.stringify(param));
    return Promise.resolve(TAG_DATA);
  }
  fetchById(id: number): Promise<Tag| undefined> {
    console.log(`Tag - fetchById`, id);
    return Promise.resolve(TAG_DATA.find(v => v.id === id));
  }
}