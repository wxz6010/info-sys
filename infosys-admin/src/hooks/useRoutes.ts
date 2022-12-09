import servers from "@/servers";
import findMenusParent from "@/utils/findChildren";
import findRoutes from "@/utils/findRoute";
import { treeToList } from "@/utils/makeCatTree";
import { IRoute, } from "@umijs/core";
import { cloneDeep } from "lodash";
import { type } from "os";
import { Reducer, useEffect, useReducer } from "react";
import { RouteProps } from "react-router";
import { history } from "umi";

export type Routes = IRoute & { access?: string, icon?: string, name?: string }

type State = {
    routes?: Routes[]
    pathname?: string
    route?: Routes
}


let flatRoutes: Routes[]

async function loading(): Promise<State> {
    const res = await servers.organization.postQuery({ current: 1, pageSize: 100, query: { p_key: 'org_distrubution' } })
    if (res.success && res.data) {
        const distrubutionRoutes: Routes[] = res.data[0].children!.map(it => ({
            name: it.name,
            path: `/sys/distrubution/${it.id}`,
            routes: [
                {
                    name: '用户管理',
                    path: `/sys/distrubution/${it.id}/users`
                }, {
                    name: '农户管理',
                    path: `/sys/distrubution/${it.id}/peasants`
                }, {
                    name: "订单管理",
                    path: `/sys/distrubution/${it.id}/orders`
                }]
        }))
        const state = {
            pathname: '/sys/org/list',
            routes: [{
                name: "组织机构",
                icon: 'icon-ic_org',
                path: "/sys/org",
                routes: [{
                    path: '/sys/org/list',
                    name: "机构管理"
                }, {
                    path: "sys/org/roles",
                    name: "角色管理"
                }]
            }, {
                name: "基层社/分销商",
                icon: 'icon-jicengzhandian',
                path: '/sys/distrubution',
                routes: distrubutionRoutes
            }, {
                name: "农技教学",
                icon: 'icon-wenzhang',
                routes: [{
                    name: "栏目管理"
                }, {
                    name: "广告内容管理"
                }]
            }, {
                name: '系统配置',
                icon: 'icon-xitongguanli',
                path: '/sysconfig',
                routes: [{
                    name: "数据字典",
                    path: '/sysconfig/dic/list'
                }]
            }]
        }
        return state
    }
    return {}
}


type ActionType = {
    type: 'init' | 'routeChange',
    payload: any
}

function reducer(state: State, action: ActionType): State {
    switch (action.type) {
        case 'init':
            return {
                ...action.payload,
            }
        case 'routeChange':
            return { ...state, route: action.payload }
    }
}

export function useLayoutRoutes() {
    const [state, dispatch] = useReducer(reducer, {})

    useEffect(() => {
        loading().then(payload => {
            dispatch({ type: 'init', payload });
            flatRoutes = treeToList(cloneDeep(payload.routes!), 'routes')
            if (history.location.pathname !== '/') {

                onPathChange(history.location.pathname, false, payload.routes)
            }
        })
    }, [])

    const onParentChange = (payload: Routes) => {
        dispatch({ type: 'routeChange', payload })
    }

    const onPathChange = (path: string, needPush?: boolean, routes?: Routes[]) => {
        const parent = findMenusParent(flatRoutes, path)
        if (parent.length > 0) {
            //---页面重新刷新后 state中值还没有----
            const route = (routes || state.routes)!.find(x => x.name == parent[0].name)
            needPush && history.push(path)
            dispatch({
                type: 'routeChange', payload: route,
            })
        } else {
            dispatch({ type: 'routeChange', payload: state.routes![0] })
        }
    }

    return {
        state,
        onParentChange,
        onPathChange
    }

}



