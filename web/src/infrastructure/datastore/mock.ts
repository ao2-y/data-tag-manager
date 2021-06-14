import { MetaKey, ItemTemplate, Item, Tag } from '../../domain/model';

export const META_KEY_DATA: MetaKey[] = [
  {
    id: `ID:1`,
    name: 'MetaKey1',
  },
  {
    id: `ID:2`,
    name: 'MetaKey2',
  },
  {
    id: `ID:3`,
    name: 'MetaKey3',
  },
].map((v) => new MetaKey(v.id, v.name));

export const TAG_DATA: Tag[] = [
  {
    id: 1,
    name: 'Tag1',
    level: 1,
    color: '#dc143c',
  },
  {
    id: 2,
    name: 'Tag2',
    level: 1,
    color: '#6a5acd',
  },
  {
    id: 3,
    name: 'Tag3',
    level: 1,
    color: '#9acd32',
  },
];

export const ITEM_TEMPLATE_DATA: ItemTemplate[] = [
  {
    id: 1,
    name: 'Mock ItemTemplate Data 1',
    metas: [
      {
        id: 1,
        metaKey: META_KEY_DATA[0],
      },
      {
        id: 2,
        metaKey: META_KEY_DATA[1],
      },
    ],
  },
  {
    id: 2,
    name: 'Mock ItemTemplate Data 2',
    metas: [
      {
        id: 3,
        metaKey: META_KEY_DATA[1],
      },
      {
        id: 4,
        metaKey: META_KEY_DATA[2],
      },
    ],
  },
].concat(
  [...Array(5)].map((_, index) => ({
    id: index + 3,
    name: `Mock ItemTemplate Data ${index + 3}`,
    metas: [],
  }))
);

export const ITEM_DATA: Item[] = [
  {
    id: 1,
    name: 'Mock Item Data 1',
    description: '',
    metaDatas: [
      {
        id: 1,
        metaKeyId: META_KEY_DATA[0].id,
        value: 'Value1',
      },
      {
        id: 2,
        metaKeyId: META_KEY_DATA[1].id,
        value: 'Value2',
      },
    ],
    tags: [
      {
        id: 1,
        tag: TAG_DATA[0],
      },
      {
        id: 2,
        tag: TAG_DATA[1],
      },
    ],
  },
  {
    id: 2,
    name: 'Mock Item Data 2',
    description: '',
    metaDatas: [
      {
        id: 3,
        metaKeyId: META_KEY_DATA[0].id,
        value: 'Value1',
      },
      {
        id: 4,
        metaKeyId: META_KEY_DATA[2].id,
        value: 'Value2',
      },
    ],
    tags: [
      {
        id: 3,
        tag: TAG_DATA[2],
      },
    ],
  },
  {
    id: 3,
    name: 'Mock Item Data 3',
    description: '',
    metaDatas: [
      {
        id: 3,
        metaKeyId: META_KEY_DATA[0].id,
        value: 'Value1',
      },
      {
        id: 4,
        metaKeyId: META_KEY_DATA[2].id,
        value: 'Value2',
      },
    ],
    tags: [
      {
        id: 3,
        tag: TAG_DATA[2],
      },
    ],
  },
];
