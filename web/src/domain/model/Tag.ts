export interface Tag {
  id: number;
  name: string;
  level: number;
  parent?: Tag;
  color: string;
}
