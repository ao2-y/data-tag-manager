import { MetaKey } from '../model';

export interface MetaKeyRepository  {

  find(param?: {keyword: string; excludes?: string[]}): Promise<MetaKey[]>;
  fetchById(id: number): Promise<MetaKey | undefined>;

}