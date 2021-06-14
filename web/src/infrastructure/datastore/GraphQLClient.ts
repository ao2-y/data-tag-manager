import { ApolloLink, DocumentNode, execute, makePromise } from 'apollo-link';
import { HttpLink } from 'apollo-link-http';

export class GraphQLClient {
  private readonly link: ApolloLink;

  constructor({ endpointUrl }: { endpointUrl: string }) {
    this.link = new HttpLink({ uri: endpointUrl });
  }

  // TODO: Genericsでいい感じに型定義したい
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  public async query(query: DocumentNode, variables: Record<string, any> = {}): Promise<Record<string, any>> {
    try {
      const result = await makePromise(execute(this.link, { query, variables }));
      if (result.errors) {
        throw result.errors;
      }
      return result.data as Record<string, unknown>;
    } catch (e) {
      console.error(e);
      throw e;
    }
  }
}
