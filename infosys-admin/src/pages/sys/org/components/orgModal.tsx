import { DrawerForm, ProFormText, ProFormInstance } from '@ant-design/pro-form'
import NiceModal, { useModal } from '@ebay/nice-modal-react'
import service from '@/servers'
import { useEffect, useMemo, useRef } from 'react'

export type OrgModalProps = {
    data?: API.SysOrgnization,
    parent?: API.SysOrgnization
}

export default NiceModal.create(({ data, parent }: OrgModalProps) => {

    const modal = useModal()
    const ref = useRef<ProFormInstance>()
    const submiiter = async (data: API.SysOrgnization) => {
        const res = await service.organization.postUpdates(data)
        modal.resolve(res.success)
        return res.success
    }


    const initValues = useMemo(() => {
        if (data && parent) {
            return {
                ...data,
                pid: parent.id,
                ext1: parent.name,
                p_key: parent.p_key,
                p_name: parent.p_name,
                path: parent.path ? `${parent.path}/${parent.id}` : `${parent.id}`
            }
        } else {
            return { ...data }
        }
    }, [parent, data])

    useEffect(() => {
        setTimeout(() => {
            ref.current?.resetFields();
            ref.current?.setFieldsValue(initValues)
        }, 100)
    }, [initValues])


    return <DrawerForm<API.SysOrgnization>
        visible={modal.visible}
        onFinish={submiiter}
        onVisibleChange={e => !e && modal.hide()}
        formRef={ref}
    >
        <ProFormText name={'id'} hidden />
        <ProFormText name={'pid'} hidden />
        <ProFormText name={'path'} hidden />
        <ProFormText name={'ext1'} disabled hidden={!(initValues?.ext1)} label="上级机构名称" />
        <ProFormText rules={[{ required: true, message: "请输入机构名称" }]} label='机构名称' name={'name'} />
        <ProFormText rules={[{ required: true, message: "请输入权限名称" }]} label='权限名称' disabled={!!initValues?.p_name} name={'p_name'} />
        <ProFormText rules={[{ required: true, message: "请输入权限标识" }]} label='权限标识' disabled={!!initValues?.p_key} name={'p_key'} />

        {initValues.p_key !== 'org_store' ? <>
            <ProFormText label='公司名称' name={'ext2'} />
            <ProFormText label='公司类型' name={'ext3'} />
            <ProFormText label='法定代表人' name={'ext4'} />
            <ProFormText label='身份证号码' name={'ext5'} />
            <ProFormText label='法定代表人' name={'ext6'} />
        </> : <>
            <ProFormText label='仓库地址' name={'ext3'} />
            <ProFormText label='使用面积' name={'ext4'} />
            <ProFormText label='房屋结构' name={'ext5'} />
            <ProFormText label='楼层数' name={'ext6'} />
            <ProFormText label='建成年度' name={'ext7'} />
            <ProFormText label='产权归属' name={'ext8'} />
         </>
        }
        <ProFormText label='联系人' name={'contact_name'} />
        <ProFormText label='联系电话' name={'contact_phone'} />
        <ProFormText label='联系地址' name={'contact_addres'} />
    </DrawerForm >
})