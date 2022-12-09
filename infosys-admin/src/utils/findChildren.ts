import { Routes } from '@/hooks/useRoutes';

function findMenusParent(
  menus: Array<Routes>,
  path: string,
): Array<Routes> {
  const strs = path.split('/');
  const parent: any[] = []

  strs.forEach((_, i) => {
    let path = strs.slice(0, i + 1).join('/')

    const menu = menus.find(x => x.path === (path === "" ? "/console/welcome" : path))
    menu && parent.push(menu)
  });
  return parent
}

export default findMenusParent;
