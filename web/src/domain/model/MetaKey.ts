export class MetaKey {
  constructor(public readonly id: string, public readonly name: string) {}
  public static fromJSON(json: any): MetaKey {
    const { id, name } = Object.assign({}, json);
    if (typeof id !== 'string') throw new Error('id must be a string');
    if (typeof name !== 'string') throw new Error('name must be a string');
    return new MetaKey(id, name);
  }
}
