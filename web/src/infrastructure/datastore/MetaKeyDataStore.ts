import { MetaKeyRepository } from '../../domain/repository';
import { MetaKey } from '../../domain/model';
import { META_KEY_DATA } from './mock';

export class MetaKeyDataStore implements MetaKeyRepository {
  find(param?: { keyword?: string; excludes?: string[] }): Promise<MetaKey[]> {
    console.log(`MetaKey - Find`, JSON.stringify(param));
    return Promise.resolve(META_KEY_DATA);
  }
  fetchById(id: number): Promise<MetaKey | undefined> {
    console.log(`MetaKey - fetchById`, id);
    return Promise.resolve(META_KEY_DATA.find(v => v.id === id));
  }

}