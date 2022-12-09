// @ts-ignore
/* eslint-disable */
import { makeCatTree } from '@/utils/makeCatTree';
import { request } from 'umi';

/** 组织机构 POST /api/system/orgs/query */
export async function postQuery(body: API.PageQuery, options?: { [key: string]: any }) {
  const res = await request<API.Response<API.SysOrgnization[]>>(
    '/api/system/orgs/query',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
  res.data = makeCatTree(res.data, { rootKey: 0 })
  return res
}

/** 组织机构 POST /api/system/orgs/updates */
export async function postUpdates(body: API.SysOrgnization, options?: { [key: string]: any }) {
  return request<API.Response>('/api/system/orgs/updates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
