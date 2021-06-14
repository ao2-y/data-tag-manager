import { MetaKey } from '../model';

export interface MetaKeyRepository {
  find(param: { keyword?: string; excludes?: string[] }): Promise<MetaKey[]>;
  fetchById(id: string): Promise<MetaKey>;
  create(value: MetaKey): Promise<MetaKey>;
  update(value: MetaKey): Promise<MetaKey>;
}
