import { Tag } from '../model';

export interface TagRepository  {

  find(param?: {}): Promise<Tag[]>;
  fetchById(id: number): Promise<Tag | undefined>;

}