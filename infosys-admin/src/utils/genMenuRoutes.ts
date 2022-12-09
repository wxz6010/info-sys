import { MenuDataItem, Route, RouterTypes } from '@ant-design/pro-layout/lib/typings'

export default function genMenuRoutes(routes: Route & { menu?: MenuDataItem }[]): Route[] {

    routes.forEach(route => {
        const menu = route.menu
        delete route.menu
        if (menu) {
            Object.assign(route, menu)
        }
    });
    return routes
}