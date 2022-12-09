// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 机构类型 POST /api/system/orgs/type/query */
export async function postOrgsTypeQuery(body: API.PageQuery, options?: { [key: string]: any }) {
  return request<API.Response>('/api/system/orgs/type/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 机构类型 POST /api/system/orgs/type/updates */
export async function postOrgsTypeUpdates(
  body: API.OrgnizationType,
  options?: { [key: string]: any },
) {
  return request<API.Response>('/api/system/orgs/type/updates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
