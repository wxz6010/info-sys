import { IRouteComponentProps } from "umi";
import ProTable, { ActionType, ProColumns } from '@ant-design/pro-table'
import { useRef, useState } from "react";
import service from '@/servers'
import { Button } from "antd";
import { useModal } from "@ebay/nice-modal-react";
import rolesModal from "./components/rolesModal";
export default (prorps: IRouteComponentProps) => {
    const columns: ProColumns<API.SysOrgnization>[] = [
        {
            title: "序号",
            dataIndex: "dataIndex",
            valueType: "index",
            width: 80
        }, {
            title: "名称",
            dataIndex: "name",
            valueType: "text",
        }, {
            title: "权限标识",
            dataIndex: "s_key",
            valueType: "text",
        }, {
            title: "描述",
            dataIndex: "desc",
            valueType: "text",
        }, {
            title: "操作",
            key: "operation",
            valueType: "option",
            render: (_, row, index, action) => [
                <a key="warn" onClick={() => showModal(row)}>编辑</a>,
            ]
        }
    ]

    const actionRef = useRef<ActionType>()
    const [dataSource, setDataSource] = useState<API.SysOrgnization[]>([])
    const modal = useModal(rolesModal)

    const showModal = async (data?: API.SysRole) => {
        (await modal.show(data)) && actionRef.current?.reload();
    }


    return <>
        <ProTable

            dataSource={dataSource}
            onDataSourceChange={data => setDataSource(data || [])}
            actionRef={actionRef}
            columns={columns}
            //@ts-ignore
            rowKey={it => it.id || it.key}
            request={(param, sort) =>
                service.SysRole.postRolesQuery({
                    pageSize: param.pageSize,
                    current: param.current,
                    sorts: sort,
                    query: param
                })}
            toolbar={{
                actions: [<Button type='primary' onClick={() => showModal()}>新建权限</Button>]
            }}

        />

    </>
}