import { DrawerForm, ProFormText, ProFormInstance } from '@ant-design/pro-form'
import NiceModal, { useModal } from '@ebay/nice-modal-react'
import service from '@/servers'
import { useEffect, useMemo, useRef } from 'react'


export default NiceModal.create((data: API.SysRole) => {

    const modal = useModal()
    const ref = useRef<ProFormInstance>()
    const submiiter = async (data: API.SysRole) => {
        const res = await service.SysRole.postRolesUpdates(data)
        modal.resolve(res.success)
        return res.success
    }


    const initValues = useMemo(() => {
        return { ...data }
    }, [data])

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
        <ProFormText hidden name={'id'} />
        <ProFormText label="名称" name={'name'} />
        <ProFormText label="权限标识" disabled={!!initValues.s_key} name={'s_key'} />
        <ProFormText label="描述" name={'desc'} />
    </DrawerForm >
})