import { IRouteComponentProps } from "umi";
import ProTable, { ActionType, ProColumns } from '@ant-design/pro-table'
import { useRef, useState } from "react";
import service from '@/servers'
import { Button, Popconfirm } from "antd";
import orgModal, { OrgModalProps } from "./components/orgModal";
import { useModal } from "@ebay/nice-modal-react";
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
            title: "权限名称",
            dataIndex: "p_name",
            valueType: "text",
        }, {
            title: "权限标识",
            dataIndex: "p_key",
            valueType: 'text'
        },
        {
            title: "联系人",
            dataIndex: "contact_name",
            valueType: "text",
            hideInSearch: true,
            editable: false
        },
        {
            title: "联系地址",
            dataIndex: "contact_addres",
            valueType: "text",
            hideInSearch: true,
        }, {
            title: "联系电话",
            dataIndex: "contact_phone",
            valueType: "text",
            hideInSearch: true,
        },
        {
            title: "操作",
            key: "operation",
            valueType: "option",
            render: (_, row, index, action) => [
                <a key="link" onClick={() => showModal(({ data: {}, parent: row }))}>新建子级</a>,
                <a key="warn" onClick={() => showModal({ data: row })}>编辑</a>,
                <div>
                    {row.pid && <Popconfirm key={'deleted'} okType="danger" title="确认删除,请谨慎进行该操作!!" onConfirm={() => { }}>
                        <a>撤销</a>
                    </Popconfirm> || <div />}
                </div>,

            ]
        }
    ]

    const actionRef = useRef<ActionType>()
    const [dataSource, setDataSource] = useState<API.SysOrgnization[]>([])
    const modal = useModal(orgModal)


    const showModal = async (data: OrgModalProps) => {

        (await modal.show(data)) && actionRef.current?.reload()
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
                service.organization.postQuery({
                    pageSize: param.pageSize,
                    current: param.current,
                    sorts: sort,
                    query: param
                })}
            toolbar={{
                actions: [<Button type='primary' onClick={() => showModal({})}>新建机构</Button>]
            }}

        />

    </>
}