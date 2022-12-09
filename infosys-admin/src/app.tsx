import { notification } from 'antd';
import { cloneDeep, merge } from 'lodash';
import React from 'react';
import { history, RequestConfig } from 'umi';
import NiceModal from '@ebay/nice-modal-react'

export const request: RequestConfig = {


    requestInterceptors: [(url, options) => {
        // const token = getState('user')['token']
        return {
            url: `${url}`,
            options: merge(cloneDeep(options), {
                headers: { Authorization: `Bearer ` },
                interceptors: true
            }),
        };
    }],

    errorHandler: (error: any) => {
        const { data } = error;
        if (data && data.statusCode) {
            notification.error({ message: data.message })
            switch (data.statusCode) {

                // case 403:
                //     history.push('/exception/403');
                //     break;
                // case 500:
                //     history.push('/exception/500')
                //     break
                // case 204:
                //     notification.success({ message: "删除数据成功" })
                //     return;
                // case 201:
                //     notification.success({ message: '新建或修改数据成功。' })
                //     return
                // case 202:
                //     notification.success({ message: '服务器后台任务已设定。' })
                //     break
                // case 400:
                //     notification.error({ message: '发出的请求有错误', description: data.message })
                //     break
                // default:
                //     notification.error({ description: data.message, message: "服务器发生错误:" + data.code })
            }
        }
        if (!data || !data.statusCode) {
            notification.error({
                description: '您的网络发生异常，无法连接服务器',
                message: '网络异常',
            });
        }
        return data
        //throw error
    },
};


export function rootContainer(container: any) {

    return <NiceModal.Provider>
        {container}
    </NiceModal.Provider>
}