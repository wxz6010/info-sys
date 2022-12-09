// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 系统数据字典 POST /api/system/dic/query */
export async function postDicQuery(body: API.PageQuery, options?: { [key: string]: any }) {
  return request<API.Response>('/api/system/dic/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统数据字典 POST /api/system/dic/updates */
export async function postDicUpdates(body: API.SysDictionary, options?: { [key: string]: any }) {
  return request<API.Response<API.SysDictionary[]>>('/api/system/dic/updates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
