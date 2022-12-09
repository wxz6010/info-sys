// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 组织机构 POST /api/system/orgs/query */
export async function postOrgsQuery(body: API.PageQuery, options?: { [key: string]: any }) {
  return request<{ Response?: API.Response; data?: API.SysOrgnization[] }>(
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
}

/** 组织机构 POST /api/system/orgs/updates */
export async function postOrgsUpdates(body: API.SysOrgnization, options?: { [key: string]: any }) {
  return request<API.Response>('/api/system/orgs/updates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
