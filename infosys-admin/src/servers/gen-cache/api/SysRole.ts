// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 角色 POST /api/system/roles/updates */
export async function postRolesUpdates(body: API.SysRole, options?: { [key: string]: any }) {
  return request<API.Response>('/api/system/roles/updates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
