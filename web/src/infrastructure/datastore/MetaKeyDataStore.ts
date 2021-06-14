import gql from 'graphql-tag';
import { GraphQLClient } from './GraphQLClient';
import { MetaKeyRepository } from '../../domain/repository';
import { MetaKey } from '../../domain/model';

export class MetaKeyDataStore implements MetaKeyRepository {
  private client: GraphQLClient = new GraphQLClient({ endpointUrl: '/query' });

  async find({ keyword, excludes }: { keyword?: string; excludes?: string[] }): Promise<MetaKey[]> {
    // TODO: サーバー側でフィルターするかローカルでキャッシュする仕組みを作ったほうが良さげ。
    const r = await this.client.query(
      gql`
        query MetaKey {
          metaKeys {
            id
            name
          }
        }
      `
    );
    let results = (r['metaKeys'] as []).map(MetaKey.fromJSON);
    if (keyword) {
      results = results.filter((v) => v.name.indexOf(keyword) > -1);
    }
    if (excludes) {
      results = results.filter((v) => excludes.indexOf(v.name) === -1);
    }
    return results;
  }

  async fetchById(id: string): Promise<MetaKey> {
    const r = await this.find({});
    const res = r.find((v) => v.id === id);
    if (!res) {
      // throw new Error('Notfound');
      return new MetaKey(id, 'Mock');
    }
    return res;
  }

  async create(value: MetaKey): Promise<MetaKey> {
    const r = await this.client.query(
      gql`
        mutation CreateMetaData($input: AddMetaKeyInput!) {
          addMetaKey(input: $input) {
            clientMutationId
          }
        }
      `,
      { input: Object.assign({ clientMutationId: `${new Date().getTime()}` }, { name: value.name }) }
    );
    console.log('[METADATA] Create', r);
    return value;
  }

  async update(value: MetaKey): Promise<MetaKey> {
    const r = await this.client.query(
      gql`
        mutation UpdateMetaData($input: UpdateMetaKeyInput!) {
          updateMetaKey(input: $input) {
            clientMutationId
          }
        }
      `,
      { input: Object.assign({ clientMutationId: `${new Date().getTime()}` }, { name: value.name, id: value.id }) }
    );
    console.log('[METADATA] Update', r);
    return value;
  }
}
