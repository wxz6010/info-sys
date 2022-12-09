import _ from 'lodash';

export function makeCatTree(
  data: any,
  { rootKey = 'undefined' as any, parentKey = 'pid', idKey = 'id' },
) {
  var groupedByParents = _.groupBy(data, parentKey);
  var catsById = _.keyBy(data, idKey);

  _.each(_.omit(groupedByParents, rootKey), function (children, parentId) {
    catsById['' + parentId]['children'] = children;
  });
  _.each(catsById, function (cat) {
    cat.isParent = !_.isEmpty(cat['children']);
    cat['children'] = _.compact(_.union(cat['children'], cat.posts));
    cat['children'].length === 0 && (cat['children'] = undefined);
  });
  return groupedByParents[rootKey];
}


export function treeToList(tree: any[], childKey = 'children') {
  let queen: any[] = []
  let out: any[] = []
  queen = queen.concat(tree)
  while (queen.length) {
    const first = queen.shift()
    if (first[childKey]) {
      queen = queen.concat(first[childKey])
      delete first[childKey]
    }
    out.push(first)
  }
  return out
}