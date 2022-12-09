import FontIcons from "@/components/Icons";
import { Routes, useLayoutRoutes } from "@/hooks/useRoutes";
import { IRouteComponentProps } from "umi";
import ProLayout, { PageContainer } from '@ant-design/pro-layout';
import './index.css'
import { useEffect } from "react";

const CustomMenu = (props: { route: Routes, activity?: boolean; onItemClick: () => void }) => {
    const { route } = props
    return <div onClick={props.onItemClick} className={props.activity ? "item activivty" : "item"}>
        <FontIcons type={route.icon!} style={{ fontSize: 22, marginBottom: 5 }} />
        <div>{route.name}</div>
    </div >
}

export default (props: IRouteComponentProps) => {
    const { state, onPathChange, onParentChange } = useLayoutRoutes()
    return <div className="layout">
        <div className="slider">
            {state.routes?.map((it, idx) => <CustomMenu onItemClick={() => onParentChange(it)} key={`${it.path}--${idx}`} route={it} activity={state.route?.name === it.name} />)}
        </div>
        {state.route && <ProLayout
            location={{
                pathname: props.location.pathname,
            }}
            route={{
                routes: state.route?.routes,
            }}
            navTheme="light"
            style={{
                height: '400px',
            }}
            logo={false}
            title={state.route?.name}
            collapsedButtonRender={false}
            rightContentRender={false}
            pageTitleRender={false}
            menu={{ defaultOpenAll: true }}
            menuItemRender={(item, dom) => {
                return <div onClick={() => onPathChange(item.path!, true)} key={item.key}>
                    {dom}
                </div>
            }}
        >
            {props.children}
        </ProLayout>}
    </div>
}