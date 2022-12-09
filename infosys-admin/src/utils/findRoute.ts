export default function findRoutes(routes: any[], path: string): any {
  let route: any = null;

  for (const key in routes) {
    route = routes[key];
    if (route.path === path) {
      break;
    }
    if (path.indexOf(route.path) > -1 && route.routes) {
      return findRoutes(route.routes, path);
    }
  }

  if (route.path !== path) {
    return undefined
  }

  return route;
}
